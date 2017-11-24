package main

import (
	"bzip"
	"os"
	"fmt"
	"io"
	"log"
)

func mustNil(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	//var compressed, uncompressed bytes.Buffer
	compressed,err := os.Create("compressed.txt.bz2")
	mustNil(err)
	defer compressed.Close()

	uncompressed,err := os.Create("uncompressed.txt")
	mustNil(err)
	defer uncompressed.Close()
	w := bzip.NewWriter(compressed)

	// Write a repetitive message in a million pieces,
	// compressing one copy but not the other.
	tee := io.MultiWriter(w, uncompressed)
	for i := 0; i < 1000000; i++ {
		io.WriteString(tee, "hello")
	}
	if err := w.Close(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("main done")
}
