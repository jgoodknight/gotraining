// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a player's batting average. The formula is hits / atBats.
// Declare a slice of this type and initialize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
package main

// Add imports.
import (
	"fmt"
)

// Declare a struct that represents a ball player.
// Include fields called name, atBats and hits.
type BaseballPlayer struct {
	name   string
	atBats int
	hits   int
}

// Declare a method that calculates the batting average for a player.
func (baseball_player *BaseballPlayer) average() float32 {
	return float32(baseball_player.hits) / float32(baseball_player.atBats)
}

func main() {

	// Create a slice of players and populate each player
	// with field values.
	dodgers_lineup := []BaseballPlayer{
		{"Who", 234, 123},
		{"What", 8675, 309},
		{"I don't know", 137, 9},
		{"Why", 2, 1},
		{"Because", 4, 3},
		{"I don' give a darn", 6, 5},
	}

	// Display the batting average for each player in the slice.
	for i := range dodgers_lineup {
		dodger := &dodgers_lineup[i]
		fmt.Println(dodger.name, " has batting average", dodger.average())
	}
}
