package main

import "fmt"

func main() {
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	var result [5]float64 = square(thing1)
	fmt.Printf("\nThe result is: %v", result)
}
func square(thing2 [5]float64) [5]float64 {
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}
