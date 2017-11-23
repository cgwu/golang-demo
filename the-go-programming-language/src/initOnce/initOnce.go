package initOnce

import (
	"fmt"
	"sync"
)

var mapData map[int]string
var mapDataOnce sync.Once // 只初始一次

func GetString(key int) string  {
	mapDataOnce.Do(initData)
	return mapData[key]
}

func initData()  {
	mapData = make(map[int]string)
	mapData[1] = "one"
	mapData[2] = "two"
	mapData[3] = "three"
	fmt.Println("initData complete")
}