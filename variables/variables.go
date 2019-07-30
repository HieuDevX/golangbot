package main

import (
	"fmt"
	"math"
)

func main() {
	// var age = 30
	// age := 30
	// fmt.Println("my age is", age)
	// age = 23
	// fmt.Println("my age is", age)
	// age = 24
	// fmt.Println("my age is", age)

	// width, height := 100, 50
	// fmt.Println("width is", width, "height is", height)

	// var (
	// 	name   = "heaven"
	// 	age    = 29
	// 	height int
	// )
	// fmt.Println("My name is", name, ", age is", age, "and height is", height)
	// fmt.Printf("My name is %v, age is %v and height is %v", name, age, height)

	// name, age := "miracle", 23
	// fmt.Printf("my name is %v, age is %d", name, age)

	// a, b := 20, 30
	// fmt.Printf("a is %v, b is %v \n", a, b)
	// b, c := 40, 50
	// fmt.Printf("b is %v, c is %v \n", b, c)
	// b, c = 40, 50
	// fmt.Printf("change b is %v, c is %v \n", b, c)

	a, b := 145.8, 543.8
	c := math.Min(a, b)
	fmt.Printf("minimum value is %v \n", c)

	// age := 29
	// age = "miracle"
}
