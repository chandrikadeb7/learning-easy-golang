package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your age: ")
	age, _ := reader.ReadString('\n')
	fmt.Printf("Age: %v \n", age)
	myAge, err := strconv.ParseInt(strings.TrimSpace(age), 10, 64)

	if err != nil {
		fmt.Println(err)
	}

	if myAge >= 18 {
		fmt.Println("Eligible to vote!")
		return
	} else if myAge > 0 && myAge < 18 {
		fmt.Println("Not eligible to vote!")
		return
	}

	fmt.Println("Please enter a valid age!")

}
