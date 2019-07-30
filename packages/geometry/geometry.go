package main

import (
	"fmt"

	"github.com/hieudevx/golangbot-tms/packages/geometry/rectangle"
)

func main() {
	var rectLen, rectWidth float64 = 6, 7
	fmt.Println("Geometrical shape properties")

	fmt.Printf("Area of rectangle %.2f \n", rectangle.Area(rectLen, rectWidth))
	fmt.Printf("Diagonal of the rectangle %.2f \n", rectangle.Diagonal(rectLen, rectWidth))
}
