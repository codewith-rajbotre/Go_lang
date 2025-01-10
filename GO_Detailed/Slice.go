package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slice in golang")

	var Names = []string{"Ram", "Shyam"}
	fmt.Printf("Typemof array is:%T\n ", Names)

	Names = append(Names, "Seeta", "Geeta")
	fmt.Println(Names)

	Names = append(Names[1:3])
	fmt.Println(Names)

	highscore := make([]int, 4)
	highscore[0] = 234
	highscore[1] = 354
	highscore[2] = 674
	highscore[3] = 114
	//highscore[4] = 456

	highscore = append(highscore, 222, 333)
	fmt.Println(highscore)

	sort.Ints(highscore)
	fmt.Println(highscore)
	fmt.Println(sort.IntsAreSorted(highscore))

}
