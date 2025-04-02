package main

import "fmt"

//1.Generic Function with type parameter T
func PrintValue[T any](value T) {
	fmt.Println(value)
}

//2. Generic struct with a type parameter T
type Box[T any] struct {
	Value T
}

func main() {
	//Type 1
	PrintValue(34)               // Works with int
	PrintValue("Hello Generics") // Works with string
	PrintValue(3.14)             // Works with float

	//Type 2
	intBox := Box[int]{Value: 100}
	stringBox := Box[string]{Value: "Golang"}
	floatBox := Box[float64]{Value: 45.34}

	fmt.Println(intBox.Value)    // 100
	fmt.Println(stringBox.Value) // Golang
	fmt.Println(floatBox.Value)  //45.34

}
