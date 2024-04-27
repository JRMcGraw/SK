package main

import (
        "fmt"
        "os"
        "strconv"
        "sync"
)

func rfactorial(n int) int {
  if n == 0 {
    return 1
  }
  return n * rfactorial(n-1)
}


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
func factorial(n int, wg *sync.WaitGroup, resultChan chan<- string) {
        defer wg.Done()

        result := 1
        for i := 2; i <= n; i++ {
                result *= i
        }

        resultChan <- fmt.Sprintf("Factorial of %d is: %d", n, result)
}

// Compute the factorial for one or more numbers and print the results.
//
// If no number is entered, print a message instructing how this tool should be used.
func main() {
        if len(os.Args) < 2 {
                fmt.Println("Usage: go run main.go <number1> <number2> ...")
                os.Exit(1)
        }

        var wg sync.WaitGroup
        resultChan := make(chan string, len(os.Args)-1)

        for _, arg := range os.Args[0:] {
                num, err := strconv.Atoi(arg)
                if err != nil {
                        fmt.Printf("Error: %s is not a valid number\n", arg)
                        continue
                }
  		fmt.Println(num, rfactorial(num))

                wg.Add(1)
                go factorial(num, &wg, resultChan)
        }

        go func() {
                wg.Wait()
                close(resultChan)
        }()

        for result := range resultChan {
                fmt.Println(result)
        }
}

