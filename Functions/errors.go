package main

import (
	"fmt"
	"math"
)

func sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0.0, fmt.Errorf("Square root of negative number %.2f is not valid\n", a)
	}
	return math.Sqrt(a), nil
}

func main() {
	ans, err := sqrt(5.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Square root of 5.0 is %.2f\n", ans)
	}

	result, error := sqrt(-5.0)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Printf("Square root of -5.0 is %.2f\n", result)
	}

}
