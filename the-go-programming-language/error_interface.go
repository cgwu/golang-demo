package main

import (
	"fmt"
	"errors"
	"syscall"
)
type errstr struct{
	text string
}

func main() {
	fmt.Println(errors.New("EOF") == errors.New("EOF"))
	fmt.Println(errstr{"EOF"} == errstr{"EOF"})
	fmt.Println(&errstr{"EOF"} == &errstr{"EOF"})

	e := fmt.Errorf("msg:%s,code:%d","错误信息",-1)
	fmt.Println(e.Error())
	var err error = syscall.Errno(2)
	fmt.Println(err.Error())
	fmt.Println(err)
}
