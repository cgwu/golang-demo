package main

import "fmt"

//可变长度参数
func myfunc(args ...int) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}
func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "整数")
		case string:
			fmt.Println(arg, "字符串")
		case int64:
			fmt.Println(arg, "int64整数")
		default:
			fmt.Println(arg, "其它类型")
		}
	}
}
func main() {
	myfunc(1, 2, 3, 4)
	var v1 int = 1
	var v2 int64 = 64
	var v3 string = "Hello世界"
	var v4 float32 = 3.14

	MyPrintf(v1, v2, v3, v4)

	fmt.Println("-----------------")
	// 匿名函数
	f := func(x, y int) int {
		return x + y
	}
	fmt.Println(f(1,2))
	fmt.Println(f(11,22))
	func(msg string){
		fmt.Println("这是匿名函数调用", msg)
	}("消息")
}
