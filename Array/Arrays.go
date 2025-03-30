package main

import "fmt"

func main() {
	//without Initialization decleration
	var list1 [5]int
	list1[0] = 3
	//Decleration with initialization
	list2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(list2)
	//Arrays in Golang

	arr := [3]string{"Go", "Java", "Python"}
	for i, value := range arr {
		fmt.Println("Index is ", i, "Values is ", value)
	}
	arr33 := [5]int{10, 20, 30, 40, 50}
	slice := arr33[1:4] // [20 30 40]

	fmt.Println("Array Length:", len(arr33), "Capacity:", cap(arr33)) // Output: Length: 5 Capacity: 5
	fmt.Println("Slice Length:", len(slice), "Capacity:", cap(slice)) // Output: Length: 3 Capacity: 4

}
