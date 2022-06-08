package main

import "fmt"

func main() {
	fmt.Println("if else in golang")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}

	if num := 13; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is not less than 10")
	}

	//https://www.youtube.com/watch?v=Up4lTPhJBvs&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=19&ab_channel=HiteshChoudhary
}
