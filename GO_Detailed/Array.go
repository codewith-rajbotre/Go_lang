package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang:")

	var FruitList [3]string
	FruitList[0] = "Apple"
	FruitList[1] = "Mango"
	FruitList[2] = "Banana"

	fmt.Println(FruitList)      //[Apple Mango Banana]
	fmt.Println(len(FruitList)) //3

	var Numbers = [10]int{122, 4, 56, 7, 3}
	// Numbers[0] = 23
	// Numbers[1] = 56
	// Numbers[2] = 15

	fmt.Println(Numbers)      //[122 4 56 7 3 0 0 0 0 0]
	fmt.Println(len(Numbers)) //10
}
