// pakage main is included in all executables
// where main is required.
package main

// import format pakage
import (
	"fmt"
	"runtime"
)

// main is the entry point for (obviously) only
// executables. Take no arguments, no return values.
func main(){
	// runtime.GOOS gives us the running program OS
	fmt.Println("Hello from", runtime.GOOS)

}