package main

import (
	"net"
	"time"
	"fmt"
	"strings"
	"bufio"
	"log"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Println("接收到消息:", shout)
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
func handleConn3(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		//echo(c, input.Text(), 1*time.Second)
		go echo(c, input.Text(), 1*time.Second) // 一个Connection多个goroutine
	}
	// NOTE: ignoreing potential errors from input.Err()
	c.Close()
}
func main() {
	l, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn3(conn)
	}
}
