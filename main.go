package main

import (
	"fmt"
	"sync"
)

const totalItem = 100

var stock uint = 100
var storeName = "One Item Store"
var totalSells = make([]UserData, 0)

type UserData struct {
	firstName     string
	email         string
	numberOfItems uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	firstName, email, userItems := getUserInput()
	isValidName, isValidEmail, isValidItemNumber := validateUserInput(firstName, email, userItems)

	if isValidName && isValidEmail && isValidItemNumber {

		buyItem(userItems, firstName, email)
		wg.Add(1)
		go buyItem(userItems, firstName, email)

		firstNames := getFirstName()
		fmt.Printf("The first names of our best costumers are: %v\n", firstNames)

		if stock == 0 {
			// end program
			fmt.Println("Our store is out of inventory. Come back next week.")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("First name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("The email address ypu entered doesn't contain @ sing")
		}
		if !isValidItemNumber {
			fmt.Println("Number of Items you entered is invalid")
		}
	}
	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %v \n", storeName)
	fmt.Println("The store where you can only buy one item")
	fmt.Printf("Total item %v \n", totalItem)
	fmt.Printf("Remaining stock %v \n", stock)
}

func getFirstName() []string {
	firstNames := []string{}
	for _, nameList := range totalSells {
		firstNames = append(firstNames, nameList.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, uint) {

	var firstName string
	var email string
	var userItems uint

	fmt.Println("Enter Your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter Your email address: ")
	fmt.Scan(&email)

	fmt.Println("Number of items you need: ")
	fmt.Scan(&userItems)

	return firstName, email, userItems
}

func buyItem(userItems uint, firstName string, email string) {
	stock = stock - userItems

	var userData = UserData{
		firstName:     firstName,
		email:         email,
		numberOfItems: userItems,
	}

	totalSells = append(totalSells, userData)
	fmt.Println("Here our best buyers: ")
	fmt.Printf("%v\n", totalSells)

	fmt.Printf("Thank you %v for buying %v items, a confirmation receipt will be sent to %v \n", firstName, userItems, email)
	fmt.Printf("%v items left \n", stock)
}
