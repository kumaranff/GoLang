package main

import (
	"fmt"
)

func main() {
	// Variables to store course ranking
	firstRank := 3333
	secondRank := 643

	if firstRank < secondRank {
		fmt.Println("The first cours is better" +
			" than the secode course")
	} else if firstRank > secondRank {
		fmt.Println("The second cours is better" +
			" than the first course")
	} else {
		fmt.Println("Both courses are same")
	}

	// if also accept optional initialization statement
	// Any variables in this statement is local to the
	// if, else if, else chunk
	if firstRank, secondRank := 12, 34; firstRank < secondRank {
		fmt.Println("The first cours is better" +
			" than the secode course")
	} else if firstRank > secondRank {
		fmt.Println("The second cours is better" +
			" than the first course")
	} else {
		fmt.Println("Both courses are same")
	}
}