package main

import "fmt"

var (
	g1 = 3.14
	g2 = "Hello世界"
)

func TestVariable() {
	var v1 int = 10
	var v2 = 20
	v3 := 30
	fmt.Printf("v1=%d, v2=%d, v3=%d\n", v1, v2, v3)
	v1, v3 = v3, v1
	fmt.Printf("v1=%d, v2=%d, v3=%d\n", v1, v2, v3)
	fmt.Printf("g1=%g, g2=%s，字符串长度:%d\n", g1, g2, len(g2))
}

func GetName() (firstName, lastName, nickName string) {
	//func GetName() (string,string, string) {
	return "May", "Chan", "最后名"
}

func TestString() {
	str := "Hello, 世界"
	n := len(str)
	for i := 0; i < n; i++ {
		ch := str[i]
		fmt.Println(i, ch)
	}
}
func TestStringUnicode() {
	str := "Hello, 世界"
	for i, ch := range str {
		fmt.Println(i, ch)
	}
}
func TestArray() {
	array := [5]int{11, 22, 33, 44, 55}
	for i, v := range array {
		fmt.Printf("Array element[%d]=%d\n", i, v)
	}
}

func TestSlice() {
	mySlice := make([]int, 3, 5)
	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))

	mySlice = append(mySlice,1,2,3)
	fmt.Println("len(mySlice):", len(mySlice))
	fmt.Println("cap(mySlice):", cap(mySlice))
	for i, v := range mySlice {
		fmt.Println(i,v)
	}
}

func main() {
	var lastName string
	TestVariable()
	_, _, lastName = GetName()
	fmt.Println(lastName)
	TestString()
	fmt.Println("----------")
	TestStringUnicode()
	fmt.Println("----------")
	TestArray()
	fmt.Println("----------")
	TestSlice()
}
