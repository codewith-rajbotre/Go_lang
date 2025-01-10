package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Learning about URl : ")
	fmt.Println(myurl)
	//Parsing
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of qparams is : %T /n", qparams)
	fmt.Println(qparams["coursename"])

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
