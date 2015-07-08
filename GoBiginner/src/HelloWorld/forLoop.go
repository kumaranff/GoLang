package main

import (
	"fmt"
	"time"
)

func main() {
	// for loop keywords
	// blank expression is infinite loop
	// Can be boolean expression or range

	for timer := 10; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("Boom")
			break
		}
		fmt.Println(timer)
		time.Sleep(1*time.Second)
	}

	// for loop in range
	coursesInProgress := []string{"C++", "Java", "PHP"}
	
	for _, i := range coursesInProgress {
		fmt.Println(i)
	}

	// Label and break
	for i:=0; i<10; i++ {
		fmt.Println("i:", i)
//Label
breakPoint:
		for j:=0; j<10; j++ {
			fmt.Println("j", j)		
			if j == 3 {
				break breakPoint
			}
		}
	}

	// Infinite loop
	count := 0
	for {
		time.Sleep(1*time.Second)
		count++
		if count == 10 {
			break
		}
	}
}