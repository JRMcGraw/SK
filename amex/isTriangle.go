package main

import (
	"fmt"
	"sort"
)

type triple struct {
	p, q, r int
}

func main() {

	tripleSet := []triple{}

	set := []int{10, 2, 5, 1, 8, 20}

	sort.Ints(set)

	end := len(set) - 2

	size := len(set)

	p, q, r := 0, 0, 0

	oneTriple := triple{}

	for i, _ := range set[:end] {

		p = set[i]
		q = set[i+1]
		r = set[i+2]

		if size < p + q + r { continue }

		if ispqr(p, q, r ) == false { continue }
		if isqrp(p, q, r ) == false { continue }
		if isrpq(p, q, r ) == false { continue }

		oneTriple = triple{p,q,r}

		// Valid triple

		tripleSet = append(tripleSet, oneTriple)
	}

	fmt.Println("Triple Set:", tripleSet)


}

func ispqr(p, q, r int) bool { 
	return ( p + q ) > r 
}
func isqrp(p, q, r int) bool { 
	return ( q + r ) > p 
}
func isrpq(p, q, r int) bool { 
	return ( r + p ) > q 
}

