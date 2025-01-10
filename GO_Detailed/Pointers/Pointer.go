package main

import "fmt"

func main() {
	fmt.Println("Welcome  to the pointer in golang")

	num := 10
	var ptr = &num

	fmt.Println("Address of pointer is: ", ptr)
	fmt.Println("Value of actual pointer is: ", num)

	*ptr = *ptr + 2
	fmt.Println(num)
}
