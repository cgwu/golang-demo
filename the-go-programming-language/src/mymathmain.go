package main

import "mymath"
import "fmt"

func main() {
	c, ok := mymath.Add(11, -2)
	if ok == nil {
		fmt.Println(c)
	} else {
		fmt.Println("nil result:" + ok.Error())
	}
	mymath.SayHi()
	c1, _ := mymath.Subtract(1,2)
	fmt.Println(c1)
}
