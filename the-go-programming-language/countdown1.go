package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second) // 返回<-chan time, 按指定时间间隔自动产生时间数据.
	for countdown := 10; countdown > 0; countdown -- {
		fmt.Println(countdown)
		<-tick
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
