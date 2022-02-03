package main

import (
	"fmt"
	"net/http"
)

// get header content type of request
func contentType(url string, out chan string) {
	resp, err := http.Get(url) //get the URL
	if err != nil {
		out <- fmt.Sprintf("Error %s in URL %s\n", err, url)
		return //error in get method
	}

	defer resp.Body.Close() //close the response body

	ctype := resp.Header.Get("Content-Type")                     //get the content type
	out <- fmt.Sprintf("URL: %s , Content-Type: %s", url, ctype) //return the content type

}

func main() {

	urls := []string{
		"https://golang.org",
		"https://google.com",
		"https://api.github.com",
	}

	ch := make(chan string)
	for _, url := range urls {
		go contentType(url, ch)
	}

	for range urls {
		out := <-ch
		fmt.Println(out)
	}
}
