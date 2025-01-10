package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader_name := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for the pizza:")

	// comma ok

	input_store_string_name, _ := reader_name.ReadString('\n')
	fmt.Println("Thanks for rating, ", input_store_string_name)
}
