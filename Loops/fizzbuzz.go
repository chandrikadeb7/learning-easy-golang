package main

import (
	"fmt"
)

func main() {
  // numbers 1 to 20
	for i := 1; i <= 20; i++ {
    // if divisible by both 3 and 5
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizz buzz")
			continue
      // if divisible by 3
		} else if i%3 == 0 {
			fmt.Println("fizz")
			continue
      // if divisible by 5
		} else if i%5 == 0 {
			fmt.Println("buzz")
			continue
		}
    // else print all numbers
		fmt.Println(i)
	}
}
