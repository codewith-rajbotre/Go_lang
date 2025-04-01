package main

import "fmt"

func main() {
	greet()
	result := adder(3, 4)
	fmt.Println("Result of Addition is : ", result)
}
func greet() {
	fmt.Println("Hello, welocme to golang world !!!")
}
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}
