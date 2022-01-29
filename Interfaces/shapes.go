package main

import "fmt"

const PI float64 = 3.14

type Square struct {
	Side float64
}

type Circle struct {
	Radius float64
}

//area of square function
func (s *Square) Area() float64 {
	return s.Side * s.Side
}

//area of circle function
func (c *Circle) Area() float64 {
	return PI * c.Radius * c.Radius
}

//define interface having an abstract function and its return type
type Shape interface {
	Area() float64
}

//sum of areas function accepting a slice of areas
func sumAreas(shapes []Shape) float64 {

	total := 0.0
	for _, s := range shapes {
		total += s.Area()
	}

	return total
}

func main() {

	square := &Square{10}
	fmt.Printf("Area of square: %.2f\n", square.Area())

	circle := &Circle{5}
	fmt.Printf("Area of circle: %.2f\n", circle.Area())

	//creating a slice
	shapes := []Shape{square, circle}

	//calling sumAreas function
	area := sumAreas(shapes)

	fmt.Printf("Sum of areas: %v\n", area)

}
