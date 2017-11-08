package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	//var personDB map[string]PersonInfo
	//personDB = make(map[string]PersonInfo)
	personDB := make(map[string]PersonInfo)
	personDB["12345"] = PersonInfo{"12345", "Tom汤", "Room 203"}
	personDB["1"] = PersonInfo{"1", "Jack", "Room 1"}
	person, ok := personDB["12345"]
	if ok {
		fmt.Println("Found Person", person.Name, "with ID", person.ID)
	}else{
		fmt.Println("未找到")
	}
	fmt.Println("map1")
}
