// Clock1 is a TCP server that periodically writes the time.
/*
连接命令:
$ nc ip 8001

$ nc -h
Ncat 6.40 ( http://nmap.org/ncat )
Usage: ncat [options] [hostname] [port]

Options taking a time assume seconds. Append 'ms' for milliseconds,
's' for seconds, 'm' for minutes, or 'h' for hours (e.g. 500ms).
*/
package main

import (
	"net"
	"log"
	"io"
	"time"
)

func main() {
	//listener, err := net.Listen("tcp", "localhost:8001")
	listener, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		handleConn(conn)
	}
}
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("2006-01-02 15:04:05.000\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
