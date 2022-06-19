package main

import "fmt"

func main() {
	fmt.Println("Stucts in golang")
	//there is no inheritance in golang; no super or parent

	novi := User{"Novi", "novi@go.dev", true, 18}
	fmt.Println(novi)
	fmt.Printf("novi details are : %+v\n", novi)
	fmt.Printf("Name is %v and email is %v\n", novi.Name, novi.Email)

	novi.GetStatus()
	novi.NewMail()
	fmt.Println("Still, the original mail is : ", novi.Email)
	//thats why we use pointer
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	// oneAge int
	//oneAge is not exportable because low case letter first
}

func (u User) GetStatus() {
	//it become method because passing a struct in argument
	fmt.Println("Is user active : ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is : ", u.Email)
}
