package main

import (
	"io"
	"os"
	//"bytes"
	"fmt"
	"bytes"
)

func main() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File) // 1. *os.File is Concrete Type
	fmt.Printf("%T\n",f)
	//c := w.(*bytes.Buffer)// panic: interface conversion: io.Writer is *os.File, not *bytes.Buffer
	//fmt.Printf("%v\n",c)

	rw := w.(io.ReadWriter) // 2. io.ReadWriter
	fmt.Printf("%T\n",rw)

	// test
	if _,ok := w.(*os.File); ok {
		fmt.Println("w is *os.File")
	} else {
		fmt.Println("w is NOT *os.File")
	}
	if _,ok := w.(*bytes.Buffer); ok {
		fmt.Println("w is *bytes.Buffer")
	} else {
		fmt.Println("w is NOT *bytes.Buffer")
	}

}
