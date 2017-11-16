package main

import (
	"math"
	"fmt"
)

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func zero(int, int) int { return 0 }

func square(n int) int   { return n * n }
func negative(n int) int { return -n }

// 产生闭包
func squares() (func() int) {
	var x int
	return func() int {
		x++
		return x * x
	}
}
// 参数个数可变
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func main() {
	fmt.Println(hypot(3, 4))
	fmt.Printf("%T\n", zero)

	var f func(int) int
	//f(3)	//panic: call of nil function.错误信息: runtime error: invalid memory address or nil pointer dereference
	f = square
	fmt.Println(f(3))
	f = negative
	fmt.Println(f(3))

	fmt.Println("--闭包测试--")
	fc := squares()
	fmt.Println(fc())
	fmt.Println(fc())
	fmt.Println(fc())
	fmt.Println(fc())
	fmt.Println("--参数个数可变测试--")
	fmt.Println(sum())
	fmt.Println(sum(1,2))
	fmt.Println(sum(1,2,3,4,5))
	values := []int{10,20,30}
	fmt.Println(sum(values...))	// slice 打散作为参数

	fmt.Printf("%T\n",fv)
	fmt.Printf("%T\n",fs)

}
func fv(...int){}
func fs([]int){}
