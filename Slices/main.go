package main

import "fmt"

func main() {
	fmt.Println("Slices in Golang :")
	//Using Slice Literal
	number := []int{1, 34, 8, 56}
	fmt.Println(number)
	for i := 0; i < len(number); i++ {
		fmt.Println(number[i])
	}
	for index, value := range number {
		fmt.Println("Index :", index, "Value : ", value)
	}
	//Creating Slice from an Array

	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[0:3]
	fmt.Println(slice)
	slice = append(slice, 3, 4, 1)
	fmt.Println(slice)

	//Using make()

	slices := make([]int, 3, 5) //length 3 and capacity 5
	fmt.Println(slices)
	fmt.Println("Length :", len(slices), "Capacity : ", cap(slices))
	slices = append(slices, 3, 4)
	slices[0] = 90
	fmt.Println(slices)

}
