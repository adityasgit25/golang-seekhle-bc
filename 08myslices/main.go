// slices are arrays only but enhanced array, they are one of the most used things in golang, you won't be using arrays in golang much.

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in Go! 🥳")
	var fruitList = []string{"apple", "banana", "mango"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)
	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("The length of the fruitlist is: ", len(fruitList))
	fruitList = append(fruitList, "peach 🍑", "can I get a ho yeah")
	fmt.Println("Updated fruit list is: ", fruitList)

	fruitList = append(fruitList[1:4])
	fmt.Println("Updated fruit list is:", fruitList)

	highScores := make([]int, 4) // this syntax is also used to create slices
	highScores[0] = 234
	highScores[1] = 123
	highScores[2] = 456
	highScores[3] = 789

	// so here what it will do is reallocate the memory and add the new values to the slice
	highScores = append(highScores, 555, 666, 777)

	fmt.Println("High scores are: ", highScores)

	//these all methods are provided in slices not arrays
	sort.Ints(highScores)

	fmt.Println("Sorted high scores are: ", highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	// removing the value from slices using the index
	var courses = []string{"Math", "Science", "History", "English", "Geography"}
	fmt.Println("Courses list is: ", courses)
	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println("Updated courses list is: ", courses)

}
