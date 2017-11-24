package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer
	b.Write([]byte("Hello "))
	fmt.Fprintf(&b, "World!")
	b.WriteString("中国")
	b.WriteTo(os.Stdout)
}
