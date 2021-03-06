package main

import "fmt"

func f() *int {
	v := 1
	return &v
}
func incr(p *int) int {
	*p++
	return *p
}

func main() {
	x := 1
	var p *int = &x
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)

	var p1 = f()
	var p2 = f()
	fmt.Println(p1, p2)
	fmt.Println(*p1, *p2)
	*p1++ // increments what p points to; does not change p
	fmt.Println(*p1, *p2)

	v := 123
	incr(&v)
	fmt.Println(v)

	fmt.Println("测试new(int)")
	pn := new(int)	// p, of type *int, points to an unnamed int variable
	fmt.Println(*pn)
	*pn = 2
	fmt.Println(*pn)

	var i1 int = 1
	fmt.Println(i1)
	//i1 = nil	//cannot use nil as type int in assignment


}
