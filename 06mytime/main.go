package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	// presentTime := time.Now()

	// fmt.Println(presentTime.Format("02-01-2006"))

	createdDate := time.Date(2020, time.August, 15, 23, 23, 0, 0, time.UTC)

	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}