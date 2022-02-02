package main

import (
	"fmt"
	"time"
)

func main() {
	num := make(chan int)

	// <-num
	// fmt.Println("We're here") //deadlock as nothing pushed in buffer

	go func() {
		num <- 99 //push in buffer
	}()

	val := <-num //receive from buffer
	fmt.Printf("We got %d\n", val)

	//push multiple values to buffer
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("Sending %d\n", i)
			num <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 3; i++ {
		val := <-num
		fmt.Printf("Received %d\n", val)
	}
}
