package main

import (
	"fmt"
	"slices"
)

type example struct {

	name	string
	age	int
	state	string
	degree	[]string
	//lived	map[string]int
	male	bool
}

func (e example) isAdult() bool {

	if e.age >= 18 {return true}
	return false
}

func (e example) sortDegrees() bool {

	slices.Sort(e.degree)
}


func main() {


	//bob := example{"Bob", 13, "FL", []string{"BS", "MSCS","MS"},{map["NC"]1,map["TX"]2},true}
	bob := example{"Bob", 13, "FL", []string{"MS", "PhD","BA","AA", "BS"},true}

	fmt.Println(bob)

	fmt.Println("Bob is an Adult:", bob.isAdult())

	fmt.Println(bob.sortDegrees())
}
