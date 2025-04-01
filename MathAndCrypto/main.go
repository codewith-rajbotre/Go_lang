package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	//"math/rand"
)

func main() {
	fmt.Println("Random Number")
	// By using Math Random package
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(6) + 1)

	//By using crypto random package
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}
