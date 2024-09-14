package main

import (
	"fmt"
	"strings"
)

func main() {
	//declear variables
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	var bookings = []string{}
	fmt.Println("Welcome to", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// ask user for thier name and details
	// array
	// var bookings [conferenceTickets]string
	// var bookings = [conferenceTickets]string{}

	//slice

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		fmt.Printf("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Printf("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Printf("Enter your email address: ")
		fmt.Scan(&email)
		fmt.Printf("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v for booking %v tickets! You will receive the confirmation email at %v\n", bookings[0], userTickets, email)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("Print are all our bookings %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come next year.")
				break
			} else {

			}
		} else {
			fmt.Printf(" We have only %v tickets remaining.", remainingTickets)
		}

	}

}
