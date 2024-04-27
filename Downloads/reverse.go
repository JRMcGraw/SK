package main

import (
	f "fmt"
	s "strings"

)


var sentence = "The quick brouwn fox jumped over the  moon."

func main () {


	words := s.Split(sentence, " ")

	f.Println(words)

	for i, word := range words {

		f.Println(i, word)
	}

	end := len(words)

	for i := end - 1; i >= 0; i-- {


		f.Println(i, words[i])
	}

}
