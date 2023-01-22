package main

import (
	"booking/common"
	"fmt"
	"sync"
	"time"
)

type UserData struct {
	firstName string
	lastName string
	email string
	ticketsPurchased uint
}

var wg = sync.WaitGroup{}

func main()  {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var availableTickets uint = 50
	bookings := make([]UserData, 0)
	// bookings := make([]map[string]string, 0) // slice of maps. 0 is the starting size of the slice

	greetings(conferenceName, conferenceTickets, availableTickets)

	firstName, lastName, email, userTickets := getUserInput()

	// Input validation
	isValidName, isValidEmail := common.ValidateUserInput(firstName, lastName, email)
	if !isValidName || !isValidEmail {
		fmt.Println("Invalid Name or Email Input")
		return
	}

	// Check ticket availability
	if userTickets > availableTickets {
		fmt.Printf("Available ticket left are %v, You are trying to book %v\n", availableTickets, userTickets)
		return
	}

	// Update the remaining tickets
	availableTickets = availableTickets - userTickets
	// userData := common.CreateUserWithMap(firstName, lastName, email, userTickets)
	userData := UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		ticketsPurchased: userTickets,
	}

	bookings = append(bookings, userData)
	wg.Add(1)
	go sendTicket(userTickets, firstName, lastName, email)

	fmt.Printf("Hi %v %v you have successfully booked %v tickets, You will receive email confirmation at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for the %v\n", availableTickets, conferenceName)
	fmt.Printf("All bookings %v\n", bookings)

	// Print only first names from booking array
	printFirstNames(bookings)

	// End cli program if tickets  are finished
	if availableTickets <= 0 { 
		fmt.Println("Tickets sold out")
		return
	}

	wg.Wait()
}

func greetings(confName string, confTickets int, availableTickets uint) {
	fmt.Printf("Welcome to %v ticket booking application\n", confName)
	fmt.Printf("We have a total of %d tickets and %d are still available.\n", confTickets, availableTickets)
	fmt.Println("Get your tickets here to attend")
}

func printFirstNames(bookings []UserData) {
	firstNames := []string{}
	for _, booking := range bookings  {
		firstNames = append(firstNames, booking.firstName)
	}

	fmt.Printf("All first names from bookings are %v\n", firstNames)
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Prompt user inputs
	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("How many tickets will you like to purchase: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second) // Simulate Blocking task
	tickets := fmt.Sprintf("%v Tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#######################################")
	fmt.Printf("Sending ticket:\n%v \nto email address %v\n", tickets, email)
	fmt.Println("#######################################")

	wg.Done()
}
