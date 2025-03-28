package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	//GetRequest()
	//GetRequestAnotherWay()
	//PostRequest()
	PostFormRequest()

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
	const myurl = "http://localhost:8000/get"       //endpoint
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

func PostRequest() {
	const myurl = "http://localhost:8000/post"

	//fake Json
	requestBody := strings.NewReader(`
	 {
	      "coursename" : "GO Language ",
		  "price" : 0,
		  "platform" : "Udemy"
	 }`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//fromData
	data := url.Values{}
	data.Add("firstname", "Mandar")
	data.Add("email", "m@gmail.com")
	data.Add("Password", "p@12345")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))

}
