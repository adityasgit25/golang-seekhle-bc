package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv" // for converting string to other data types
	"strings" // powerful package for manipulating strings
)

func main() {
	fmt.Println("Welcome to my pizza shop!")
	fmt.Println("rate my pizza bc :).  ")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating my pizza: ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Adding 1 to your rating: ", numRating+1)
	}
}
