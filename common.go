package main

import "strings"

func validateUserInput(firstName string, email string, userItems uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidItemNumber := userItems > 0 && userItems <= stock
	return isValidName, isValidEmail, isValidItemNumber
}
