package main

import (
	"time"
	"log"
	"fmt"
)
// defer很巧妙的用法
func bigSlowOperation() {
	log.Println("bigSlowOperation start")
	defer trace("bigSlowOperation")()
	time.Sleep(5 * time.Second)
	log.Println("bigSlowOperation end")
}
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

// defer 打印返回值
func double(x int) (result int) {
	if x == 0 {
		panic("x不能为0")	 // 仅用于测试panic
	}
	defer func(){
		fmt.Printf("double(%d) = %d\n",x,result)
	}()
	return x + x
}

func main() {
	//bigSlowOperation()
	 double(3)
	 double(0)
	log.Println("main done")
}
