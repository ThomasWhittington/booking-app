package main

import (
	"fmt"
	"strconv"
)

const conferenceTickets uint = 50

var conferenceName = "Go conference"
var remainingTickets uint = conferenceTickets
var bookings = make([]map[string]string, 0)

func main() {

	greetUsers()

	//for { // infinite loop
	for remainingTickets > 0 && len(bookings) < 50 { //equivalant of a while loop. Set a condition after 'for' keyword
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(firstName, lastName, email, userTickets)
			firstNames := getFirstNames(bookings)
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

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []map[string]string) []string { // return type is put after the parameters
	firstNames := []string{}

	for _, booking := range bookings { //typical foreach loop (for index, value:= range list). Use _ to mark as not used
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets -= userTickets

	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)

	fmt.Printf("Thankyou %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
