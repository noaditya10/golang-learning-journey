package main

import "fmt"

func main() {
	fmt.Println("Stucts in golang")
	//there is no inheritance in golang; no super or parent

	novi := User{"Novi", "novi@go.dev", true, 18}
	fmt.Println(novi)
	fmt.Printf("novi details are : %+v\n", novi)
	fmt.Printf("Name is %v and email is %v\n", novi.Name, novi.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
