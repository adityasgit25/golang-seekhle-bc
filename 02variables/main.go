package main

import "fmt"

// go mod init variables

// when the variable's starting letter is capitalized, it becomes public
const LoginToken string = "asdasdasdasdasdasdasdasd" // public variable

func main() {
	var username string = "aditya"

	// see here as well Println is public function
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.325624234
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var number int
	fmt.Println(number)
	fmt.Printf("Variable is of type: %T \n", number)

	var rassi string
	fmt.Println(rassi)
	fmt.Printf("Variable is of type: %T \n", rassi)

	// implicit type
	var name = "aditya"
	fmt.Println(name)

	// no var style
	number2 := 24
	fmt.Println(number2)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
