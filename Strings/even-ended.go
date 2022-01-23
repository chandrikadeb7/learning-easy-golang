package main

import (
	"fmt"
)

func main() {
	// Create result variable
	result := 0

	//loop through all 4-digit numbers
	for i := 1000; i <= 9999; i++ {
		for j := i; j <= 9999; j++ {
			//multiply the numbers
			n := i * j

			//convert number to string format
			s := fmt.Sprintf("%d", n)

			//check for even ended digits (same start and end digits)
			if s[0] == s[len(s)-1] {
				//increment the result count by 1
				result++
			}
		}
	}
	//output the result
	fmt.Printf("Even ended numbers on multiplying two 4-digit numbers: %v\n", result)
}
