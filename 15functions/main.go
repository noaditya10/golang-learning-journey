package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(4, 5)
	fmt.Println("Result is : ", result)

	proRes, myMessage := proAdder(2, 5, 3, 6)
	fmt.Println("Result in proAdder : ", proRes)
	fmt.Println("Message ", myMessage)

}

//dont forget function signature : type of parameter and type of return value
//e.g.: int
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "hi from proAdder"
}

func greeterTwo() {
	fmt.Println("Another method")
}

func greeter() {
	fmt.Println("Hello from golang")
}
