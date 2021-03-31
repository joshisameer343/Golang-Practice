package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("sameer")
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", handler2)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "welcome to home page")
}

func handler2(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "this is hello page")
}
