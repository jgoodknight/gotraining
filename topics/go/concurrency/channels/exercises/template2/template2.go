// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that uses a fan out pattern to generate 100 random numbers
// concurrently. Have each goroutine generate a single random number and return
// that number to the main goroutine over a buffered channel. Set the size of
// the buffered channel so no send ever blocks. Don't allocate more capacity
// than you need. Have the main goroutine store each random number it receives
// in a slice. Print the slice values then terminate the program.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Add imports.

// Declare constant for number of goroutines.

func init() {
	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())
}

func main() {
	num_rand_desired := 100
	// Create the buffered channel with room for
	// each goroutine to be created.

	var random_output_list []int

	storage_channel := make(chan int, num_rand_desired)

	// Storage goroutine:

	// Iterate and launch each goroutine.
	for i := 0; i < num_rand_desired; i++ {
		go func() {
			new_rand := rand.Intn(500)
			storage_channel <- new_rand
		}()
	}
	// for new_rand := range storage_channel {
	// 	random_output_list = append(random_output_list, new_rand)
	// }
	for len(random_output_list) < num_rand_desired {
		new_rand := <-storage_channel
		random_output_list = append(random_output_list, new_rand)
	}

	for i, rand_val := range random_output_list {
		fmt.Println(i, rand_val)
	}
}
