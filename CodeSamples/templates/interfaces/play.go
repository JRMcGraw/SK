package main

import (
	"fmt"
)

func main() {

	allTypes := []interface{}{1, "hello", 1.55, []int{1,2,3}, "World", map[int]string{2:"two", 3:"three"},6}

	for i, v := range allTypes {

		fmt.Println(i,v)

		//t := v.(type)

		switch n := 1 ; v.(type) {
		case int: 
		
			num := v.(int)

			fmt.Println("I got an int:", v, n, num + n)

		case string: fmt.Println("I got a string:", v)
		case map[int]string: 

			//fmt.Println("I got a", t, v)
			fmt.Println("I got a", v)

			myMap, _ := v.(map[int]string)

			for index, value := range myMap {
				fmt.Println(index, value)
			}

		case []int: 
		
			fmt.Println("I got an int slice:", v)

			mySlice, _ := v.([]int)

			for i, value := range mySlice {
				fmt.Println(i, value)
			}

		case float64: fmt.Println("I got a float:", v)
		default: fmt.Println("I got squat.")
		}

	}

	fmt.Println(allTypes)
}
