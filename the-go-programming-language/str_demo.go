package main

import (
	"fmt"
	"unicode/utf8"
	"bytes"
	"strconv"
)

const GoUsage = `Go is a tool for managing Go source code.

Usage:
	go command [arguments]
...`

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	//fmt.Println(GoUsage)
	s := "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	// 按每个字输出
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\t(%d)\n", i, r, size)
		i += size
	}
	// 等价于
	for i, r := range s {
		fmt.Printf("%d\t%c\n", i, r)
	}
	// 统计字数
	n := 0
	for range s {
		n++
	}
	fmt.Println("字符数:", n)
	// <==>
	fmt.Println(utf8.RuneCountInString(s))

	fmt.Println(string(65))
	fmt.Println(string(0x4eac))

	fmt.Println("-------------")
	fmt.Println(intsToString([]int{1, 2, 3, 4}))
	fmt.Println(intsToString([]int{}))

	fmt.Println("-------------")
	i123 := 123
	y123 := fmt.Sprintf("%d", i123)
	fmt.Println(y123, strconv.Itoa(i123))
	fmt.Println(y123 == strconv.Itoa(i123))
	fmt.Println(strconv.FormatInt(int64(i123), 2)) //转换进制
	fmt.Println(fmt.Sprintf("%b", i123) == strconv.FormatInt(int64(i123), 2))

	px123, _ := strconv.Atoi("123")
	py123, _ := strconv.ParseInt("123", 10, 64)
	fmt.Println(px123, py123)
}
