package main

import "fmt"

// struct
type Point struct {
    x, y int
}


// app entry point
func main() {

	// struct
	p := Point{1, 2}
	fmt.Println(p)

	// string
	fmt.Println("hello...")
	fmt.Println("hello", "world")

	// array
	var aryOne [10]int // declare
	fmt.Println("emp: ", aryOne)
	aryOne[8] = 10
	fmt.Println("set:", aryOne)
	fmt.Println("get:", aryOne[8])

	aryTwo := [6]int{1,2,3,4,5,6} // init
	fmt.Println("dcl: ", aryTwo)

	var twoD [2][3]int
	for ith:= 0; ith < 2; ith++ {
		for jth := 0; jth < 3; jth++ {
			twoD[ith][jth] = ith + jth
		}
	}
	fmt.Println("2d: ", twoD)
}