package main

import (
	"fmt"
	"os"
)

//centre point struct
type Point struct {
	X int
	Y int
}

//square struct
type Square struct {
	Centre Point
	Length int
}

//new square return function
func NewSquare(x int, y int, length int) (*Square, error) {

	//validation if length is negative
	if length <= 0 {
		return nil, fmt.Errorf("Length of a square must be positive")
	}

	//centre point new
	point := &Point{
		X: x,
		Y: y,
	}

	//square new
	square := &Square{
		Centre: *point,
		Length: length,
	}

	return square, nil

}

//centre point move function
func (s *Square) Move(dx int, dy int) {
	s.Centre.X += dx
	s.Centre.Y += dy
}

//square area calculate
func (s *Square) Area() int {
	return s.Length * s.Length
}

func main() {

	//initialise new square with centre point and side length
	s1, err := NewSquare(1, 2, 10)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Printf("Original Square: %+v\n", *s1)
	}

	//call move function
	s1.Move(2, 1)
	fmt.Printf("After moving Square: %+v\n", *s1)

	//call area function
	area := s1.Area()
	fmt.Printf("Area of Square: %+v\n", area)
}
