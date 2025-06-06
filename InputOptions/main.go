package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Method : 1
	//Disadvantage :  reads input until space or newline

	// var name string
	// fmt.Println("Enter the string")
	// fmt.Scanln(&name)
	// fmt.Println("Name : ", name)

	//Method : 2

	Reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the String : ")
	input, _ := Reader.ReadString('\n')
	fmt.Println("String :", input)
}
