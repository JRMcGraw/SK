
package main

 import "fmt"
import "sort"


func main() {

	A := []int{1,2,3,4,5,7,8}

	res := Solution(A)

	fmt.Println(res)
}

var current  int

func Solution(A []int) int {

	sort.Ints(A)

	if len(A) > 0 { 

		fmt.Println("All negative")
		return 1
	}

	current = A[0]

	for _, next := range A[1:] {

		if next - current > 1 { break }
	}
	return current + 1
}
