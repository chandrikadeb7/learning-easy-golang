package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "Error in URL", err
	}

	defer resp.Body.Close()

	ctype := resp.Header.Get("Content-Type")
	if ctype == " " {
		return "", fmt.Errorf("Cannot find Content-Type header in URL")
	}

	return ctype, nil

}

func main() {
	ctype, err := contentType("https://golang.org")

	if err != nil {
		fmt.Printf("Error %v\n", err)
	} else {
		fmt.Println(ctype)
	}
}
