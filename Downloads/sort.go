package main

import (
"fmt"
"sort"
)

type  flight struct {
	plane	string
	origin	string
	dest	string
	price	int
}

type  flights []flight


func (f flights) Less(i, j int) bool {

	return f[i].price > f[j].price
}

func (f flights) Len() int {

	return len(f)

}

func (f flights) Swap(i, j int) {

	f[i], f[j] = f[j], f[i]
	return
}

func main() {

	schedule := []flight{flight{"Delta","MCO", "BOS",400}, flight{"United","MCO", "SFO",100}, flight{"Southwest","MCO", "ATL",300}}

	fmt.Println(schedule)

	sort.Sort(flights(schedule))

	fmt.Println(schedule)

}
