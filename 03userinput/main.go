package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "WElcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("ENter the rating for our Pizza :")

	//comma ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)

	// https://www.youtube.com/watch?v=zYIZtbyUIDY&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=7&ab_channel=HiteshChoudhary
}
