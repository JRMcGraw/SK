package main

import (
	"fmt"
)


func main() {


	mux := http.NewServer()

	// Create waitgroup

	wg := ?

	// Launch APIs, each will Initialize

	go launch(mux, "HomeDepot")
	go launch(mux, "EY")
	go launch(mux, "IBM")
	go launch(mux, "JSOC")
	go launch(mux, "GE")
	go launch(mux, "AAA")
	go launch(mux, "Apple")
	go launch(mux, "Disney")

	// Wait for all to finish


	fmt.Println("Lazy man.")

}

func launch(httpServer *type, API string) {

	endpoint := "/" + API

	what := httpServer.HandleFunc(endPoint, serveEndpoint())

	huh := httpServer.ListenAndServe(":8080")
}

func serveEndpoint() {


}
