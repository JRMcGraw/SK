package main

import (
	"fmt"
)


func main() {



	field := make(chan string)
	done := make(chan []string)

	go pitch(field, done)
	go toss(field, done)

	fmt.Println("Main, waitng.")

	one := <- done
	two := <- done

	fmt.Println("What did he say?", one)
	fmt.Println("What did he say?", two)
	fmt.Println("Main, done.")
}

func pitch(field chan string, done chan []string) {

	messages := []string{"Four", "and", "years" }

	sentence := []string{"Pitcher"}

	for i, msg := range messages {

		field <- msg
		answer := <- field

		sentence = append(sentence, msg)
		sentence = append(sentence, answer)

		fmt.Println("Pitched (",i, ")","Sent:", msg, "Got:", answer)

	}

	done <- sentence

	fmt.Println("Done. And I said:", sentence)
}

func toss(field chan string, done chan []string ) {

	messages := []string{"score","seven",  "ago"}

	sentence := []string{"Tosser"}

	for i, msg := range messages {

		answer := <- field
		field <- msg

		sentence = append(sentence, answer)
		sentence = append(sentence, msg)

		fmt.Println("Tossed (",i, ")","Got:", answer, "Sent:", msg)

	}

	done <- sentence

	fmt.Println("Done. And I said:", sentence)
}


