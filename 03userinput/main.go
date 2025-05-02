package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"

	fmt.Println(welcome)

	// using : operator coz we don't know what will come as the input, (that's one of the usecase)
	reader := bufio.NewReader(os.Stdin) // Stdin is the standard input
	fmt.Println("Enter the rating for the pizza: ")

	// comma ok || err

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating ", input)
	fmt.Printf("The type of the input is %T\n", input)

}
