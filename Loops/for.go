package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		if i > 8 {
			fmt.Printf("Breaking loop. Value of i is %v\n", i)
			break
		} else if i == 5 {
			fmt.Printf("Continuing loop. Value of i is %v\n", i)
			continue
		}
		fmt.Println(i)
	}
}
