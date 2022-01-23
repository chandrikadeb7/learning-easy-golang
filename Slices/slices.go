package main

import (
	"fmt"
)

func main() {

	fruits := []string{"apple", "orange", "banana", "peach", "pineapple"}

	for i, fruit := range fruits { //loop over fruits using for loop and index
		fmt.Printf("Fruit %d: %s\t\n", i, fruit)
	}

	fmt.Println(fruits[4])  //output pineapple
	fmt.Println(fruits[:3]) //output till index 2
	fmt.Println("Appending kiwi to fruits")
	fruits = append(fruits, "kiwi") //append kiwi fruit
	fmt.Println(fruits)
	fmt.Printf("Value of fruit is: %v and it's type is %T \n", fruits[5], fruits[5]) //output index 5 and it's type

}
