package main

import "fmt"

type Stack struct {
	items	[]string
}

func ( s *Stack) Pop() string {

	length := len(s.items) - 1 
	popped := s.items[length]
	s.items= s.items[:length]
/*
	length := len(*s) - 1 
	popped := (*s)[length]
	(*s) = (*s)[:length]
*/

	return popped

}

func (s *Stack) Push(msg string) {

	s.items = append(s.items, msg)

	return

}

func main() {

	messages := []string{"Hello", "World", "Hi", "USA", "Stop", "Go"}

	var stacked Stack

	for _, m := range messages {

		stacked.Push(m)
	}

	fmt.Println(stacked)

	popped := stacked.Pop()
	fmt.Println(popped)

	fmt.Println(stacked)

	popped = stacked.Pop()
	fmt.Println(popped)

	fmt.Println(stacked)

	stacked.Push("ONE")
	stacked.Push("TWO")


	fmt.Println(stacked)


}

