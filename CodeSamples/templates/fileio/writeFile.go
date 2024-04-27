package main

import (
	"fmt"
	"os"
//	"net/http"
)

func main() {

// ReadFile(name string) ([]byte, error)

	myFile := "json.data"

	myFileContents, err := os.ReadFile(myFile)
	if err != nil {
	}

	fmt.Println(string(myFileContents))

	err = os.WriteFile("output.file", myFileContents, 644)
	if err != nil {
	}
}

