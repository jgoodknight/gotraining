// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Copy the code from the template. Declare a new type called hockey
// which embeds the sports type. Implement the finder interface for hockey.
// When implementing the find method for hockey, call into the find method
// for the embedded sport type to check the embedded fields first. Then create
// two hockey values inside the slice of haystacks and perform the search.
package main

import (
	"fmt"
	"strings"
)

// finder defines the behavior required for performing matching.
type finder interface {
	find(needle string) bool
}

// sport represents a sports team.
type sport struct {
	team string
	city string
}

// find checks the value for the specified term.
func (s *sport) find(needle string) bool {
	return strings.Contains(s.team, needle) || strings.Contains(s.city, needle)
}

// Declare a struct type named hockey that represents specific
// hockey information:
// - Have it embed the sport type first.
// - Have it include a field with the country of the team.
type hockey struct {
	sport
	country string
}

// find checks the value for the specified term.
func (h *hockey) find(needle string) bool {

	// Make sure you call into find method for the embedded sport type.

	// Implement the search for the new fields.
	return h.sport.find(needle) || strings.Contains(h.country, needle)
}

func main() {

	// Define the term to find.

	// Create a slice of finder values and assign values
	// of the concrete hockey type.
	my_finders := []finder{
		&hockey{sport{"isotopes", "Albuquerque"}, "USA"},
		&hockey{sport{"dodgers", "LA"}, "USA"},
		&sport{"dodgers", "Brooklyn"},
	}

	// Display what we are trying to find.
	target_needle := "od"
	// Range over each finder value and check the term.
	for i := range my_finders {
		f := my_finders[i]
		if f.find(target_needle) {
			fmt.Println(f)
		}
	}
}
