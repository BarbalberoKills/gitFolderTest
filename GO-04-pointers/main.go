package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {

	var first, second = 5, 299

	fmt.Println(first, second)
	swap(&first, &second)
	fmt.Println(first, second)
}
