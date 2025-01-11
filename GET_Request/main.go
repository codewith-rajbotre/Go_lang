package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("GET Request")
	//PerformGetRequest()
	//PostRequest()
	PostForm()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close() //To close the response is our Responsibility
	fmt.Println("Status Code:", response.StatusCode)

	var responseString strings.Builder //Contains Librery

	content, _ := ioutil.ReadAll(response.Body)

	byteCount, _ := responseString.Write(content)

	fmt.Println(byteCount)               //Only Prints ByteCount
	fmt.Println(responseString.String()) //Prints actual Data
	//fmt.Println(string(content))

}

func PostRequest() {
	const myurl = "http://localhost:8000/post"

	//Json Data

	requestBody := strings.NewReader(`
	{
	     "name":"Jara",
		 "Learning":"Golang",
		 "LearnigHours" : 2
	}
	`)

	response, err := http.Post(myurl, "application/json", requestBody) //(url,typeOfData,requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))

}

func PostForm() {
	const myurl = "http://localhost:8000/postform"
	data := url.Values{}
	data.Add("firstname", "Raj")
	data.Add("Company", "Kcloud")
	data.Add("email", "r@gmail.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
