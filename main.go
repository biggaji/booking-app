package main

import "fmt"

// entry point for all go projects
func main() {
	var conferenceName = "Crypto Geeks"
	const totalAvailableTickets = 60
	var remainingTickets = 60
	const programExitCode = 0

	fmt.Println("Welcome to", conferenceName, "conference booking service!")
	fmt.Println("We have a total of", totalAvailableTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Grab a ticket to attend!")

	//using printf() we will replace values with the placeholder %v
	fmt.Printf("Program exited %v\n", programExitCode)

	//%T is used to get the datatype
}
