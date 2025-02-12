package main

import (
	"fmt"
)

func calc(a, b *int, operation *string) (int, error) {
	if *operation == "sum" {
		return *a + *b, nil
	} else if *operation == "product" {
		return *a * *b, nil
	} else {
		return 0, fmt.Errorf("Unsupporte operation: %s", *operation)
	}
}

func main() {
	var a, b int = 5, 5
	var op string = "sum"
	res, err := calc(&a, &b, &op)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result", res)
	}
}
