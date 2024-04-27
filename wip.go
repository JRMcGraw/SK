package main

import (
	f "fmt"
	//p1 "pkg1"
	//p2 "pkg2"

	p1 "github.com/KitchenSink/pkg1"
	p2 "github.com/KitchenSink/pkg2"
	sy "github.com/KitchenSink/syntax"
	pr "github.com/KitchenSink/practicing"
	ch "github.com/KitchenSink/channels"
	dt "github.com/KitchenSink/datatypes"
	t "github.com/KitchenSink/interfaces"
	"github.com/KitchenSink/files"
)


func main() {


	f.Println("Packages, packages, packages.")

	p1.Hello()
	p2.SayHello()
	p2.SayThere()
	p2.SayWorld()

	sentence := "The quick brouwn fox jumped over the  moon."
	sy.Reverse(sentence)

	l1 := []string{"Ava", "Emma", "Olivia"}
	l2 := []string{"Olivia", "Sophia", "Emma"}

	res := pr.UniqueNames(l1, l2)
	f.Println(res)


	ch.Channels()

	t.Travel()


	files.Reading("README.md")

	dt.Types()
}
