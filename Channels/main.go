package main

import (
	"fmt"
	"sync"
)

func main() {
	BufferdChannels()
	//fmt.Println("Channels in Golang")

	myCh := make(chan int, 1) //second parameter is buffer -Number of channels
	wg := &sync.WaitGroup{}

	// myCh <- 5
	// fmt.Println(<-myCh)
	wg.Add(2)
	//Recieve only
	go func(ch <-chan int, wt *sync.WaitGroup) {
		//close(myCh)
		val, isChanelOpen := <-myCh
		fmt.Println(isChanelOpen)
		fmt.Println(val)

		//fmt.Println(<-myCh)
		//fmt.Println(<-myCh)
		wg.Done()

	}(myCh, wg)
	//Send only
	go func(ch chan<- int, wt *sync.WaitGroup) {

		myCh <- 5
		// myCh <- 6
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
