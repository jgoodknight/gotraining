// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that inspects a user's name and greets them in a certain way
// if they are on a list or in a different way if they are not. Also look at
// the user's age and tell them some special secret if they are old enough to
// know it.
package main

// Add imports
import "fmt"

func main() {

	// Change these values and rerun your program.
	name := "Carter"
	age := 6

	// If the user's name is on a special list then give them a secret greeting.
	switch name {
	case "Joey", "Chetty", "Dustin":
		fmt.Println("Job Alerts is the best!")
	default:
		fmt.Println("No!  Not for you...")
	}
	// If the user is old enough then tell them a secret.
	if age > 5 {
		fmt.Println("Someday, you will know...")
	} else if age > 12 {
		fmt.Println("You are not yet wise enough")
	} else if age > 21 {
		fmt.Println("Beer before Liquor, never been sicker")
	}

}
