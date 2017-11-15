package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	// <=>
	ages2 := make(map[string]int)
	ages2["alice"] = 31
	ages2["charlie"] = 34

	fmt.Println(ages, ages2)
	delete(ages2, "alice")
	fmt.Println(ages2)
	ages2["charlie"]++
	fmt.Println(ages2)
	delete(ages2, "not_exit") // safe

	ages["zella"] = 10
	ages["dannis"] = 29

	fmt.Println("--------map顺序--------")
	//var names []string
	names := make([]string, 0, len(ages)) // allocates a slice of length 0 and capacity 10.
	//for name := range ages {	// 也可只循环key值
	for name, age := range ages {
		names = append(names, name)
		fmt.Println(name, age)
	}
	sort.Strings(names)
	fmt.Println("-------按名字排序输出---------")
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
	searchName := "dannisx"
	if search_age, ok := ages[searchName]; !ok {
		fmt.Println(searchName + "不存在")
	} else {
		fmt.Printf("%s age:%d", searchName, search_age)
	}

}
