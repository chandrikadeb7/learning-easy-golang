package main

import (
	"fmt"
)

func main() {

	v := "E"

	switch v {
	case "A", "a", "E", "e", "I", "i", "O", "o", "U", "u":
		fmt.Printf("You entered a vowel letter %v \n", v)
	default:
		fmt.Println("You entered a consonant letter!")
	}

}
