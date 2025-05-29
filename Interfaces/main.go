package main

import "fmt"

// Define an interface
type Speaker interface {
	Speak() string
}

// Define a Person struct
type Person struct {
	Name string
}

// Define a Dog struct
type Dog struct {
	Breed string
}

// Implement Speak() for Person
func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}

// Implement Speak() for Dog
func (d Dog) Speak() string {
	return "Woof! I'm a " + d.Breed
}

// Main function
func main() {
	// Create instances
	p := Person{Name: "Alice"}
	d := Dog{Breed: "Labrador"}

	// Declare a Speaker interface
	var s Speaker

	// Assign and call Speak()
	s = p
	fmt.Println(s.Speak())

	s = d
	fmt.Println(s.Speak())

	// Alternatively, use a slice of Speaker
	speakers := []Speaker{p, d}
	fmt.Println("\nAll speakers say:")
	for _, speaker := range speakers {
		fmt.Println(speaker.Speak())
	}
}
