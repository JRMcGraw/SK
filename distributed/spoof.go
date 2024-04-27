package main

import (
	"fmt"
	"io"
	"io/ioutil"
	//	"context"
	"net/http"
)

const keyServerAddr = "serverAddr"

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)
	mux.HandleFunc("/oldway", oldWay)

	mux.HandleFunc("/HomeDepot", HomeDepot)
	mux.HandleFunc("/EY", EY)
	mux.HandleFunc("/IBM", IBM)
	//mux.HandleFunc("/JSOC", JSOC)
	//mux.HandleFunc("/GE", GE)

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
/*
package main

import (
	"fmt"
	"io"
	//	"context"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

const keyServerAddr = "serverAddr"

func main() {


	mux := http.NewServeMux()

	// Create waitgroup


	wg.Add(9)

	// Launch APIs, each will Initialize

	mux.HandleFunc("/getRoot", getRoot)
	mux.HandleFunc("/HomeDepot", HomeDepot)
	mux.HandleFunc("/EY", EY)
	mux.HandleFunc("/IBM", IBM)
	mux.HandleFunc("/JSOC", JSOC)
	mux.HandleFunc("/GE", GE)
	mux.HandleFunc("/AAA", AAA)
	mux.HandleFunc("/Apple", Apple)
	mux.HandleFunc("/Disney", Disney)

	go http.ListenAndServe(":8090", mux)

	// Maybe we can replace this with another API (like the Master)
	// The others hit this endpoint to share their status
	// This could be "Here's something" (and we're now passing information
	// The payload could be an "I'm Done" message
	// Instead of using the browser to hit the endpoint, use the API

	//http.ListenAndServe(":8080", mux)

	fmt.Println("Waiting")

	// Wait for all to finish

	wg.Wait()

	fmt.Println("Everybody got their $hit done.")

}
*/

func HomeDepot(w http.ResponseWriter, r *http.Request) {

	Initialize("Home Depot")

	io.WriteString(w, "Home Depot here; I am at your service!\n")
	fmt.Println("You found me: HomeDepot")

	CleanUpShit("Home Depot")
}

func CleanUpShit(API string) {

	fmt.Println("I'm", API, "and I finished, now for Clean Up.")
	fmt.Println("One more cleaned up.")
//	wg.Done()
}

func Initialize(API string) {

	fmt.Println("I'm", API, "and this function() is doing basic init stuff.")
}

func EY(w http.ResponseWriter, r *http.Request) {

	message := "What can the EY API do for you?"
	//io.WriteString(w, "EY here; I am at your service!\n")
	fmt.Fprintf(w, message)
	fmt.Println("EY: I just sent a message to the Requester.")
	fmt.Println("EY: Requester URL:", r.URL)
	fmt.Println("EY: Requester Method:", r.Method )

//	wg.Done()
}


//

func IBM(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "IBM here; I am at your service!\n")
	fmt.Println("IBM: You found me.")

	response, _ := http.Get("http://localhost:8090/EY")

	fmt.Println("IBM: Got response from EY:", response.StatusCode)

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	message := string(bodyBytes)
	fmt.Println("IBM: The response from EY:", message)


	fmt.Println("IBM: Request URL:", r.URL)
	fmt.Println("IBM: Request Method:", r.Method )

	printResponse(response)
/*
	switch r.Method {
	case "GET": fmt.Println("Get something.")
	case "PUT": fmt.Println("Put something.")
	default: fmt.Println("Huh?", r.Method)
	}

	fmt.Println(r.URL)


	for i, v := range r.Header { fmt.Println(i, v) }

	fmt.Println("Body", r.Body)
	fmt.Println(r.TransferEncoding)
	fmt.Println(r.PostForm)
	fmt.Println("Response:", r.Response)

//	wg.Done()
*/

	fmt.Println("Did you see all the stuff I printed?")
}

/*

func JSOC(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "JSOC here; I am at your service!\n")
	fmt.Println("You found me: JSOC")
//	wg.Done()
}

func GE(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "GE here; I am at your service!\n")
	fmt.Println("You found me: GE")
//	wg.Done()
}



func AAA(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "AAA here; I am at your service!\n")
	fmt.Println("You found me: AAA")
	wg.Done()
}

func Apple(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "Apple here; I am at your service!\n")
	fmt.Println("You found me: Apple")
	wg.Done()
}

func Disney(w http.ResponseWriter, r *http.Request) {

	//
		Hello
	//
	io.WriteString(w, "Disney here; I am at your service!\n")

	fmt.Println("You found me: Disney")

	switch r.Method {
	case "GET": fmt.Println("Get something.")
	case "PUT": fmt.Println("Put something.")
	default: fmt.Println("Huh?", r.Method)
	}

	fmt.Println(r.URL)


	for i, v := range r.Header { fmt.Println(i, v) }

	fmt.Println("Body", r.Body)
	fmt.Println(r.TransferEncoding)
	fmt.Println(r.PostForm)
	fmt.Println("Response:", r.Response)

	fmt.Println("Disney here; I'm done.")

	wg.Done()
}

func getRoot(w http.ResponseWriter, r *http.Request) {

	//	ctx := r.Context()

	io.WriteString(w, "This is my website!\n")

	fmt.Println("What's this?  A root.")

	// fmt.Printf("%s: got / request\n", ctx.Value(keyServerAddr))
}

*/

func printResponse(response *http.Response) {

fmt.Println("Status:", response.Status)
fmt.Println("StatusCode:", response.StatusCode)
fmt.Println("Proto:", response.Proto)
fmt.Println("ProtoMajor:", response.ProtoMajor)
fmt.Println("ProtoMinor:", response.ProtoMinor)

// Header Header
for i, v := range response.Header { 
	fmt.Println("Header:", i, v) 
}

//Body io.ReadCloser

fmt.Println("ContentLength:", response.ContentLength)
fmt.Println("Close:", response.Close)
fmt.Println("Uncompressed:" , response.Uncompressed)

// Trailer Header,
fmt.Println("Trailer:", response.Trailer)

// Request *Request,
fmt.Println("Request:", *response.Request)

//fmt.Println("TLS:", *response.TLS)

}
