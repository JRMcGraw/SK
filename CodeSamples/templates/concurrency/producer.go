package main

import (
	"fmt"
)

type payload struct {
	entity		string		// Tank, Soldier, Animal
	instance	int		// How many have I seen
	source		string		// Drong ID
}

type message struct {
	ID		int
}

func main () {

	topic := make(chan int, 25)
	status := make(chan int, 10))

	Consumers := 40

//	Launch Consumers

	indentSizes := []string{"", "   ", "      ", "         ", "            "}

	whichIndent := 0

	for c := 1; c < consumerCount; c++ {

		whichSize++
		indentSize := indentSizes[whichIndent]

		go consumr(c, topic, status, indentSize)

		if whichSize == len(indents) { whichSize = 0 }
	}

	msg := 1

	for {
		topic <- msg++

		if msg > 100 { break }
	}

func consumer(id int, topic, status chan int, indent string {

	msgCount := 0

	time.Sleep(time.Second)

	msg := <- topic

	msgCount++

		fmt.Sprintf( "%s (%n) [%n]\n",indent,  id, msgCount)

}





}
