package main

import (
	"fmt"

	"github.com/BarbalberoKills/gitFolderTest/userinput"
)

func swap[T userinput.AnyValue](a, b *T) {
	*a, *b = *b, *a
}

func main() {

	var prompt string

	prompt = "Give me your first value"
	first := userinput.GetUserValue[float64](&prompt)

	prompt = "Give me your second value"
	second := userinput.GetUserValue[float64](&prompt)

	fmt.Println("Before the swap - ", first, second)
	swap(&first, &second)
	fmt.Println("After the swap - ", first, second)
}
