package main

// Packages

import (
	"fmt"
	//"net/http"
	//"strings"
	//"database/sql"
	//"testing"
)

func main() {

	samples()

}

// Common statements, types

/*
map
struct
slice
*/

type bob struct {
	a int
	b string
	c map[int]string
}

func samples() {

	bobby := bob{}

	x := 5

	m := make(map[int]string)
	m = map[int]string{1: "Hi", 3: "Low", 6: "More"}

	this := m
	z := []int{1, 2, 3, 4, 5}

	bobby.a = x
	bobby.b = "Yes."
	bobby.c = this

	fmt.Println(bobby)

	fmt.Println(m)

	bigList := []int{}

	for _, b := range []int{1, 4, 2, 6, 4, 8} {
		switch b {
		case 1:
			//fmt.Println(b)
			bigList = append(bigList, b)
		case 2:
			//fmt.Println(b)
			bigList = append(bigList, b)
		default:
			bigList = append(bigList, b)
			//fmt.Println(b)
		}
	}

	fmt.Println(bigList)

	switch {
	case x == 65:
	case x > 3:
	default:
	}

	for _, stuff := range z {

		fmt.Println(stuff)
	}

	bobby.show(7)
}

func (s *bob) show(x int) {

	fmt.Println(s)

	num := s.a + x

	fmt.Println("The total is:", num)

	for i, item := range s.c {
		fmt.Println("show map:", i, item)
	}

}
