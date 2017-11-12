package main

import (
	"time"
	"fmt"
)

func main() {
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)

	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a, b, c, d)

	type Weekday int
	const (
		Sunday    Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Saturday, Friday, Sunday)

	const (
		_  = 1 << (10 * iota)
		KB  // 1<< 20
		MB  // 1<< 30
		GB
		TB
		PB
		EB
		ZB
		YB
	)
	fmt.Printf("KB=%#X\n", KB)
	fmt.Printf("MB=%#X\n", MB)
	fmt.Printf("GB=%#X\n", GB)
}
