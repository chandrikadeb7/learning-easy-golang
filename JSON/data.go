package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

//example data
var data = `
{
	"user" : "Chandrika Deb",
	"type" : "deposit",
	"amount" : 100000.50
}
	`

//request data mapped to Json format
type Request struct {
	Login  string  `json:"user"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

func main() {
	//Simulate a file/socket
	rdr := bytes.NewBufferString(data)

	//Decode request
	dec := json.NewDecoder(rdr)

	req := &Request{}

	//error handling
	if err := dec.Decode(req); err != nil {
		log.Fatalf("Cannot decode request: %+v\n", err)
	}

	fmt.Printf("Got %+v\n", *req)

	//create response
	prevBalance := 450000.30
	resp := map[string]interface{}{
		"ok":           true,
		"totalBalance": prevBalance + req.Amount,
	}

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(resp); err != nil {
		log.Fatalf("Failed to encode request: %+v\n", err)
	}

}
