package main

import "fmt"

func checkEvenOdd(numToCheck *int) {
	if *numToCheck%2 == 0 {
		fmt.Printf("The number %v is even\n", *numToCheck)
	} else {
		fmt.Printf("The number %v is odd\n", *numToCheck)
	}

}

func fibonacciSeriesGenerator(length *int) []int {
	series := make([]int, *length)
	for i := 0; i < *length; i++ {
		if i < 2 {
			series[i] = i
		} else {
			series[i] = series[i-1] + series[i-2]
		}
	}
	return series
}

func main() {

	var numToCheck int
	var fibonacciLength int

	for {
		fmt.Printf("Please give me a number:")
		_, err := fmt.Scan(&numToCheck)
		if err == nil {
			break
		}
		fmt.Printf("Invalid input. Please enter a number.\n")
		fmt.Scanln() // Clear inpit buffer to avoid infinte loop
	}
	checkEvenOdd(&numToCheck)

	for {
		fmt.Printf("How many element of the Fibonacci series would you like to see?\n")
		_, err := fmt.Scan(&fibonacciLength)
		if err == nil {
			break
		}
		fmt.Printf("Invalid input. Please enter a number.\n")
		fmt.Scanln()
	}
	s := fibonacciSeriesGenerator(&fibonacciLength)
	for i, v := range s {
		fmt.Printf("Fibonacci series: Term %v = %v\n", i+1, v)
	}
}
