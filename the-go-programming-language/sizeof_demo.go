package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x int
	var x32 int32
	var y int64
	fmt.Println(unsafe.Sizeof(x32), unsafe.Sizeof(x), unsafe.Sizeof(y))
}
