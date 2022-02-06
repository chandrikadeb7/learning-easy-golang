package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//user struct to match Github json
type User struct {
	Name               string `json:"name"`
	PublicRepositories int    `json:"public_repos"`
}

func userInfo(login string) (*User, error) {
	//url for a particular user
	url := fmt.Sprintf("https://api.github.com/users/%s", login)
	//http get request
	resp, err := http.Get(url)
	//error handling
	if err != nil {
		return nil, err
	}

	//close response body
	defer resp.Body.Close()

	user := &User{}
	//decode the data
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(user); err != nil {
		return nil, err
	}

	return user, nil

}

func main() {
	//find data for given username
	user, err := userInfo("chandrikadeb7")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	//return the user data
	fmt.Printf("User: %+v\n", *user)
}
