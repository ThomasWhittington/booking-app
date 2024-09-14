package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets uint = 50

var conferenceName = "Go conference"
var remainingTickets uint = conferenceTickets
var bookings = make([]UserData, 0)

type UserData = struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	//for { // infinite loop
	for remainingTickets > 0 && len(bookings) < 50 { //equivalant of a while loop. Set a condition after 'for' keyword
		userData := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(userData, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userData)
			wg.Add(1)               //add to wait group the number of new threads that main needs to wait for before closing
			go sendTicket(userData) // just add the 'go' keyword to run function in new thread!
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
	wg.Wait() //wait at this point for all threads to finish
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []UserData) []string { // return type is put after the parameters
	firstNames := []string{}

	for _, booking := range bookings { //typical foreach loop (for index, value:= range list). Use _ to mark as not used
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() UserData {
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

	return UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
}

func bookTicket(userData UserData) {
	remainingTickets -= userData.numberOfTickets

	bookings = append(bookings, userData)

	fmt.Printf("Thankyou %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", userData.firstName, userData.lastName, userData.numberOfTickets, userData.email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userData UserData) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userData.numberOfTickets, userData.firstName, userData.lastName)
	fmt.Println("##########")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, userData.email)
	fmt.Println("##########")
	wg.Done() //mark thread as completed
}
