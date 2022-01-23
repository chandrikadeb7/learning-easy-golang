package main

import (
	"fmt"
)

func main() {
	// Create a map
	cities := map[string]string{
		"IXR": "Ranchi",
		"CCU": "Kolkata",
	}

	//Print the cities map
	fmt.Println(cities)

	//Add Delhi to the map
	cities["DEL"] = "Delhi"
	fmt.Println(cities)

	//Return the value of the city from key
	fmt.Println(cities["IXR"])

	//Loop through the cities map
	for key, city := range cities {
		fmt.Printf("City %s has code %s\n", city, key)
	}

	//Delete Ranchi from the map
	delete(cities, "IXR")
	fmt.Println(cities)

}
