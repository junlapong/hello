package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!\n")
	})

	fmt.Println("start http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
