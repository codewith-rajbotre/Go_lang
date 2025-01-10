package main

import "fmt"

func main() {
	var intArr [3]int32
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	intArr2 := [3]int32{1, 2, 3}
	fmt.Println(intArr2)

	intArr3 := [...]int32{11, 22, 33}
	fmt.Println(intArr3)

}
