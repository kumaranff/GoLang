package main

import (
	"fmt"
)

/*
Arrays are fixed sized contiguous memory.
Slices are built on top of arrays. Slice
is just a reference to a contiguous portion
of an Array. Slices can be resized and what 
stroed in slice is: name, type, starting offset
of the underling array and the length.
*/

func main() {
	// Declares a slice with built in make()
	// make(<type>, <len>, <cap>)
	// <cap> is the maximum size of underling array
	// [] says we are making a slice
	myCourses := make([]string, 5, 10)

	// Note that % doesn't work with Println
	fmt.Printf("Lenght is %d and Capacity is %d\n", len(myCourses), cap(myCourses))

	// Also we can create slice like this
	myCourses2 := []string{"OK", "We", "Go"}

	fmt.Printf("Lenght is %d and Capacity is %d\n", len(myCourses2), cap(myCourses2))

	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	mySlice[1] = 0
	// Referencing a slice will refrence the entire slice
	fmt.Println(mySlice) 

	// You can print it using range which return two values
	// index (which is disregard with _), and data
	for _,i := range mySlice {
		fmt.Println("for range loop: ", i)
	}

	// Slices of slice. See the range is [), first element included
	// But not the last element.
	sliceOfSlice1 := mySlice[2:5]
	fmt.Println(sliceOfSlice1)

	// Start is 0
	sliceOfSlice2 := mySlice[:5]
	fmt.Println(sliceOfSlice2)

	// End is end of mySlice
	sliceOfSlice3 := mySlice[2:]
	fmt.Println(sliceOfSlice3)
	
	/* 
	Adding values to slice with append()
	If the underling array is not enough to append
	a new value, the append() will automatically
	create a new array of size of 2 times of current
	array and copy all values.
	*/ 
	newSlice := make([]int, 1, 4)

	fmt.Printf("Lenght is %d and Capacity is %d\n", len(newSlice), cap(newSlice))

	for i:=1; i<17; i++{
		newSlice = append(newSlice, i)
		fmt.Printf("Lenght: %d - Capacity: %d\n", len(newSlice), cap(newSlice))
	}

	// Append slice to another slice
	// Note that eclipse ... which means we are
	// appending elements of the slice. Otherwise
	// it will try to append entier slice as the
	// next element. This will be rather slice of slices
	// or multidimentional slice
	newSlice = append(newSlice, mySlice...)
	fmt.Println(newSlice)
}