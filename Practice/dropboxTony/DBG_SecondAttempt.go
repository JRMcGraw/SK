package main

import (
        "fmt"
        "os"
        "strconv"
        "sync"
)

//
// The following program should be run via the following command:
// go run main.go 1 7 8 9
//
// where the numbers are the input to the program. For each number the program
// should output the factorial of the input in the order the input was given. For
// the above input, the output should be:
// 
// Factorial of 1 is: 1
// Factorial of 7 is: 5040
// Factorial of 8 is: 40320
// Factorial of 9 is: 362880
//
// The program should work for all positive integers. For unacceptable input, the
// program should give reasonable feedback to the user.
//

// Compute the factorial of `n`, print the result to `resultChan`, then
// signal to `wg` that the goroutine for this execution is complete
func factorial(order, n int, wg *sync.WaitGroup, resultChan chan<- results) {
	fmt.Println("In factorial.  Order:", order, "n:", n)
        defer wg.Done()

        result := 1
        for i := 2; i <= n; i++ {
                result *= i
        }

        //resultChan <- fmt.Sprintf("Factorial of %d is: %d", n, result)

	resultSet := results{order, n, result}

        resultChan <- resultSet
	fmt.Println("Done factorial.  Order:", order, "n:", n)
}

type results struct {
	order	int
	input	int
	fact	int
}

// Compute the factorial for one or more numbers and print the results.
//
// If no number is entered, print a message instructing how this tool should be used.
func main() {

	fmt.Println("DropBox starting.")

        if len(os.Args) < 2 {
                fmt.Println("Usage: go run main.go <number1> <number2> ...")
                os.Exit(1)
        }

	ordering := make(map[int]results)

        var wg sync.WaitGroup
        //resultChan := make(chan string, len(os.Args)-1)
        resultChan := make(chan results, len(os.Args)-1)

        for outputPosition, arg := range os.Args[1:] {
                num, err := strconv.Atoi(arg)
                if err != nil {
                        fmt.Printf("Error: %s is not a valid number\n", arg)
                        continue
                }
                if num < 0 {
                        fmt.Printf("Error: %s is not a positive number\n", arg)
                        continue
                }
		fmt.Println("DropBox next Input:", num, "Position:", outputPosition)

                wg.Add(1)
                go factorial(outputPosition, num, &wg, resultChan)
        }

//        go func() {
                wg.Wait()
                close(resultChan)
//        }()
	fmt.Println("DropBox ready to read Result Channel.")

        for result := range resultChan {

		ordering[result.order] = result
                fmt.Println(result)
        }

	fmt.Println("DropBox done reading channel.")
	fmt.Println("DropBox, reading", len(os.Args), "map entries.")

	for i := 0; i < len(os.Args)-1; i++ {

		n := ordering[i].input
		fact := ordering[i].fact

        	fmt.Println(fmt.Sprintf("Factorial of %d is: %d", n, fact))
	}
	fmt.Println("DropBox done.")

}

