package main

import (
	"example/helper"
	"fmt"
	"strings"
)

func main() {
	conferanceName := "Go Conference"

	const conferanceTikets int = 50
	var remainingTickets int = conferanceTikets
	bookings := []string{"John Doe", "Jane Doe", "John Smith", "Jane Smith"}

	fmt.Printf("conferenceTickets is %T, rmainingTickets is %T, conferenceName is %T \n", conferanceTikets, remainingTickets, conferanceName)
	fmt.Printf("Welcome to %v booking application\n", conferanceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferanceTikets, remainingTickets)

	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
		fmt.Printf("%s %s\n", helper.Greeting, names[1])
	}
	fmt.Printf(("The first names of the bookings are %v\n"), firstNames)

	if remainingTickets == 0 {
		fmt.Println("Sorry, we are sold out!")
		return
	}
	// for {

	// 	var firstName string
	// 	var lastName string
	// 	var email string
	// 	var userTickets int
	// 	fmt.Println("Enter your first name")
	// 	fmt.Scan(&firstName)

	// 	fmt.Println("Enter your last name")
	// 	fmt.Scan(&lastName)

	// 	fmt.Println("Enter your email")
	// 	fmt.Scan(&email)

	// 	fmt.Println("Enter number of tickets you want to book")
	// 	fmt.Scan(&userTickets)

	// 	remainingTickets = remainingTickets - userTickets

	// 	bookings = append(bookings, firstName+" "+lastName)

	// 	fmt.Printf("Array length is %d\n", len(bookings))

	// 	fmt.Printf("Thank you %s for booking %d tickets. You will receive a confirmation email at %s\n", bookings[0], userTickets, email)

	// 	firstNames := []string{}
	// 	for index, booking := range bookings {
	// 		fmt
	// 	}
	// 	fmt.Printf("We have %d tickets remaining\n", remainingTickets)
	// }
}
