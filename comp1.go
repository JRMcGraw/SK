package main

import (
	f "fmt"
)


type huge struct {
	top int
	bot int
}

type tiny struct {
	ins int
	outs int
}

func (h huge) hello() {

	f.Print(h.top)
}

func (t tiny) hello() {

	f.Print(t.ins)
}

type form struct {
	huge
	tiny
}

func main() {

	f.Println("Lazy man.")
/*
	t1 := tiny{1,2}
	h1 := huge{5,6}


	f.Println("Tiny:", t1)
	f.Println("Huge:", h1)
	f.Println("Lazy man.")
*/

}
