package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

// use sync package inthat WaitGroup
// Three tasks we have to do 1.Add 2.Wait 3.Done

var waitGroup sync.WaitGroup //pointers

var mut sync.Mutex

func main() {

	websitelist := []string{
		"https://go.dev",
		"https://google.com",
		"https://github.com",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		waitGroup.Add(1)
	}
	// go greeter("Hello")
	// greeter("world")
	waitGroup.Wait()

	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(2 * time.Microsecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer waitGroup.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OPPs in Endpoint")
	} else {

		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d staus code for %s \n", res.StatusCode, endpoint)
	}
}
