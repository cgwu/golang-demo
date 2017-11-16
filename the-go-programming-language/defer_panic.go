package main

import (
	"fmt"
	"runtime"
	"os"
)

func fp(x int) {
	//if x == 1 {
	//	return
	//}
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	fp(x - 1)
}
func fpWrapp(x int) (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("出现panic: %v", p)
		}
	}()
	fp(x)
	return nil
}
func printStack() {
	fmt.Printf("------------------")
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
func main() {
	//defer printStack()
	//fp(3)
	if err := fpWrapp(3); err != nil {
		fmt.Println(err)
	}
}
