package main

import (
	"fmt"
	"net/http"
)

// get header content type of request
func contentType(url string) (string, error) {
	resp, err := http.Get(url) //get the URL
	if err != nil {
		return "Error in URL", err //error in get method
	}

	defer resp.Body.Close() //close the response body

	ctype := resp.Header.Get("Content-Type") //get the content type
	if ctype == " " {
		return "", fmt.Errorf("Cannot find Content-Type header in URL") //error in content type
	}

	return ctype, nil //return the content type

}

func main() {
	ctype, err := contentType("https://golang.org") //call the content type of URL

	if err != nil {
		fmt.Printf("Error %v\n", err) //error handling
	} else {
		fmt.Println(ctype) //content type output
	}
}
