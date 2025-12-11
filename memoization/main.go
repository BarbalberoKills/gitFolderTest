package main

import (
	"fmt"
	"time"
)

var cache = make(map[int]int)

func main() {
	result := timeit(func() int {
		return fibonacci(1)
	})
	fmt.Println(result)
}

func timeit(f func() int) int {
	start := time.Now()
	result := f()
	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %v\n", elapsed)
	return result
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	if v, ok := cache[n]; ok {
		return v
	}
	result := fibonacci(n-1) + fibonacci(n-2)
	cache[n] = result
	return result
}
