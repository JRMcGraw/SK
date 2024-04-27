package main

import (

	"fmt"
	"sort"
	"os"
	"time"
	"strconv"
	"strings"
	"math"
)

func main() {

	//problem1()
	//problem2()
	//problem3()
	data_types()
}

func data_types() {

	type complex_key struct {

		key1	int		// Street Number
		key2	string		// Strdet Name
		key3	int		// Apt. Number
		key4	string		// Zip

	}

	type kitchen_sink struct {

		cars	int
		bars	int
		trails	int
		
	}


	my_map := make(map[int]string)

//	my_structured_map := make(map[complex_key]kitchen_sink


	my_map = map[int]string{5:"five",4:"four"}

				
	fmt.Println(my_map)


	hi := "Hello World"

	hw_sl := strings.Split(hi,"l")

	fmt.Println( hw_sl )

	dice := math.Rand(6)

}

func problem3() {


	birth := os.Args[1]

	birthdate, _ := time.Parse("2006-01-02", birth)

	age_in := os.Args[2]

	age, _ := strconv.Atoi(age_in)



	fmt.Println("Birth Date:", birthdate)
	fmt.Println("Birth Year:", birthdate.Year())

	fmt.Println("Age:", age)

	future := birthdate.Year() + age

	fmt.Println("The year is", future)

}

func problem2() {

	cmd_line := []string{}

	os_cmd := os.Args[0]
	all_args := os.Args[1:]

	fmt.Println("Commend Line:", os.Args)
	fmt.Println("Commend:", os_cmd)
	fmt.Println("Args:", all_args)

	//for i := 0; i < len(all_args) ; i++ {
	for _, arg := range all_args {

		cmd_line = append(cmd_line, arg)
	}

	fmt.Println("Args Built:", cmd_line)
	fmt.Println("Command Line Built:", append([]string{os_cmd}, cmd_line...))

}

func problem1() {

	first := []string{"Ava", "Emma", "Olivia"} 
	second := []string{"Olivia", "Sophia", "Emma"}

	result := uniqueNames(first, second)

	fmt.Println(result)

}

func uniqueNames(a, b []string) []string {

	fmt.Println(a)
	fmt.Println(b)

	both := append (a, b...)

	fmt.Println(both)

    	sort.Strings(both)

	fmt.Println(both)

	sorted := []string{}

	previous := ""

	for _, s := range both {

		if s == previous { continue }

		previous = s

		sorted = append(sorted,s)
	}
	fmt.Println(sorted)

	return sorted

	
	}


	// should print Ava, Emma, Olivia, Sophia
	    
