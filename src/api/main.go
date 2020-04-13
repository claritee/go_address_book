package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting app")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
