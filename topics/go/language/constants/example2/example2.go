// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how constants do have a parallel type system.
package main

import (
	"fmt"
	"math"
)

const (
	// Max integer value on 64 bit architecture.
	maxInt = 9223372036854775807

	// Much larger value than int64.
	bigger = 922337203685477580854352234599999900000

	// Will NOT compile
	// biggerInt int64 = 9223372036854775808543522345
)

func main() {
	fmt.Println("Will Compile")
	fmt.Println(math.Pi)
}
