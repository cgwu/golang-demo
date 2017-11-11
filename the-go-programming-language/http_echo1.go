package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.path = %q\n", r.URL.Path)
}
