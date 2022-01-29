package main

import (
	"fmt"
)

//declare a struct for Employee - initial char uppercase so exported variable
type Employee struct {
	FirstName string
	LastName  string
	EmpID     int64
	Phone     int
	Address   string
}

func main() {

	//instantiate a struct
	e1 := Employee{"Chandrika", "Deb", 1779, 834011190, "Jamshedpur, India"}
	fmt.Println(e1)         //field names not printed
	fmt.Printf("%+v\n", e1) //field names printed along with values

	e2 := Employee{
		FirstName: "James",
		LastName:  "Bond",
		EmpID:     700,
		Phone:     911,
		Address:   "Pasadena, CA",
	}

	fmt.Printf("Address of James: %s\n", e2.Address) //dot notation for a particular field value

	e3 := Employee{} //go assigns 0 to int and blank to string types
	fmt.Printf("%+v\n", e3)

}
