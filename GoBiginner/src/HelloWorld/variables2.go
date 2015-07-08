package main

import (
	"fmt"
)

func main() {
	name := "Nethali"	// Name of the person
	course := "GoLang"	// Name of the course

	fmt.Println("\nHi ", name, ". You are following ", course)

	ChageCourseNameByVal(course)

	fmt.Println("\nHi ", name, ". You are following ", course)

	ChageCourseNameByRef(&course)

	fmt.Println("\nHi ", name, ". You are following ", course)
}

// Return value can be named or unnamed.
// But naming it always help to the reader
func ChageCourseNameByVal(course string) string {
	// This is not initial assignment. Just use =
	course = "GoLangExtender"

	fmt.Println("\nCourse chaged to ", course)

	return course
}

func ChageCourseNameByRef(course *string) string {
	// This is not initial assignment. Just use =
	*course = "GoLangExtender"

	fmt.Println("\nCourse chaged to ", *course)

	return *course
}