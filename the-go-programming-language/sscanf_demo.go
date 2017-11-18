package main

import "fmt"

func main() {
	var i int
	var str string
	var f float64
	s := "123min分 3.1415926"
	//fmt.Sscan(s, &i, &str)
	fmt.Sscanf(s, "%d%s %f", &i, &str, &f)
	fmt.Println(i, str, f)
}
