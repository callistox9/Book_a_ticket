package main

import "fmt"

func main() {

	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	// go through Data tyoes in go.<<-uint etc..-->>
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v Tickets and %v Are still remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var userTickets uint
	var lastName string
	var midName string
	var email string

	//ask user for their name
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("enter your midname")
	fmt.Scan(&midName)
	fmt.Println("enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("enter your email")
	fmt.Scan(&email)
	//use of pointer.Memory address//Special varialbe//println(variable) to print memory address
	fmt.Println("Enter no of tickets")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	//Arrays and slices

	//var bookings [50]string
	var bookings []string // syntax var bookings =[]string{}; or bookings:=[]string{};
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf(" %v booked %v tickets", firstName, userTickets)
	fmt.Println("	you will get a confirmation at this email address")
	fmt.Printf("These are all your bookings %v\n", bookings)

	fmt.Println("Remaining tickets left	", remainingTickets)
}
