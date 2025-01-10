package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Hello lets learn something about file Handling ")
	content := "This is a contnet of file"
	file, err := os.Create("./myfile.txt")

	if err != nil {
		panic(err)
	}

	lenght, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("The length of file : ", lenght)
	defer file.Close()
	ReadFile("./myfile.txt")

}

func ReadFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("The data inside the file is : \n ", string(databyte))

}
