package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, `{"message":"pong"}`)
	})

	fmt.Println("start http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
