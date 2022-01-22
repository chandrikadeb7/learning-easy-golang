package main

import (
	"fmt"
)

func main() {
	myString := `
	I love to have ice-cream
	I have a chic fashion sense!
	` // multiple lines string
	fmt.Println(myString)
	fmt.Println(len(myString))                                                 //builtin string length function
	fmt.Printf("First Character: %v\n type is %T\n", myString[0], myString[0]) //unint8 is byte

	fmt.Println(myString[5:12]) //slicing 5th to 11th char
	fmt.Println(myString[:12])  //slicing till 11th char
	fmt.Println(myString[12:])  //slicing from 12th char to last

	fmt.Println(myString + "JK...!!") //concatenate string

}
