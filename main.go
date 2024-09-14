package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets
	var bookingsArray [50]string
	var bookingsSlice []string

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

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

	remainingTickets -= userTickets
	bookingsArray[0] = firstName + " " + lastName
	bookingsSlice = append(bookingsSlice, firstName+" "+lastName)

	fmt.Printf("The whole array: %v\n", bookingsArray)
	fmt.Printf("the first value: %v\n", bookingsArray[0])
	fmt.Printf("Array type: %T\n", bookingsArray)
	fmt.Printf("Array length: %v\n", len(bookingsArray))

	fmt.Printf("The whole slice: %v\n", bookingsSlice)
	fmt.Printf("the first value: %v\n", bookingsSlice[0])
	fmt.Printf("Slice type: %T\n", bookingsSlice)
	fmt.Printf("Slice length: %v\n", len(bookingsSlice))

	fmt.Printf("Thankyou %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
