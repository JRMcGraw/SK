package main

import (
	"fmt"
	"playing"
)


func main() {


	fmt.Println("Lazy man.")
	fmt.Println(playing.Playing())
	//fmt.Println(pl.Playing())

}

func globals() {

	fmt.Println("Globally:",playing.Playing())
}
