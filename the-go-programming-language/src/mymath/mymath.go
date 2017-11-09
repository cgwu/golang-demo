package mymath

import (
	"errors"
	"fmt"
)

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

func SayHi()  {
	fmt.Println("Hello foo bar")
}
