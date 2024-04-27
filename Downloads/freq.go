package main

import (
"fmt"
)

func main() {


	linked := make(map[string]int)

	set := []string{"A","B","C","D","A","C","C","D","F","H","X","F","A","B","A"}

	fmt.Println(set)

	for _, n := range set {

		f := linked[n]
		if f == 0 { linked[n] = 0 }

		f++

		linked[n] = f

	}

	fmt.Println(linked)

}
