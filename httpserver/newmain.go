package main

import (
	"fmt"
	"io"
	//	"context"
	"net/http"
)

const keyServerAddr = "serverAddr"

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)
	mux.HandleFunc("/oldway", oldWay)

	go http.ListenAndServe(":8090", mux)
	http.ListenAndServe(":8080", mux)

}

func getRoot(w http.ResponseWriter, r *http.Request) {

	//	ctx := r.Context()

	io.WriteString(w, "This is my website!\n")

	fmt.Println("What's this?  A root.")

	// fmt.Printf("%s: got / request\n", ctx.Value(keyServerAddr))
}

func getHello(w http.ResponseWriter, r *http.Request) {

	/*
		Hello
	*/
	fmt.Println("Hi there.")
	/*	Hello
	 */

	switch r.Method {
	case "GET": fmt.Println("Get something.")
	case "PUT": fmt.Println("Put something.")
	default: fmt.Println("Huh?", r.Method)
	}

	// Hello
	fmt.Println(r.URL)

	//	Hello

	for i, v := range r.Header {
		// Hello
		fmt.Println(i, v)
	}
	// TheEnd

	fmt.Println("Body", r.Body)
	fmt.Println(r.TransferEncoding)
	fmt.Println(r.PostForm)
	// Help
	fmt.Println("Response:", r.Response)

	io.WriteString(w, "Hello, HTTP!\n")
}

func oldWay(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Old Way; it always worked.")
	fmt.Println("Old Way; it always worked.")

}
