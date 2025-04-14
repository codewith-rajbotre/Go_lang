package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func sub(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, number := range nums {
		total -= number
	}
	fmt.Println(total)
}

func main() {

	sum(1, 2)
	sum(1, 2, 3)
	sub(7, 5, 3, 2)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
