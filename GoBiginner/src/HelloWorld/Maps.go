/*
Maps are unordered.
Maps are set of key value pairs
Maps are implimentation of hash-table
Maps are references
Maps are dynamically resizable (obviosly with overhead)
Giving a size at the declaration for big Maps improves the performance
*/

package main

import (
	"fmt"
)

func main() {
	// Declare a map, the key type must be copariable type
	// make(map[<key>]vlaue, size), where size is optional
	leaguTitles := make(map[string]int, 10)

	leaguTitles["Nethali"] = 3
	leaguTitles["Medi"] = 4

	// Another way of creating map
	persons := map[int]string {
		23: "Nethali",
		22: "Medi",
	}

	// Print both maps
	fmt.Printf("\nleguTitles: %v\nperson: %v",
		leaguTitles, persons)

	testMap := map[string]int{"A":1, "B":2, "C":3, "D":4, "E":5}

	// Iterating maps
	// This will give random starting point
	for key, value := range testMap {
		fmt.Printf("Key: %v, Value: %d\n", key, value)
	}

	// Update the Map
	testMap["A"] = 100
	fmt.Println(testMap)

	// Instert into the Map
	testMap["F"] = 123
	fmt.Println(testMap)

	// Delete
	delete(testMap, "A")
	fmt.Println(testMap)
}