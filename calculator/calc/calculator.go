package calc

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

func Divide(a, b int) (int, error) {
	if b == 0 {
		return -1, fmt.Errorf("impossible to divide by 0")
	}
	return a / b, nil
}
