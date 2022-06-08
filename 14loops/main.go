package main

import "fmt"

func main() {
	fmt.Println("welcome to loops in golang")

	days := []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// looping through index i
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v \n", index, day)
	// }

	// MOST PRACTICE:
	// for _, day := range days {
	// 	fmt.Printf("value is %v \n", day)
	// }

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 5 {
			// break

			goto labelNovi
		}

		fmt.Println("Value is :", rougueValue)
		rougueValue++
	}

labelNovi:
	fmt.Println("jump into label")
}
