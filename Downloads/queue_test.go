package main

import (

	"testing"
)

var people_in []string{"One", "Two", "Three"}
var queue_order []string{"Three", "Two", "One"}


func TestRemove () {

	for _, person_removed := range people_in {

		removed := Remove()

		if removed != person_removed {
			t.Error()
		}
	}

}

func TestAdd(t *testing.T) {


	for _, p := range people_in {

		Add(p)
	}

	for i, p := range people_order {

		if queue.items[i] != p {

			t.Error()
		}
	}

}

