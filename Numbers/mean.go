package main

import (
	"fmt"
)

func main() {
	// var x int
	// var y int

	x := 10.0
	y := 15.0

	// var mean int
	mean := (x + y) / 2.0

	fmt.Printf("Number x=%v and it's type is %T\n", x, x)
	fmt.Printf("Number y=%v and it's type is %T\n", y, y)

	fmt.Printf("The mean of x and y is %v and it's type is %T\n", mean, mean)
}
