package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("welcome")

	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:7000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)

	fmt.Println("Content length is : ", response.ContentLength)

	//to read

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount : ", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(string(content))

	//in some cases we use strings.Builder to get response in origin form and can convert in desired form

}
