package main

import (
	"fmt"
	"strings"
)

// Get called automatically, no arguments, no return
func main() {
	module := "function basics"
	author := "nethali"

	fmt.Println(converter(module, author))

	bestFinish := bestLeagueFinishes(13, 10, 13, 17, 14, 16, 12, 2, 3, 1)

	fmt.Println(bestFinish)
}

// Multiple arguments and return is possible
func converter(module, author string) (s1, s2 string) {
	module = strings.Title(module)
	author = strings.ToUpper(author)

	return module, author
}

// Variadic function is a function
// with unknown number of arguments
func bestLeagueFinishes(finishes ... int) int {
	best := finishes[0]
	for _, i := range finishes {
		if best > i {
			best = i
		}
	}

	return best
}