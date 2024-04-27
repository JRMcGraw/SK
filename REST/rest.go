package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	http.HandleFunc("/", Base)
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/bye", Bye)
	http.HandleFunc("/nice", Nice)


	http.ListenAndServe(":8090", nil)
}

func Base(w http.ResponseWriter, r *http.Request) {

	MoreWork(w,r)

}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
	fmt.Println("Hello")
}

func Bye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bye")
	fmt.Println("Bye")
}

func Nice(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Nice")
	fmt.Println("Nice")
}

func MoreWork(w http.ResponseWriter, r *http.Request) {

	msg := fmt.Sprintf("Hello %s", r.URL.Path[1:])
	fmt.Println(msg)
	//fmt.Fprintf(w, msg)

	firstbody := body + files + bodyend
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
      <p>Next Line</p>
      <p>Line 2</p>
      <p>Line 3</p>
      <img src="img015.jpg" alt="My Image">`

var files = `<p>Hello world!</p>`
var crlf = `<CR><LF>`

var bodyend = `</body>
    </html>`

var newLineMsg = ` <!DOCTYPE html>
<html>
  <body>
      File:`


func read(w http.ResponseWriter) {

	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Fprintf(w, "Error:", err)
		return
	}
	for _, file := range files {

		fmt.Fprintf(w, file.Name())
		fmt.Fprintf(w, newLineMsg)
		fmt.Println(file.Name())

        }

}
