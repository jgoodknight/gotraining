// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Create a program that declares two anonymous functions. One that counts down from
// 100 to 0 and one that counts up from 0 to 100. Display each number with an
// unique identifier for each goroutine. Then create goroutines from these functions
// and don't let main return until the goroutines complete.
//
// Run the program in parallel.
package main

// Add imports.
import (
	"fmt"
	"runtime"
	"sync"
)

func init() {

	// Allocate one logical processor for the scheduler to use.
	runtime.GOMAXPROCS(2)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	// Declare a wait group and set the count to two.

	// Declare an anonymous function and create a goroutine.
	go func() {
		// Declare a loop that counts down from 100 to 0 and
		// display each value.
		count_up("UP!")
		wg.Done()
		// Tell main we are done.
	}()
	// Declare an anonymous function and create a goroutine.
	go func() {
		// Declare a loop that counts down from 100 to 0 and
		// display each value.
		count_down("DOWN!")
		wg.Done()
		// Tell main we are done.
	}()

	// Wait for the goroutines to finish.
	wg.Wait()
	// Display "Terminating Program".
	fmt.Println("\nTerminating Program...")
}

func count_up(id string) {
	for i := 1; i <= 100; i++ {
		fmt.Printf("%s %d | ", id, i)
	}
}
func count_down(id string) {
	for i := 100; i > 0; i-- {
		fmt.Printf("%s %d | ", id, i)
	}
}
