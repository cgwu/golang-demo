/*
ref:
https://studygolang.com/articles/2877
https://studygolang.com/articles/2073

*/
package ioutil_demo_test

import (
	"io/ioutil"
	"testing"
	"os"
	"bufio"
	"fmt"
)

func mustNil(e error) {
	if e != nil {
		panic(e)
	}
}
func TestIoutilWrite(t *testing.T) {
	err := ioutil.WriteFile("dat1.txt", []byte("你好Hello\ngo\n"), 0644)
	mustNil(err)
}
func TestIoutilRead(t *testing.T) {
	data, err := ioutil.ReadFile("dat1.txt")
	mustNil(err)
	fmt.Println(string(data))
}

func TestOsCreate(t *testing.T) {
	f, err := os.Create("dat2.txt")
	mustNil(err)
	defer f.Close()
	f.WriteString("Hello\n中国\n")
	f.Sync() //调用 Sync 来将缓冲区的信息写入磁盘。

	//bufio 提供了和我们前面看到的带缓冲的读取器一样的带缓冲的写入器。
	w := bufio.NewWriter(f)
	w.Write([]byte("buffered缓冲的writer\n"))
	w.Flush()
}
func TestOsRead(t *testing.T) {
	f, err := os.Open("dat2.txt")
	mustNil(err)
	defer f.Close()

	b3 := make([]byte,40)
	n, err := f.Read(b3)
	if err == nil {
		fmt.Println(n, string(b3))
	}
}