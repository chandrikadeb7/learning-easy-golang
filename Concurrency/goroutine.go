package main

import (
	"fmt"
	"net/http"
	"sync"
)

// get header content type of request
func contentType(url string) {
	resp, err := http.Get(url) //get the URL
	if err != nil {
		fmt.Printf("Error in URL %v\n", err) //error in get method
	}

	defer resp.Body.Close() //close the response body

	ctype := resp.Header.Get("Content-Type")             //get the content type
	fmt.Printf("URL: %s , Content-Type: %s", url, ctype) //return the content type

}

func main() {
	urls := []string{
		"https://golang.org",
		"https://google.com",
		"https://api.github.com",
	}

	var wg sync.WaitGroup //add a wait group to wait for goroutines to finish

	for _, url := range urls {
		wg.Add(1) //add another worker for every spin
		go func(url string) {
			contentType(url)
			wg.Done() //signal done
		}(url)
	}
	wg.Wait() //wait till all workers finish
}
