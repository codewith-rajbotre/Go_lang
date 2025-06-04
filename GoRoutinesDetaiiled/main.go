package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Goroutine started")
	time.Sleep(1 * time.Second)
	fmt.Println("Goroutine finished")
}

func main() {
	go sayHello()
	time.Sleep(2 * time.Second) // Let goroutine complete
	fmt.Println("Main function done")
}
