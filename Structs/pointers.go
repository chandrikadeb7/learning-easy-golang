package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
	Z int
}

func (p *Point) Move(dx int, dy int, dz int) {
	p.X += dx
	p.Y += dy
	p.Z += dz
}

func main() {
	p := &Point{1, 2, 3}
	p.Move(3, 2, 1)

	fmt.Printf("After Moving: %+v\n", p)
}
