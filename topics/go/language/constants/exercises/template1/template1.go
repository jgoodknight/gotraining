// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare an untyped and typed constant and display their values.
//
// Multiply two literal constants into a typed variable and display the value.
package main

// Add imports.
import (
	"fmt"
)

const (
	// Declare a constant named server of kind string and assign a value.
	server = "867.53.09"
	// Declare a constant named port of type integer and assign a value.
	port = 9999
)

func main() {

	// Display the value of both server and port.
	fmt.Println("Server: ", server)
	fmt.Println("Port: ", port)

	// Divide a constant of kind integer and kind floating point and
	// assign the result to a variable.
	const float_example = .0000001
	a := float_example / port

	// Display the value of the variable.
	fmt.Println(a)
}
