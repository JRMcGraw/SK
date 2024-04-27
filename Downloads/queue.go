package main

import f "fmt"

type Queue struct {

	items []string
}


var queue Queue


func (q *Queue) Add (i string) {

	item := []string{i}

	q.items = append(item, q.items...)

	f.Println(q.items)

	return 

}


func (q *Queue) Remove () (string, error) {

	front := len(q.items) - 1

	first := q.items[front]

	q.items = q.items[:front]

	f.Println(q.items)

	return first, nil

}

func main () {


	people := []string{"One", "Tow", "Three"}

	for _, p := range people {

		queue.Add(p)
	}

	f.Println(queue.items)


	next, _ := queue.Remove()

	f.Println(next)



}
