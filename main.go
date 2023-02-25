package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

// package level variables
// Slice -- syntax var bookings =[]string{}; or bookings:=[]string{}; //its a slice

var conferenceName string = "Go Conference"

const conferenceTickets int = 50

// go through Data tyoes in go.<<-uint etc..-->>
var remainingTickets uint = 50
var bookings = make([]UserData, 0) //empty list of user data struct

// var bookings = make([]map[string]string, 0) //creating empty list of maps
// Learning about struct
// type keyword creates a new type, with the neame you specify
type UserData struct {
	firstName string
	lastName  string //mixed datatype >> defining a structure (which fields) of the User Type
	//it iws somewhat like a light weight class, which e.g doesn't support inheritance
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	// Infinite loop
	//for {

	println("Do you wantr to book a ticket?y/n")
	var ch string
	fmt.Scan(&ch)
	if ch == "y" {
		if remainingTickets == 0 {
			fmt.Println("We are completely booked,Try next year")
			//break
		}
		var firstName, lastName, email, userTickets = userInput()
		// checking user validation
		isValidName, isValidEmial, isValidTicketNumber := helper.Validation(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmial && isValidName && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)

			go sendTicket(userTickets, firstName, lastName, email)

			//concurrency in go

			var k = firstnames()                                   // stroing the return value of function firstname into k
			fmt.Printf("The first names of bookings are %v \n", k) //printing the return value of the function that is firstname

			fmt.Printf("\n %v booked %v tickets \n", firstName, userTickets)
			fmt.Println("	you will get a confirmation at this email address", email)

			fmt.Printf("These are all your bookings %v\n", userTickets)

			fmt.Println("Remaining tickets left	", remainingTickets)
			println("Thank you")

		} else {
			if !isValidEmial {
				fmt.Println("Please enter valid email")
			}
			if !isValidName {
				fmt.Println("Enter valid Name")

			}
			if !isValidTicketNumber {
				fmt.Printf("We only have %v tickets with us, Please enter valid ticket number\n", remainingTickets)
			}

		}
	} else {
		println("Thank you for showing interest")
		//break
	}
	wg.Wait()
}

//}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have total of %v Tickets and %v Are still remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
func firstnames() []string {
	firstNames := []string{}
	for _, booking := range bookings { //blank identifier '_'.To ignore variables you dont want to use
		//var names = strings.Fields(booking) // names is an array which stores the split string after space
		//var firstName = names[0];//the first index of the names which is first name is stored in it
		firstNames = append(firstNames, booking.firstName)

	}
	return firstNames
}

// passing multiple values from 1 function

func userInput() (string, string, string, uint) {
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
	return firstName, lastName, email, userTickets
}
func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTickets

	//Arrays and slices
	//use of mapsss

	//var bookings [50]string
	//useOfMAps
	//var userData = make(map[string]string) //bookings map

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	//userData["firstName"] = firstName
	//userData["lastName"] = lastName
	//userData["email"] = email
	//FormatUint(i uint64, base int) string
	//FormatUint returns the string representation of i in the given base,
	//for 2 <= base <= 36. The result uses the lower-case letters 'a' to 'z'
	//for digit values >= 10.
	//userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) //converting uint to string

	bookings = append(bookings, userData)
	fmt.Printf("List of Bookings is  %v \n", bookings)

}

// GOroutines
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(30 * time.Second) // the Sleep function stops or blocks the current "thread" (goRoutine) execution for the defined duration
	var ticket = fmt.Sprintf("%v tickets for %v %v \n", userTickets, firstName, lastName)
	fmt.Println("#########")
	fmt.Printf("Sending %v to email address %v \n", ticket, email)
	fmt.Println("#########")
	wg.Done()

}
