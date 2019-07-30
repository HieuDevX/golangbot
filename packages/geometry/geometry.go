package main

import (
	"fmt"
	"log"
	_ "log"

	"github.com/hieudevx/golangbot-tms/packages/geometry/rectangle"
)

/*
 * 1. Package variables
 */

var rectLen, rectWidth float64 = 6, 7

/*
 * 2. init func to validate if the length and width less than 0
 */

func init() {
	println("main package intiialized")
	if rectLen < 0 {
		log.Fatal("length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}

func main() {
	// var rectLen, rectWidth float64 = 6, 7
	fmt.Println("Geometrical shape properties")

	fmt.Printf("Area of rectangle %.2f \n", rectangle.Area(rectLen, rectWidth))
	fmt.Printf("Diagonal of the rectangle %.2f \n", rectangle.Diagonal(rectLen, rectWidth))
}
