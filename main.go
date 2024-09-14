package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets
	var bookings []string

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	//for { // infinite loop
	for remainingTickets > 0 && len(bookings) < 50 { //equivalant of a while loop. Set a condition after 'for' keyword
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name:	")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name:	")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email:	")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets:	")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets -= userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thankyou %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}

			for _, booking := range bookings { //typical foreach loop (for index, value:= range list). Use _ to mark as not used
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Booked out!")
				break //break and continue behave very similar to C#
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is to short")
			}

			if !isValidEmail {
				fmt.Println("Email didn't contain @")
			}

			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	}
}
