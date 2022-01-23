package main

import (
	"fmt"
)

func add(a float64, b float64) float64 {
	return a + b
}

func divide(a int, b int) (int, int) {
	return a / b, a % b
}

func main() {
	sum := add(10, 15)
	fmt.Printf("The sum is %v\n", sum)

	ans, rem := divide(20, 3)
	fmt.Printf("The quotient is %v and the remainder is %v\n", ans, rem)

}
