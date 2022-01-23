package main

import (
	"fmt"
	"strings"
)

func main() {
	//multiline string
	text := `
	This is a beautiful picture
	The picture is created with a creative mind
	Creative thinking often results in beautiful picture
	`
	//break string into words
	words := strings.Fields(text)

	//empty map
	count := map[string]int{}

	//lower case words and increment map
	for _, word := range words {
		count[strings.ToLower(word)]++
	}

	//print word frequency map
	fmt.Println(count)

}
