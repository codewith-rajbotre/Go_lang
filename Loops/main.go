package main

import "fmt"

func main() {
	//Simple for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	//While like for loop
	num := 11
	for num > 0 {
		fmt.Println(num)
		num--
	}
	//Range for loop
	arr := []int{11, 3, 2, 55, 32}
	for _, element := range arr {
		fmt.Println(element)
	}

}
