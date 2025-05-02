package main

import "fmt"

func main() {
	// Pointers are used to store the memory address of a variable
	fmt.Println("Pointers in Go 🫦")
	// var ptr *int
	// fmt.Println("Pointer: ", ptr)

	// reference is &
	// "*" isse andar chale jaate hai.

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of the pointer: ", ptr)
	fmt.Println("Value of the pointer: ", *ptr)

	// operation will be performed on the actual value.
	*ptr = *ptr + 1
	fmt.Println("Value of the pointer: ", myNumber)
}
