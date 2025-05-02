package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Mai samay hoon")
	presentTime := time.Now()

	fmt.Println("Current Time: ", presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdTime := time.Date(2023, time.October, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Created Time: ", createdTime.Format("01-02-2006 15:04:05 Monday"))
}
