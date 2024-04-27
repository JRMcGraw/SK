
	package main

	import ( 
	
		"fmt"
		"strings"
		"errors"
	)


	var my_strings = []string{"Hello", "World", "Yes", "No"}
	var my_ints = []int{5, 4, 3, 2, 1}

	type my_stuff struct {

		name	string
		age	int
		fam	[]string
	}

	func main() {

		my_map := make(map[int]int)
		my_other_map := make(map[int]string)

		stuff := my_stuff{"Gobs", 5, []string{"Bonnie", "Kaden", "Kenna"}}


		for i := 0; i < len(my_ints); i++ {

			my_map[i] = my_ints[i]

		}

		for i, person := range my_strings {

			my_other_map[i] = person


		}

		for _, person := range stuff.fam {

			switch person {

			case "Bonnie" : fmt.Println("Hello, Bonnie McGraw")
					fmt.Println("I'm not done.")
			case "Kaden" : fmt.Println("Hello, ", person)
			case "Kenna" : fmt.Println("Hello, ", person, "Baby Girl.")

			}	
		}	

		if stuff.name == "Gobs" {

			fmt.Println(stuff.fam)

		} else {

			fmt.Println("Not Gobs")

		}


		fmt.Println(my_strings)
		fmt.Println(my_ints)
		fmt.Println(my_map)
		fmt.Println(my_other_map)
		fmt.Println(stuff)

		sentence := "The person pets a cat"

		words := strings.Split(sentence, " ")

		fmt.Println(words)

		result, err := check_sentence(words)


		if err != nil { 
		
			fmt.Println(err)
		}
		fmt.Println(result)
	}




/*


	<sentence> ::= <subject> <verb> <object>
	<subject> ::= <article> <noun>
	<object> ::= <article> <noun>
	<article> ::= the | a
	<noun> ::= dog | cat | person
	<verb> ::= pets | fed
*/

func check_sentence(remaining []string) (string, error) {

	s, remaining, _ := subject(remaining)
	v, remaining, _ := validate("verb", remaining)
	o, remaining, _ := object(remaining)

	complete := s + " " + v + " " + o

	fmt.Println(s, v, o, complete)

	return complete, nil
}

func subject(remaining []string) (string, []string, error) {

	a, remaining, _ := validate("article", remaining)
	n, remaining, _ := validate("noun", remaining)

	s := a + " " + n

	fmt.Println(a, n, s)

	return s, remaining, nil
}

func object(remaining []string) (string, []string, error) {

	a, remaining, _ := validate("article", remaining)
	n, remaining, _ := validate("noun", remaining)

	o := a + " " + n

	fmt.Println(a, n, o)

	return o, remaining, nil
}

func validate(which_part string, remaining []string) (string, []string, error) {

	token := remaining[0]
	remaining = remaining[1:]

	ok := test_part(which_part, token)

	if ok { return token, remaining, nil }

	return "", remaining, errors.New(which_part)

}
/*

*/

func test_part(which_part, token string) bool {


	ok := false

	switch which_part {

	case "verb": switch token {

		case "pets": ok = true
		case "fIed": ok = true

		}
	
	case "noun": switch token {

		case "dog": ok = true
		case "cat": ok = true
		case "person": ok = true

		}

	case "article": switch token {

		case "the": ok = true
		case "The": ok = true
		case "a": ok = true
		case "A": ok = true

		}

	}

	return ok

}
