package common

import (
	"fmt"
	"strconv"
	"strings"
)

func ValidateUserInput(firstName string, lastName string, email string) (bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2 
	isValidEmail := strings.Contains(email, "@")

	return isValidName, isValidEmail
}

func CreateUserWithMap(firstName string, lastName string, email string, noOfTickets uint) map[string]string {
	userData := make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["ticketsPurchased"] = strconv.Itoa(int(noOfTickets)) // convert int to string

	return userData
}

func PrintFirstNamesString(bookings []string) {
	firstNames := []string{}
	for _, booking := range bookings  {
		// Split sting to individual array values
		name := strings.Fields(booking)
		firstNames = append(firstNames, name[0])
	}

	fmt.Printf("All first names from bookings are %v\n", firstNames)
}

func PrintFirstNamesMap(bookings []map[string]string) {
	firstNames := []string{}
	for _, booking := range bookings  {
		firstNames = append(firstNames, booking["firstName"])
	}

	fmt.Printf("All first names from bookings are %v\n", firstNames)
}