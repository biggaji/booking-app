package main

import (
	"fmt"
	"strings"
)

type bookingsData struct {
	firstName     string
	lastName      string
	userName      string
	no_of_tickets uint
}

var conferenceName = "Crypto Geeks"
var no_of_network = 30

const totalAvailableTickets = 60

var remainingTickets uint = 60
var attendees = make([]bookingsData, 0) //A slice for all attendees

// entry point for all go projects
func main() {

	greetAttendees(conferenceName, totalAvailableTickets, remainingTickets, no_of_network)

	for {

		firstName, lastName, userName, ticket := collectTicketBookingInputs()
		fmt.Println(firstName, lastName, userName, ticket)
		//write some validations for our booking system
		isValidTicketNum, isValidName, isValidUsername := validateTicketBooking(firstName, lastName, userName, ticket)

		bookingValidations := isValidName && isValidUsername && isValidTicketNum
		fmt.Println(isValidName, isValidTicketNum, isValidUsername)

		if bookingValidations {
			createTicketBooking(firstName, lastName, userName, ticket)

			// check if tickets are sold out
			if remainingTickets == 0 {
				fmt.Printf("Tickets for %v conference has been sold out! Please try again next year!\n", conferenceName)
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Firstname and Lastname should be greater than 2 characters")
			}
			if !isValidUsername {
				fmt.Println("Username should be at least 3 characters long")
			}
			if !isValidTicketNum {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}
	}
}

// functions in Go
func greetAttendees(confName string, totalTickets int, remainTickets uint, networks int) {
	fmt.Println("Welcome to", confName, "conference booking service!")
	fmt.Println("We have a total of", totalTickets, "tickets and", remainTickets, "are still available")
	fmt.Println("Grab a ticket to attend!")
	fmt.Println("We have", networks, "network providers at this event.")
}

func collectTicketBookingInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var userName string
	var ticket uint

	println("Please enter your first name:")
	fmt.Scan(&firstName)

	println("Please enter your last name:")
	fmt.Scan(&lastName)

	println("Please enter your username:")
	fmt.Scan(&userName)

	println("How many ticket do you want to buy:")
	fmt.Scan(&ticket)

	return firstName, lastName, userName, ticket
}

// using return key word we have to specifically provide it return type e.g
func getFirstNames(names []bookingsData) []string {
	// use _ for blank identifiers
	users := []string{} //replace index to _, since it wont be used now
	for _, variable := range names {
		names := strings.Fields(variable.firstName)
		users = append(users, names[0])
	}
	return users
}

func validateTicketBooking(firstName string, lastName string, userName string, ticket uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidUsername := len(userName) >= 3
	isValidTicketNum := ticket >= 1 && ticket <= remainingTickets

	return isValidName, isValidUsername, isValidTicketNum
}

func createTicketBooking(firstName string, lastName string, userName string, ticket uint) {

	remainingTickets -= ticket

	bookings := bookingsData{
		firstName:     firstName,
		lastName:      lastName,
		userName:      userName,
		no_of_tickets: ticket,
	}

	attendees = append(attendees, bookings)

	fmt.Println(getFirstNames(attendees))
	fmt.Println(attendees)

	tickStr := "tickets"
	if ticket <= 1 {
		tickStr = "ticket"
	}

	remainTickStr := "tickets"
	if remainingTickets <= 1 {
		remainTickStr = "ticket"
	}

	userLenStr := "users"
	if len(attendees) == 1 {
		userLenStr = "user"
	}

	fmt.Printf("Thank you %v for buying %v conference %v.\n", userName, ticket, tickStr)
	fmt.Printf("%v %v are available for %v conference event!\n", remainingTickets, remainTickStr, conferenceName)
	fmt.Printf("%v %v had successfully booked.\n", len(attendees), userLenStr)
	fmt.Printf("%v will be attending the %v conference this year\n", attendees, conferenceName)
}
