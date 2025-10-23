package main

import (
	"errors"
	"fmt"
	"os"
)

var ErrTaskAlreadyExists = errors.New("word not allowed")

func checker(myString string) (string, error) {
	if myString != "hello" {
		return "", ErrTaskAlreadyExists
	}
	return myString, nil
}

func main() {
	stringa := "hell"
	checkedStringa, err := checker(stringa)
	fmt.Printf("%v\n", err)
	fmt.Printf("%v\n", ErrTaskAlreadyExists)
	if errors.Is(err, ErrTaskAlreadyExists) {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("String %v is correct\n", checkedStringa)
}
