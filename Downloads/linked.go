package main

import (

	f "fmt"
)

type item struct {
	ID		int
	next		int
	position	int
	payload		content
}

type linked_list []item

var list_end int

type content struct {
	info string
}

func (ll linked_list) insert_item(record item) {

	return
}

func (ll linked_list) get_next() item, error {

	if ll.next == 0 { return item{}, nil }

	return linked_list[ll.next], nil
}

func (ll linked_list) get_item(ID string) item, error {

	postion := Nodes[ID]

	return ll[postion], nil

}

var entries []string{1,2,3,4,5,6,7,8,9,10,11}


func main() {

	list_end = 1

	Nodes := make(map[int]int)

	for _, e := range entries {


	}
	f.Println()
}
