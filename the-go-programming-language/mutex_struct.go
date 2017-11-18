package main

import (
	"sync"
	"fmt"
)

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string), // ,号不可少：syntax error: unexpected newline, expecting comma or }
}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

func main() {
	cache.mapping["abc"] = "Hello中国"
	fmt.Println(Lookup("abc"))

}
