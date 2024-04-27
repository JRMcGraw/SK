package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	http.HandleFunc("/", SimpleServer)


	http.ListenAndServe(":8080", nil)
}

func SimpleServer(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprintf(w, "Hello %s", r.URL.Path[1:])

	firstbody := body + files
	fmt.Fprintf(w, firstbody)
	read(w)
}


var body = ` <!DOCTYPE html>
<html>
  <head>
      <title>Server Hello World: Static Website</title>
      </head>
  <body>
      <h1>Happy Coding</h1>
      <img src="./b38.png" alt="My Image">`

var files = `<p>Hello world!</p>`


func read(w http.ResponseWriter) {

	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Fprintf(w, "Error:", err)
		return
	}
	for _, file := range files {

		fmt.Println(w, file.Name())


        }

}
