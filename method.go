package main

import (
	"fmt"
	"math"
)

/* define a cirlce */
type Circle struct {
	radius float64
}

/* define a method for circle */
func (circle Circle) area() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}

func main() {
	circle := Circle{radius:5}
	fmt.Printf("Circle area: %f\n", circle.area())
}
