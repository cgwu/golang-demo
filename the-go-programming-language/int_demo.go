package main

import (
	"fmt"
	"math"
)

func main() {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%08b\n", x&^y)

	var apples int32 = 1
	var oranges int16 = 2
	var compote int = int(apples) + int(oranges)
	fmt.Println(compote)

	f := 3.541 // a float64
	i := int(f)
	fmt.Println(f, i)

	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)
	x64 := int64(0xdeadbeef)
	//[1]意思为使用下标为1的参数值(此处即x64)，#号表示显示前缀0x或0X
	fmt.Printf("%d %[1]x %#[1]X\n", x64)

	// rune: n. 诗歌；古代北欧文字；神秘的记号
	// Runes are printed with %c, or with %q if quoting is desired.
	ascii := 'a'
	//unicode :='国'
	var unicode rune = '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]c %[1]q\n", newline)

	var f1 float32 = 16777216
	fmt.Println(f1)
	fmt.Println(f1 + 1)
	fmt.Println(f1 == f1+1)

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}
}
