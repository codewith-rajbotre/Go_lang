package main

import "fmt"

func main() {
	greet()
	result := adder(3, 4)
	fmt.Println("Result of Addition is : ", result)
	fmt.Println("Substraction is : ", sub(112, 12))
}
func greet() {
	fmt.Println("Hello, welocme to golang world !!!")
}
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}
func sub(num1 int, num2 int) int {
	return num1 - num2
}
