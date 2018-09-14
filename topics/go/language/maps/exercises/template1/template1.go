// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare and make a map of integer values with a string as the key. Populate the
// map with five values and iterate over the map to display the key/value pairs.
package main

// Add imports.
import (
	"fmt"
	"sort"
)

func main() {

	// Declare and make a map of integer type values.
	my_int_map := map[int]int{
		1: 2,
		2: 3,
	}
	// Initialize some data into the map.
	my_int_map[3] = 4
	my_int_map[4] = 5

	// Display each key/value pair.
	for key, value := range my_int_map {
		fmt.Println(key, value)
	}
	fmt.Println("****************************")
	var key_list []int
	for key := range my_int_map {
		key_list = append(key_list, key)
	}
	sort.Ints(key_list)

	for _, key := range key_list {
		fmt.Println(key, my_int_map[key])
	}
	fmt.Println("****************************")

	_, ok := my_int_map[99]

	if !ok {
		fmt.Println("I got problems, but 99 ain't one")
	}
}
