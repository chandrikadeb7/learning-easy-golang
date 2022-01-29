package main

import "fmt"

// This a struct
type Mango struct {
	Name     string
	Quantity float64
	Price    float64
	DidBuy   bool
}

//function points to the mango struct
func (m *Mango) getValue(quantity float64, price float64, buy bool) float64 {
	value := 0.0
	//calculate the overall amount if mango is bought
	if buy {
		value = m.Quantity * m.Price
	}
	return value
}

func main() {

	//initialise a new mango
	john := Mango{
		Name:     "Alphonso",
		Quantity: 15.50,
		Price:    90.99,
		DidBuy:   true,
	}

	//call the function
	amount := john.getValue(john.Quantity, john.Price, john.DidBuy)
	//print the amount
	fmt.Printf("Amount spent by John: %.2f\n", amount)
}
