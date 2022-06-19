package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Golang")

	content := "this needs to go in a file"

	file, err := os.Create("./file.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNillErr(err)

	length, err := io.WriteString(file, content)

	checkNillErr(err)

	fmt.Println("length is : ", length)
	defer file.Close()

	readFile("./file.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	checkNillErr(err)
	fmt.Println("Text data inside the file is \n", string(databyte))
}

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}
