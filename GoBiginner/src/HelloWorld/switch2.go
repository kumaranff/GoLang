package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	switch tmpNum := random(); tmpNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("You got an even number :", tmpNum)
	case 1, 3, 5, 7, 9:
		fmt.Println("You got an odd number :", tmpNum)
	}
}

func random() int {
	rand.Seed(time.Now().Unix())
	// Generate numbe between 0-1
	return rand.Intn(10)
}
