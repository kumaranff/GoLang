package main

import (
	"fmt"
)

func main() {
	// switch as default (implicit) break trough behaviour
	switch "nethali"{
	case "nethali":
		fmt.Println("Nethali")
	case "kumara":
		fmt.Println("Kumara")
	case "zoysa":
		fmt.Println("Zoysa")
	default:
		fmt.Println("None")
	}	

	// Simple statement is allowed in switch 
	switch name := "kumara1"; name {
	case "nethali":
		fmt.Println("Nethali")
	case "kumara":
		fmt.Println("Kumara")
	case "zoysa":
		fmt.Println("Zoysa")
	default:
		fmt.Println("None")
	}	

	// If you need fall through behaviour do it as follows
	switch name := "kumara"; name {
	case "nethali":
		fmt.Println("Nethali")
		fallthrough
	case "kumara":
		fmt.Println("Kumara")
		fallthrough
	case "zoysa":
		fmt.Println("Zoysa")
		fallthrough
	default:
		fmt.Println("None")
	}	
}