package main

import (
	"fmt"
	"time"
)

type Memoized struct {
	f     func(int) int
	cache map[int]int
}

func main() {
	var mem *Memoized
	mem = memoize(func(n int) int {
		if n <= 1 {
			return n
		}
		return mem.call(n-1) + mem.call(n-2)
	})

	// fmt.Println(mem.call(10))
	result := timeit(func() int {
		return mem.call(100)
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

func memoize(f func(int) int) *Memoized {
	return &Memoized{f: f, cache: make(map[int]int)}
}

func (m *Memoized) call(x int) int {
	if v, ok := m.cache[x]; ok {
		return v
	}
	result := m.f(x)
	m.cache[x] = result
	return result
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	result := fibonacci(n-1) + fibonacci(n-2)
	return result
}
