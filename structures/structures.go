package main

import (
	"fmt"
	_ "fmt"

	_ "github.com/hieudevx/golangbot/structures/computer"
)

// Employee struct example
type Employee struct {
	firstName, lastName string
	age, salary         int
}

// Person struct
type Person struct {
	string
	int
}

// Address struct
type Address struct {
	city, state string
}

// Person1 struct
type Person1 struct {
	name    string
	age     int
	address Address
}

// Person2 struct
type Person2 struct {
	name string
	age  int
	Address
}

// name struct
type name struct {
	firstName string
	lastName  string
}

// image struct
type image struct {
	data map[int]int
}

func main() {
	// creating structure using field names
	// emp1 := Employee{
	// 	firstName: "Sam",
	// 	age:       25,
	// 	salary:    500,
	// 	lastName:  "Anderson",
	// }

	// emp2 := Employee{"Thomas", "Paul", 29, 800}

	// fmt.Printf("Employee 1 %v \n", emp1)
	// fmt.Printf("Employee 2 %v \n", emp2)

	// var emp3 Employee // zero valued structure
	// fmt.Printf("Employee 3 %v \n", emp3)

	// emp4 := Employee{
	// 	firstName: "John",
	// 	lastName:  "Paul",
	// }
	// fmt.Printf("Employee 4 %v \n", emp4)

	// emp5 := Employee{"Sam", "Anderson", 55, 6000}
	// fmt.Printf("First Name: %v \n", emp5.firstName)
	// fmt.Printf("Last Name: %v \n", emp5.lastName)
	// fmt.Printf("Age: %v \n", emp5.age)
	// fmt.Printf("Salary: $%v \n", emp5.salary)

	// var emp6 Employee
	// emp6.firstName = "Jack"
	// emp6.lastName = "Adams"
	// fmt.Printf("Employee 6: %v \n", emp6)

	// emp7 := &Employee{"Sam", "Anderson", 55, 6000}
	// // fmt.Printf("First Name: %v \n", (*emp7).firstName)
	// // fmt.Printf("Age: %v \n", (*emp7).age)
	// fmt.Printf("First Name: %v \n", emp7.firstName)
	// fmt.Printf("Age: %v \n", emp7.age)

	// p := Person{"Naveen", 50}
	// fmt.Println(p)

	// var p1 Person
	// p1.string = "naveen"
	// p1.int = 50
	// fmt.Println(p1)

	// var p Person1
	// p.name = "Naveen"
	// p.age = 50
	// p.address = Address{
	// 	city:  "Chicago",
	// 	state: "Illinois",
	// }
	// fmt.Printf("Name: %v \n", p.name)
	// fmt.Printf("Age: %v \n", p.age)
	// fmt.Printf("City: %v \n", p.address.city)
	// fmt.Printf("State: %v \n", p.address.state)

	// var p Person2
	// p.name = "Naveen"
	// p.age = 50
	// p.Address = Address{
	// 	city:  "Chicago",
	// 	state: "Illinois",
	// }
	// fmt.Printf("Name: %v \n", p.name)
	// fmt.Printf("Age: %v \n", p.age)
	// fmt.Printf("City: %v \n", p.city)
	// fmt.Printf("State: %v \n", p.state)

	// var spec computer.Spec
	// spec.Maker = "apple"
	// spec.Price = 50000
	// // spec.model = "Mac Mini"	// error
	// fmt.Printf("Spec: %v \n", spec)

	name1 := name{"Steve", "Jobs"}
	name2 := name{"Steve", "Jobs"}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}

	name3 := name{firstName: "Steve", lastName: "Jobs"}
	name4 := name{}
	name4.firstName = "Steve"
	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}

	image1 := image{data: map[int]int{
		0: 155,
	}}

	fmt.Println(image1)

	image2 := image{data: map[int]int{
		0: 155,
	}}

	fmt.Println(image2)

	// if image1 == image2 { // error
	// 	fmt.Println("image1 and image2 are equal")
	// }
}
