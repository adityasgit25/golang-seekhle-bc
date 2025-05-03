// pkgm for package + func
package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps in Go! 🥳")
	myMap := make(map[string]int) // make keyword is also used to create maps
	myMap["JS"] = 10
	myMap["Python"] = 20
	myMap["Go"] = 30
	myMap["Java"] = 40
	fmt.Println("Map is: ", myMap)
	fmt.Println("JS score is: ", myMap["JS"])

	// loops in maps
	for key, value := range myMap {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
	fmt.Println("--------------------------------------")
	for key, _ := range myMap {
		fmt.Printf("For key %v, value is %v\n", key, myMap[key])
	}

}
