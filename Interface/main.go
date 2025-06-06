package main

import (
	"fmt"
	"strings"
)

// Def
type Examples interface {
	ToUpper() string
}

//Implement

type SchoolName struct {
	names string
}

func (s SchoolName) ToUpper() string {
	return strings.ToUpper(s.names)
}

func main() {
	var m Examples = SchoolName{"hello world"}
	fmt.Println("Upper:", m.ToUpper())
}

// // Defining an interface

// type Shape interface {
// 	Area() float64
// }

// // Implementing the interface in a struct
// type Rectangle struct {
// 	Length  float64
// 	Breadth float64
// }

// // Implementing the Area method for Rectangle
// func (r Rectangle) Area() float64 {
// 	return r.Length * r.Breadth
// }

// //
// func calculateArea(s Shape) float64 {
// 	return s.Area()
// }

// // Empty Interface
// func describeValue(t interface{}) {
// 	fmt.Printf("Type : %T , Value : %v \n", t, t)
// }

// func main() {
// 	// Empty Interface
// 	mysteryBox := interface{}("Hello")
// 	describeValue(mysteryBox)

// 	rect := Rectangle{Length: 45, Breadth: 2}
// 	//Another Way
// 	// var sa Shape
// 	// sa = rect
// 	// fmt.Println("Rectangle Area:", sa.Area())

// 	fmt.Println("Rectangle Area : ", calculateArea(rect))

// }
