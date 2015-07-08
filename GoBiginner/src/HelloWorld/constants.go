package main

import (
	"fmt"
	"os"	// Give infromation about environment
)

func main() {
	// Print envrironment information
	fmt.Println(os.Environ())

	// Print envrironment information one per line
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	name := os.Getenv("USERNAME")

	fmt.Println(name)
}