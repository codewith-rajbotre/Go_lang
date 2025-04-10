package main

import "fmt"

func main() {
	var intSlice []int32 = []int32{4, 5, 6}
	var stringSlice []string = []string{"Mandar", "Manoj", "Gopal"}
	fmt.Println(stringSlice)
	fmt.Println(intSlice) //[4,5,6]
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice) //[4,5,6,7]

	var iniSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, iniSlice2...)
	fmt.Println(intSlice) //[4,5,6,7,8,9]
}
