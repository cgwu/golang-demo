package main

import (
	"fmt"
	//"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 50 ; x++ {
			naturals <- x
		}
		close(naturals)
	}()
	// Squarer
	go func(){
		for {
			x, ok := <- naturals
			if !ok {
				break // channel was closed and drained
			}
			squares <- x * x
		}
		close(squares)
	}()
	// Printer (in main goroutine)
	/*法1*/
	//for {
	//	fmt.Println(<-squares)
	//	time.Sleep(500 * time.Millisecond)
	//}

	/*法2*/
	for x := range squares {
		fmt.Println(x)
		//time.Sleep(10 * time.Millisecond)
	}
}
