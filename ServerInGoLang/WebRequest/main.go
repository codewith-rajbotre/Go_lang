package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	//fmt.Println("Hello This file is regarding to GET request in GoLang")
	//GetRequest()
	GetRequestAnotherWay()

}

func GetRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status Code", response.StatusCode)

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))

}

func GetRequestAnotherWay() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Status Code", response.StatusCode)

	var responseString strings.Builder

	content, _ := ioutil.ReadAll(response.Body)

	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is", byteCount)

	fmt.Println(responseString.String())

	defer response.Body.Close()

}
