package main

import (
	"net"
	"log"
	"os"
	"io"
)

func main() {
	conn, err := net.Dial("tcp", ":8001")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatalln(err)
	}
}
