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
      <img src="./b38.png" alt="My Image">`

var files = `<p>Hello world!</p>`
var crlf = `<CR><LF>`

var bodyend = `</body>
    </html>`



func read(w http.ResponseWriter) {

	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Fprintf(w, "Error:", err)
		return
	}
	for _, file := range files {

		//filename := file.Name() + '\r' + '\n'
		//filename := []byte(file.Name()) + '\r' + '\n'
		//fmt.Fprintf(w, filename)
		//filename := []byte(file.Name()) + '\r' + '\n'

		fmt.Fprintf(w, file.Name())

		//filename := `<p>` + []byte(file.Name()) + `<p>`
		//files = files + filename
		fmt.Fprintf(w, crlf)

        }

	//totalbody := body + files + bodyend
	//fmt.Fprintf(w, totalbody)
}
