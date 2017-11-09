package main

import (
	"fmt"
	//"log"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}
func (p *Point) Print() {
	fmt.Println("(", p.x, ",", p.y, ")")
}

// 匿名组合（相当于继承）
type Rect struct {
	pt     Point
	width  float64
	height float64
	//*log.Logger	/* 匿名组合 */
}

func NewRect(x, y, width, height float64) *Rect {
	return &Rect{Point{x, y}, width, height}
}
func (r *Rect) Print() {
	r.pt.Print()
	fmt.Printf("Width:%g, Height:%g\n", r.width, r.height)
}

func main() {
	p1 := new(Point)
	p1.Print()
	p2 := &Point{100, 200}
	p2.Print()
	p3 := NewPoint(3, 4)
	p3.Print()

	r1 := &Rect{Point{1, 2}, 100, 200}
	r1.Print()
	r2 := NewRect(1.1, 2.2, 300.3, 400.4)
	r2.Print()
}
