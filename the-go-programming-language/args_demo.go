package main

import (
	"os"
	"fmt"
	"strings"
)

func PrintArgs1() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
}
func PrintArgs2() {
	for i, v := range os.Args {
		fmt.Println(i, v)
	}
}

func PrintArgs3()  {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	PrintArgs2()
	PrintArgs3()
}
