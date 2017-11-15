package main

import "fmt"

type PointInt struct {
	X,Y int
}

type address struct {
	hostname string
	port int
}

func main() {
	p := PointInt{1,2}
	q := PointInt{1,2}
	fmt.Println(p == q)
	q.Y++
	fmt.Println(p == q)
	fmt.Println(q)

	// 结构体作为map的key
	hits := make(map[address]int)
	hits[address{"golang.org",443}]++
	hits[address{"golang.org",443}]++
	for k,v := range hits {
		fmt.Println(k,v)
	}
}
