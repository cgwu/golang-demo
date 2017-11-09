package main

import "fmt"

func main() {
	var a = [3]int{1, 2, 3}
	//值复制
	var b = a
	b[1]++
	fmt.Println(a, b)	//直接打印数组
	//引用(地址)复制
	var c = &a
	c[1]++
	fmt.Println(a, c)
}
