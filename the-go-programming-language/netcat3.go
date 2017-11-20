package main

import (
	"net"
	"log"
	"io"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8001")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("goroutine done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy3(conn, os.Stdin)
	conn.Close()
	<-done // wait for the background goroutine to finish
	log.Println("main done")
}

func mustCopy3(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
