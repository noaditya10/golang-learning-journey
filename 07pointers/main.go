package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// var ptr *int

	// fmt.Println("Value of pinter is ", ptr)
	//result is <nil>

	myNumber := 23
	//auto init type int

	var ptr = &myNumber

	fmt.Println("Value of actual pointer is", ptr)
	fmt.Println("Value of actual pointer is", *ptr)

	//*ptr get the value
	*ptr = *ptr * 2
	fmt.Println("New value is : ", myNumber)

}
