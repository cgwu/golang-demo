package main

import "fmt"

func TestSwitch() {
	i := 23
	switch i {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("Default其它值")
	}
}

func TestFor() {
	sum := 0
	for {
		sum ++
		if sum >= 100 {
			break
		}
	}
	fmt.Println(sum)

	a := []int{1, 2, 3, 4, 5, 6}
	for i, j := 0, len(a)-1;
		i < j;
	i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	for _, v := range a {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func TestBreak() {
JLoop:
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i == 3 {
				goto DONE
			}
			if i > 5 {
				break JLoop
			}
			fmt.Println(j, i)
		}
	}
DONE:
	fmt.Println("done")
}

func main() {
	TestSwitch()
	TestFor()
	TestBreak()
}
