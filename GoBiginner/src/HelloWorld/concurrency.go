/*
goroutinges vs OS Threads
-Go uses light weight construct get layerd on top of OS therad
-go routings are light weight (OS stack size may be some Megabytes,
 while go routing stack size is sigle digit kilobytes)
-Can dynamically grow
-Go manages goroutinges (creating to death)
-Less switching of OS threads because many go routings run on same thread
-Faster start-up times
-go routing easily communicate with each others, thanks to channels

Go's concurrency model:
-Actor model
-communicating Sequential Processes (CSP)
*/
package main

import (
	"fmt"
	"time"	
	"sync"		// For syncronization
	"runtime"	// To request more OS threads

)

func main() {
	
	// Anonymus funciton
	func() {
		time.Sleep(5*time.Second)
		fmt.Println("Hello")
	}()	// Placing the () at the end will self execute it

	func() {
		fmt.Println("Nethali")
	}()

	// Now on go routing
	// To request more OS threads (Virtual Processes)
	// Simple but dangoruos
	runtime.GOMAXPROCS(2)

	// Sync
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		defer waitGroup.Done()

		time.Sleep(5*time.Second)
		fmt.Println("Go")
	}()	// Placing the () at the end will self execute it

	go func() {
		defer waitGroup.Done()

		fmt.Println("Routing")
	}()

	waitGroup.Wait()
	/*
	You can have following code. But the problem is even
	main() exit unwinding the stack before  function execution
	So need to sync.
	go func() {
		time.Sleep(5*time.Second)
		fmt.Println("Go")
	}()	// Placing the () at the end will self execute it

	go func() {
		fmt.Println("Routing")
	}()
	*/
}