package main

import (
	"fmt"
)

// Ordered constraint for types that support comparison
type Ordered interface {
	~int | ~float64 | ~string
}

// Generic function to find the maximum value in a slice
func Max[T Ordered](items []T) T {
	if len(items) == 0 {
		panic("empty slice")
	}
	max := items[0]
	for _, item := range items {
		if item > max {
			max = item
		}
	}
	return max
}

func main() {
	ints := []int{1, 3, 2, 10, 5}
	floats := []float64{1.1, 3.5, 2.2, 0.9}
	strs := []string{"apple", "banana", "cherry"}

	fmt.Println("Max int:", Max(ints))
	fmt.Println("Max float:", Max(floats))
	fmt.Println("Max string:", Max(strs))
}
