package main

import "fmt"

// This a struct
type Mango struct {
	Name     string
	Quantity float64
	Price    float64
	DidBuy   bool
}

func NewMango(name string, quantity float64, price float64, didBuy bool) (*Mango, error) {

	//validation conditions

	//name empty
	if name == "" {
		return nil, fmt.Errorf("Name cannot be empty!")
	}

	//quantity less than zero
	if quantity < 0.0 {
		return nil, fmt.Errorf("Quantity cannot be negative!")
	}

	//price less than zero
	if price < 0.0 {
		return nil, fmt.Errorf("Price cannot be negative!")
	}

	//new mango
	mango := &Mango{
		Name:     name,
		Quantity: quantity,
		Price:    price,
		DidBuy:   didBuy,
	}

	return mango, nil
}

func main() {

	m, err := NewMango("", 1.5, 99.99, true) //name empty error message

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", *m)
	}

}
