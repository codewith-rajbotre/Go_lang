package main

import "fmt"

func main() {
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Seeta": 10, "Geeta": 45}
	fmt.Println(myMap2["Seeta"]) //10
	var age, ok = myMap2["Geeta"]
	if ok {
		fmt.Printf("The age is %v", age) //The age is 45
	} else {
		fmt.Println("Invalid Name")
	}
	for name := range myMap2 {
		fmt.Printf("Name: %v\n", name) // Name: Seeta   Name: Geeta
	}
	var intArr [3]int32
	for name, age := range myMap2 {
		fmt.Printf("Name;%v,Age:%v \n", name, age)
	}

	for i, v := range intArr {
		fmt.Printf("Index: %v, value: %v \n", i, v)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
