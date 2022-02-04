package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://httpbin.org/get") //get the URL - handles https requests
	if err != nil {
		log.Fatalf("Error getting URL: %v", err) //error handling
	}

	defer resp.Body.Close() //close the response body

	io.Copy(os.Stdout, resp.Body) //get the response body to Stdout
}
