package main

import (

	"fmt"
	"time"
	"strconv"
)

var global_counter = 1
var global_Name = "First Name"
var global_wait = 1

var endings = []int{24,18,15,40}

func main() {

	fmt.Println(global_Name, "Started.")
	fmt.Println(endings)

	start := 0

	increment := 3

	for _, end := range endings {

		start = start + increment

		my_name := "Worker-" + strconv.Itoa(end)

		fmt.Println("Launching:", my_name, "Start, End:", start, end)


		go collide_with_global_variable (my_name, start, end) 

	}
	time.Sleep(10 * time.Second)

	fmt.Println("Worker:", global_Name, " aka,", "Main", "Ended.")
}

func collide_with_global_variable (my_name string, start, end int) {

	fmt.Println("Who am I:", my_name, "Started.  Start, End:", start, end)


	current := start


	for i:= 1; current < end; current++ {
		
		i = i

		global_counter++
		time.Sleep(time.Duration(global_wait) * time.Millisecond)

		global_wait++
		fmt.Println("I am:", my_name, "Global Name:", global_Name, "Wait:", global_wait, "Counter:", global_counter)

		time.Sleep(time.Duration(global_wait) * time.Millisecond)
		global_Name = my_name

		global_counter++
		global_wait++
	}

	fmt.Println("Worker:", my_name, "is now", global_Name, "and Ended @ ", global_counter)
}
