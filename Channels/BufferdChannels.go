package main

import "fmt"

func BufferdChannels() {
	ch := make(chan int, 3) // capacity 3

	ch <- 10
	ch <- 20
	ch <- 30

	fmt.Println(<-ch) // 10
	fmt.Println(<-ch) // 20
	fmt.Println(<-ch) // 30
}
