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

func (h huge) hugeHello() {

	f.Println("huge:", h)
}

func (t tiny) tinyHello() {

	f.Println("tiny:", t)
}

type form struct {
	huge
	tiny
}

func main() {

	f.Println("Lazy man.")

	t1 := tiny{1,2}
	h1 := huge{5,6}

	both := form{huge{1,2},tiny{2,4}}

	f.Println("Tiny:", t1)
	f.Println("Huge:", h1)
	f.Println("Lazy man.")

	both.hugeHello()
	both.tinyHello()

	f.Println("Both:", both)

}
