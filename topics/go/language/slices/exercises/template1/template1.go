// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

// Add imports.
import "fmt"

func main() {

	// Declare a nil slice of integers.
	var my_slice []int

	// Appends numbers to the slice.
	my_slice = append(my_slice, 8)
	my_slice = append(my_slice, 6)
	my_slice = append(my_slice, 7)
	my_slice = append(my_slice, 5)
	my_slice = append(my_slice, 3)
	my_slice = append(my_slice, 0)
	my_slice = append(my_slice, 9)

	// Display each value in the slice.
	for i, val := range my_slice {
		fmt.Println(i, val)
	}

	fmt.Println("==========================")
	// Declare a slice of strings and populate the slice with names.
	
	var my_team = make([]string, 3, 5)

	// Display each index position and slice value.
	my_team[0] = "Peter"
	my_team[1] = "Joey"
	my_team[2] = "Chetty"
	// Take a slice of index 1 and 2 of the slice of strings.
	var my_new_team = my_team[1:3]
	// Display each index position and slice values for the new slice.
	for i, val := range my_new_team {
		fmt.Println(i, val)
	}
	fmt.Println(":(")
}
