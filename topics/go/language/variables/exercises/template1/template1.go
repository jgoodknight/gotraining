// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

// Add imports
import "fmt"

// main is the entry point for the application.
func main() {

	// Declare variables that are set to their zero value.
	var a string
	var b int
	var c bool

	// Display the value of those variables.
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	aa := "hello"
	bb := 137
	cc := false

	// Display the value of those variables.
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)

	// Perform a type conversion.
	bbb := int64(137)

	// Display the value of that variable.
	fmt.Println(bbb)

	fmt.Printf("%T [%v]\n", bbb, bbb)
}
