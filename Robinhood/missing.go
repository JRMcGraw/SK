
package main



// you can also use imports, for example:
 import "fmt"
// import "os"
//import "sort"


func main() {

	A := []int{1,2,3,4,5,7,8}

	res := Solution(A)

	fmt.Println(res)
// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")
}

func Solution(A []int) int {

//	sort.Ints(A)

	end := len(A)

	if A[end] < 0 { fmt.Println("All negative") }

	for i, _ := range A {

		fmt.Println(A[i], A[i+1])

	}
	return 0
}
    // Implement your solution here
