package mymath

import (
	"errors"
	"fmt"
)

// 包初始函数，程序启动时自动调用
func init() {
	fmt.Println("1 #mymath package init func called.#")
}
// 可包含多个，按声明顺序自动调用
func init() {
	fmt.Println("2 #mymath package init func called.#")
}

func Add(a, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("参数应该非负!")
		return
	}
	return a + b, nil
}

func Subtract(a, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("参数应该非负!")
		return
	}
	return a - b, nil
}

func SayHi() {
	fmt.Println("Hello foo bar")
}
