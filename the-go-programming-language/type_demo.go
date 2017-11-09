package main

import "fmt"

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) {
	*a += b
}

func main() {
	var a Integer = 3
	if a.Less(3) {
		fmt.Println(a, "小于4")
	} else {
		fmt.Println(a, ">= 4")
	}
	//var b Integer = 4
	a.Add(10) // 参数b:10 自动转化为Integer
	fmt.Println("a = ", a)
}
