package main

import (
	"net/http"
	"log"
	"fmt"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	/*
	Behind the scenes, the ser ver runs
	the handler for each incoming request in a separate goroutine
	so that it can ser ve multiple
	requests simultane ously.
	*/
	http.HandleFunc("/url", handlerURL) // Pattern "/" 会匹配所有路径
	http.HandleFunc("/req", handlerRequest)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

// handler echoes the HTTP request.
func handlerRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func handlerURL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle url")
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.path = %q\n", r.URL.Path)
}
func counter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("counter")
	mu.Lock()
	fmt.Fprintf(w, "Count: %d\n", count)
	mu.Unlock()
}
