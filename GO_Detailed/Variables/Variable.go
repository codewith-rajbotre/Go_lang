package main

import (
	"fmt"
	"unicode/utf8"
)

// For empty int value id 0
// For empty String value is empty string (" ")
// For empty bool value is false
func main() {
	var num int = 32767
	num = num + 1
	fmt.Println(num)

	var username string = "Rajshri"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T", username)

	var floatnum float64 = 35873209.9
	fmt.Println(floatnum)

	var floatnum32 float32 = 10.5
	var intnum32 int32 = 2
	var result float32 = floatnum32 + float32(intnum32)
	fmt.Println(result)

	var intnum1 int = 5
	var intnum2 int = 2
	fmt.Println(intnum1 / intnum2)
	fmt.Println(intnum1 % intnum2)

	var myString string = "Hello" + " " + "World"
	fmt.Println(myString)

	fmt.Println(utf8.RuneCountInString("Y"))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)

	var intNum3 rune
	fmt.Println(intNum3)

	// var myVar string = foo()
	// fmt.Println(myVar)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	const myConst string = "const value"
	fmt.Println(myConst)
}
