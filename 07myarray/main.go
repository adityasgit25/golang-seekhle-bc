package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in Go! 🥳")
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Banana"

	fmt.Println("Fruit List: ", fruitList)
	fmt.Println("Length of the array: ", len(fruitList))

	var vegList = [4]string{"Potato", "Tomato", "Onion"}
	fmt.Println("Veg List: ", vegList)
	fmt.Println("Length of the array: ", len(vegList))

}
