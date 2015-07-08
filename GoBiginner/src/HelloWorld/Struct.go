/*
Struct is user defined type.
Go have no Objects, no Classes, no Inheritance.
Although you can do Object Oriented programming with
Go, it may not be the best. In Go, any cutom type
can have mothod associated with it. Not just struct
*/

package main

import (
	"fmt"
)

func main() {
	// Define a struct
	type courseMeta struct{
		Author string
		Level string
		Rating float64
	}

	// Define the user defined types
	// Zero intialized 
	var course1 courseMeta
	fmt.Println(course1)

	// Custom initialized
	course2 := courseMeta{
		"Nethali",
		"1",
		1.23,
	}
	fmt.Println(course2)	

	// Pointer 
	course3 := new(courseMeta)
	fmt.Println(*course3)

	// Custom initialized
	course4 := courseMeta{
		Author: "Medi",
		Level: "3",
	}
	fmt.Println(course4)	
	fmt.Println("Course author:",course4.Author)
}