package main

import "fmt"

type MyStringer interface {
	MyString() string
}

type Base struct {
	msg string
}

func (b *Base) MyString() string {
	return b.msg
}

func main() {
	var b1 interface{} = &Base{"Hello中国"}
	//var b1 interface{} = 1

	//fmt.Println(b1.msg)
	if v, ok := b1.(MyStringer); ok {
		val := v.MyString()
		fmt.Println(val)
	} else {
		fmt.Println("未实现MyStringer接口")
	}
}
