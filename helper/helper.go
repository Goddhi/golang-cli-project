package helper

import (
	"strings"
)

func ValidateUserInput (firstName string, lastName string, email string, userTickets uint, remainingTicket uint) (bool, bool, bool) {
	isValidFirstNameLastName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTicket

	return isValidFirstNameLastName, isValidEmail, isValidTicketNumber
}