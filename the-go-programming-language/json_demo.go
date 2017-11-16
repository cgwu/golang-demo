package main

import (
	"fmt"
	"encoding/json"
	"log"
	"os"
)

type PointXY struct {
	X, Y int
}
type WheelEx struct {
	PointXY //anonymous fields
	Radius float32 `json:"半径"`
	Brand  string
}

func main() {
	w1 := new(WheelEx)
	w1.X = 1 //直接引用匿名域的属性
	w1.Y = 2
	w1.Radius = 3.14
	w1.Brand = "好牌子"
	fmt.Println(w1)
	//data, err := json.Marshal(w1) // 对象序列化成JSON
	data, err := json.MarshalIndent(w1, "", "    ") // 对象序列化成JSON,格式化输出
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", data)
	var w2 WheelEx
	if err := json.Unmarshal(data, &w2); err != nil { // 反序列化
		log.Fatalf("JSON unmarshaling failed: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(w2)
	fmt.Println(*w1 == w2)
}
