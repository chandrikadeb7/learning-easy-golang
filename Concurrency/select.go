package main

import (
	"fmt"
	"time"
)

func main() {
	num1, num2 := make(chan int), make(chan int) //make two channels

	go func() {
		num1 <- 99
	}()

	select {
	case val := <-num1:
		fmt.Printf("Got %d from num1\n", val)
	case val := <-num2:
		fmt.Printf("Got %d from num2\n", val)
	}

	out := make(chan float64)

	go func() {
		time.Sleep(100 * time.Millisecond)
		out <- 9.0
	}()

	select {
	case val := <-out:
		fmt.Printf("Got %.2f from out\n", val)
	case <-time.After(50 * time.Millisecond): //go routine still running in background
		fmt.Println("Timeout")
	}

}
