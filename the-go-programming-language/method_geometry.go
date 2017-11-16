package main

import (
	"math"
	"fmt"
)

type Pt struct {
	X, Y float64
}

// traditional function
func Distance(p, q Pt) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Pt type
func (p Pt) Distance(q Pt) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
func (p *Pt) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

type Path []Pt

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

type IntList struct{
	Value int
	Tail *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

func main() {
	p := Pt{1, 2}
	q := Pt{4, 6}
	fmt.Println(Distance(p, q)) // function call
	fmt.Println(p.Distance(q))  // method call
	path := Path{p, q, {7, 10}}
	fmt.Println("路径长度:", path.Distance())

	p.ScaleBy(3) // the compiler will perform an implicit &p on the var iable
	q.ScaleBy(3) // 编译器隐式转换为指针形式 q => (&q)
	fmt.Println("放大后距离:",p.Distance(q))

	n1 := IntList{Value:10}
	n2 := IntList{20,&n1}
	n3 := IntList{30,&n2}
	fmt.Println(n3.Sum())
	fmt.Println(n2.Sum())
	fmt.Println(n1.Sum())
}
