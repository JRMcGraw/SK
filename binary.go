package main

import (
	f "fmt"
	e "errors"
)


func main() {

	sample := []int{3,8,9,18,22,27,35,48,55,77}

	key := 9

	match := false
	next := sample
	mid := len(sample)/2

	for {

		f.Println( key, mid, next, match )

		match, next, _ = pick( key, mid, next )

		mid = len(next)/2

		f.Println( key, mid, next, match )

		if match { break }

		if len(next) <= 1 { break }
	}

	f.Println("Matched:", key, match)

}

func pick ( key, mid int, sample []int ) (bool, []int, error) {

	f.Println( key, mid, sample )

	test := sample[mid] 

	if  key == test { return true, nil, nil }
	if  key > test { return false, sample[mid:], nil }
	if  key < test { return false, sample[:mid], nil }

	return false, nil, e.New("WTH")
}
