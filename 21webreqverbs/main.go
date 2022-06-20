package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("welcome")

	// PerformGetRequest()
	PerformPostJsonRequest()
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

func PerformPostJsonRequest() {
	const myurl = "http://localhost:7000/post"

	//fake json payload

	requestBody := strings.NewReader(`
		{
			"coursename" : "golang fundamentals",
			"price" : 0,
			"platform" : "camp404"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
