package main

import "fmt"

// Describer interface
type Describer interface {
	Describe()
}

func main() {
	var d1 Describer
	// d1.Describe()	// error
	if d1 == nil {
		fmt.Printf("d1 is nil and has type %T value %v\n", d1, d1)
	}
}
