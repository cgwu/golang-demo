package main

import "fmt"

func counter3(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}
func squarer3(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}
func printer3(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
	//close(in)	// no need
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter3(naturals)
	go squarer3(squares, naturals)
	printer3(squares)
}
