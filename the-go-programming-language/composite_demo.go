package main

import (
	"fmt"
)

func Zero(ptr *[3]int) {
	for i := range ptr {
		ptr[i] = 0
	}
}
func Zero2(ptr *[3]int) {
	*ptr = [3]int{}
}
func Zero3(ptr [3]int) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func Reverse(s []int) { // Slice作为参数
	for i, j := 0, len(s)-1;
		i < j;
	i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	var a1 [3]int = [3]int{11, 22, 33}
	a2 := [...]int{11, 22, 33}
	fmt.Printf("%T,%T\n", a1, a2)
	fmt.Println(a1, a2)
	fmt.Println("相等测试: ", a1 == a2)
	Zero(&a1)
	//Zero3(a1)	// 此操作并未清空main中的a1数组，说明数组也是"值传递"
	Zero2(&a2)
	fmt.Println("清空数组: ", a1, a2)

	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	// 此方法声明数组，下标可以间断
	symbol := [...]string{USD: "$", /* EUR: "E", */ GBP: "F", RMB: "Y"}
	fmt.Println(symbol)
	for i, v := range symbol {
		fmt.Printf("index: %d, value: %v\n", i, v)
	}

	r := [...]int{99: -1}
	fmt.Println(r)
	fmt.Println(len(r))

	fmt.Println(32*8, 256/4)

	arr1 := [...]int{0, 1, 2, 3, 4, 5}
	// Slice 作为参数可以在子函数中被修改
	Reverse(arr1[:])
	fmt.Println(arr1)
	Reverse(arr1[:2])
	Reverse(arr1[2:])
	fmt.Println(arr1)

	Reverse(nil) // safe

	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r) // built-in函数，append
	}
	fmt.Printf("%q\n", runes)
	runes2 := []rune("Hello, 世界") // 类型转换语法
	fmt.Printf("%q\n", runes2)

}
