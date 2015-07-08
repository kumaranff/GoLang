package main

import (
	"fmt"
	"os"
)

func main() {
	// In go language, always the last return value is error status
	// The following function return two values, we ignore the first with _
	_, err := os.Open("c:\\temp\\test.txt")

	if err != nil{
		fmt.Println("Error occured: ", err)
	}

	if fileFound, err := IsFileAvailable("c:\\temp\\test1.txt"); fileFound {
		fmt.Println("File Found")
	} else {
		fmt.Println("File not found. Reason :", err)
	}

}


// Good way to write produection code with the error status
// return types are bool and error
func IsFileAvailable(path string) (bool, error) {
	_, err := os.Open(path)

	fileFound := false;
	if err == nil{
		fileFound = true;
	}
	return fileFound, err
}