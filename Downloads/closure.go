package main

import (
	"fmt"
)

func main() {

	a := 1
	b := 2

	f, g := func(one, two int) (int, int)  {

		fmt.Println(a, b, a+b)
		fmt.Println(one, two, one+two)
		return (a+b), (one+two)

	} (3,4)

	fmt.Println(f, g)
	



}
