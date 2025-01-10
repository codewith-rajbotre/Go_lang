package main

import "fmt"

func main() {

	//sliceCopy Function
	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)     // [1 2 4]
	fmt.Println(sliceCopy) //[1 2 4]

	//Pointer
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("The value p points to is: %v", *p)
	fmt.Printf("\nThe value if i is: %v", i)
	p = &i
	*p = 1
	fmt.Printf("\nThe value p points to is: %v", *p)
	fmt.Printf("\nThe value if i is: %v", i)
}
