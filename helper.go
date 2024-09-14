package main

import "strings"

func validateUserInput(userData UserData, remainingTickets uint) (bool, bool, bool) { // multiple returns! declare inside ()

	isValidName := len(userData.firstName) >= 2 && len(userData.lastName) >= 2
	isValidEmail := strings.Contains(userData.email, "@")
	isValidTicketNumber := userData.numberOfTickets > 0 && userData.numberOfTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
