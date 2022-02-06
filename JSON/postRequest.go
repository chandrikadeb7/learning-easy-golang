package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type Job struct {
	User   string `json:"user"`
	Action string `json:"action"`
	Count  int    `json:"count"`
}

func main() {
	//post request data
	job := &Job{
		User:   "Chandrika",
		Action: "Codes",
		Count:  1,
	}

	//buffer - IO writer in-memory
	var buf bytes.Buffer

	//encode
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(job); err != nil {
		log.Fatalf("Error! Can't encode job: %v", err)
	}

	//post request
	resp, err := http.Post("https://httpbin.org/post", "application/json", &buf)
	if err != nil {
		log.Fatalf("Error! Can't post job: %v", err)
	}

	defer resp.Body.Close()

	//copy the response body to standard output
	io.Copy(os.Stdout, resp.Body)

}
