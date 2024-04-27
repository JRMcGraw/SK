package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	http.ListenAndServe(":8080", nil)

}

func getRoot(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "This is my website!\n")
	fmt.Println("What's this?  A root.")

}

func getHello(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Hi there.")
	io.WriteString(w, "Hello, HTTP!\n")
}

