package main

import (
	"fmt"
)

func cleanup(s string) {
	fmt.Printf("Cleaning up resource %s\n", s)
}

func worker() {
	defer cleanup("A")
	defer cleanup("B")

	fmt.Println("Worker Code")
}

func main() {
	worker()
}
