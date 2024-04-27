package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {

    r := mux.NewRouter()

    bookHandler := func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        title := vars["title"]
	page := vars["page"]

        msg := fmt.Sprintf("You've requested the book: %s on page %s\n", title, page) 
        fmt.Fprintf(w, msg)
        fmt.Println(msg)
    }
/*
*/
    r.HandleFunc("/books/{title}/page/{page}", bookHandler)

    http.ListenAndServe(":80", r)
}
