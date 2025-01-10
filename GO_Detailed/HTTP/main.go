package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.geeksforgeeks.org/"

func main() {
	fmt.Println("* Let's Learn About HTTP WEB Request *")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response of the type :%T \n", response)
	response.Body.Close() //Caller's Responsibility to close the connection
	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data inside the response is:\n", string(databytes))
}
