package main
/*
go build sleep.go
./go -period 1s
./go -period 50ms
./go -period 2m30s
./go -period 1.5h
*/
import (
	"flag"
	"time"
	"fmt"
)

var period  = flag.Duration("period", 1*time.Second, "sleep period")

func main() {
	flag.Parse()
	fmt.Printf("Sleep for %v...", *period)
	time.Sleep(*period)
	fmt.Println()

}
