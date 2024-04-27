package main

import ( "net/http")

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", getHello)

	http.ListenAndServe(":8080", mux)
}

func getHello(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
        case "GET": fmt.Println("GET.")
        }
	fmt.Println(r.URL)
	for i, v := range r.Header {
		fmt.Println(i, v)
	}
	fmt.Println("Body", r.Body)
	fmt.Println("Response:", r.Response)

	io.WriteString(w, "Hello, HTTP!\n")
}
