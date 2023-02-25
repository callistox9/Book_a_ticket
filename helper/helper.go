package helper

//Scope:->> Package level
//variables and functions defined outside function, can be accessed in all other files within the same package
import (
	"strings"
)

// we are importing package from user created package.We set the import path, then call by . like fmt.Print and capital letter of the function
func Validation(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmial := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets //we can also export variable
	return isValidName, isValidEmial, isValidTicketNumber

}

//Multiple packages in your application
