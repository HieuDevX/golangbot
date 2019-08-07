package main

import (
	"fmt"
)

// Tester interface
type Tester interface {
	Test()
}

// MyFloat alias float64
type MyFloat float64

// Test method
func (m MyFloat) Test() {
	fmt.Println(m)
}

func describe(t Tester) {
	fmt.Printf("Interface type %T value %v\n", t, t)
}

func main() {
	var t Tester
	f := MyFloat(89.7)
	t = f
	describe(t)
	t.Test()
}
