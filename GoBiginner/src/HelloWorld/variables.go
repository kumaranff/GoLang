/*
Variables in go is statically typed which means
do type checking (the process of verifying and 
enforcing the constraints of types) at compile-time 
as opposed to run-time. Dynamically typed programming 
languages do type checking at run-time as opposed to 
Compile-time.
*/
package main

import (
	"fmt"
	"reflect"
)

/*
Use var keyword if declaring at package level
ie: out side of function. Variables are set to 
intial zero values if not inisialized. Thye are
global variables.
*/
var(
	name string	// Name of subscriber
	course string // Name of current course
	module float64	// Current place in course

	// Go is smart enough to understand the types bu initialization values
	// Note that assignment is always use "=" sign
	name2, course2, module2 = "Nethali", "Go", 3.4
	test = 34
)

func main() {
	fmt.Println("Name is set to ", name, " and is type of ", reflect.TypeOf(name))
	fmt.Println("module is set to ", module, " and is type of ", reflect.TypeOf(module))
	fmt.Println("Name is set to ", name2, " and is type of ", reflect.TypeOf(name2))
	fmt.Println("module is set to ", module2, " and is type of ", reflect.TypeOf(module2))

	// This is local variables. Initial assignements are always use ":="
	a := 10.00000000
	b := 3

	fmt.Println("\nA is type ", reflect.TypeOf(a), " B is type of ", reflect.TypeOf(b))

	c := int(a)+b //c := a + b gives an error

	fmt.Println("\nC = ", c, " and is type of ", reflect.TypeOf(c))

	fmt.Println("\nThe memory address of C is ", &c)

	// Pointers can be assigned on the fly
	ptr := &c
	fmt.Println("\nThe memory address of C is ", ptr, " and the value of C is ", *ptr)
}