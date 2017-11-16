package main

import "fmt"

type PointInt struct {
	X, Y int
}
type Wheel struct {
	PointInt //anonymous fields
	Radius float32
}
type address struct {
	hostname string
	port     int
}

func main() {
	p := PointInt{1, 2}
	q := PointInt{1, 2}
	fmt.Println(p == q)
	q.Y++
	fmt.Println(p == q)
	fmt.Println(q)

	// 结构体作为map的key
	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
	hits[address{"golang.org", 443}]++
	for k, v := range hits {
		fmt.Println(k, v)
	}

	w1 := new(Wheel)
	w1.X = 1 //直接引用匿名域的属性
	w1.Y = 2
	w1.Radius = 3.14
	fmt.Println(w1)

}
