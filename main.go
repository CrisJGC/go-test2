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

	firstNsme, email, userItems := getUserInput()

	if userItems > stock {
		fmt.Printf("We only have %v items, so we can't sell you %v items.\n", stock, userItems)
		continue
	}

	stock = stock - userItems
	totalSells = append(totalSells, firstName)

	// fmt.Printf("Array: %v\n", totalSells)
	// fmt.Printf("First value %v\n", totalSells[0])
	// fmt.Printf("Array type: %T\n", totalSells)
	// fmt.Printf("Array length: %v\n", len(totalSells))

	if stock == 0 {
		fmt.Println("All our items are sold")
		break
	}

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
