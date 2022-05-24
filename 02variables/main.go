package main

import "fmt"

const LoginToken string = "abc123" //public because initial case

func main() {
	var username string = "Novi"

	fmt.Println(username)
	fmt.Printf("Variable is type : %T \n", username)

	//boolean
	var isLoggedIn bool = false

	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is type : %T \n", isLoggedIn)

	//small
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is type : %T \n", smallVal)

	//small
	var smallFloat float32 = 255.2342344
	fmt.Println(smallFloat)
	fmt.Printf("Variable is type : %T \n", smallFloat)
}
