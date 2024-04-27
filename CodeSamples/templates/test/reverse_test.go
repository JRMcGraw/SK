package main

import (
    "testing"
    "fmt"
)

func TestReverse(t *testing.T) {

    testcases := []struct {

    	in, want string
    
    }{ 
    	{"Hello there","ereht olleH"},
	{"HOW NOW","WON WOH"},
	{"12345","54321"},
    }

    for _, c := range testcases {

	reversed := Reverse(c.in)


	if reversed != c.want { t.Error() }
    }

    fmt.Println("Passed.")
}
