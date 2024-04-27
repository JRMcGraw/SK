package main

import (
	"fmt"
	"sort"
)


func main() {

	list := []int{1,6,4,8,12,23,4}

	fmt.Println(list)

	sort.Ints(list)

	fmt.Println(list)

	previous :=  0

	for _, next := range list {

		if previous  == next { 
		
			fmt.Println("Duplicate Found", next) 
			return
		}
		previous = next
	}
	fmt.Println( "No Duplicates" )


}
