package main

import "fmt"

// Describer interface
type Describer interface {
	Describe()
}

// Person struct
type Person struct {
	name string
	age  int
}

// Describe method: Person implements Describe
func (p Person) Describe() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}

func findType(i interface{}) {
	switch v := i.(type) {
	case Describer:
		v.Describe()
	default:
		fmt.Printf("unknown type\n")
	}
}

func main() {
	findType("Naveen")
	p := Person{
		name: "Naveen R",
		age:  25,
	}
	findType(p)
}
