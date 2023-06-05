package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Your code here

	 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	    fmt.Fprintln(w, "Hello, World!")
	})

	// Additional handlers...

	http.ListenAndServe(":8080", nil)
}