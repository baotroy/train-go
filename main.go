package main

import "fmt"

func main() {
	var conferanceName string = "Go Conference"

	const conferanceTikets int = 50
	var remainingTickets int = conferanceTikets

	fmt.Printf("conferenceTickets is %T, rmainingTickets is %T, conferenceName is %T \n", conferanceTikets, remainingTickets, conferanceName)
	fmt.Printf("Welcome to %v booking application\n", conferanceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferanceTikets, remainingTickets)

	fmt.Println("Get your ticket here")

	var userName  string
	var lastName string
	var email string
	var userTickets int
	fmt.Println("Enter your first name")
	fmt.Scan(&userName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)
	

	// userName = "John Doe"
	userTickets = 2
	fmt.Printf("%s booked %d tickets\n", userName, userTickets)
}