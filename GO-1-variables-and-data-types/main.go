package main

import "fmt"

func main() {
	var number int    // declaration
	var fNumber = 0.6 // declaration and initialization
	str := "string"   // short declaration and initialization
	number = 5        // initialization
	boolVar := true

	fmt.Printf("number variable is a type %T, and it's value is: %v\n", number, number)
	fmt.Printf("fnumber variable is a type %T, and it's value is: %v\n", fNumber, fNumber)
	fmt.Printf("str variable is a type %T, and it's value is: %v\n", str, str)
	fmt.Printf("boolvar variable is a type %T, and it's value is: %v\n", boolVar, boolVar)

	var age int
	var height float64
	fmt.Printf("Please tell me your age\n")
	fmt.Scan(age)

	fmt.Printf("Please tell me your heigth\n")
	fmt.Scan(height)
	fmt.Printf("Your age is a type %T, and it's value is: %v\n", age, age)
	fmt.Printf("Your height is a type %T, and it's value is: %v\n", height, height)

}
