package main

import (
	"fmt"
)

func main() {
	// var personSalary map[string]int
	// if personSalary == nil {
	// 	fmt.Println("map is nil. Going to make one")
	// 	personSalary = make(map[string]int)
	// }

	// personSalary := make(map[string]int)
	// personSalary["steve"] = 12000
	// personSalary["jamie"] = 15000
	// personSalary["mike"] = 9000
	// fmt.Printf("personSalary map contents: %v \n", personSalary)

	personSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary["mike"] = 9000
	// fmt.Printf("personSalary map contents: %v \n", personSalary)

	// personSalary["mike"] = 9000
	// employee := "jamie"
	// fmt.Printf("Salary of %v is %v \n", employee, personSalary[employee])
	// fmt.Printf("Salary of %v is %v \n", employee, personSalary["hieuht"])

	// newEmp := "hieuht"
	// value, ok := personSalary[newEmp]
	// if ok == true {
	// 	fmt.Printf("Salary of %v is %v \n", newEmp, value)
	// } else {
	// 	fmt.Printf("%v not found \n", newEmp)
	// }

	// fmt.Println("All items of a map")
	// for key, value := range personSalary {
	// 	fmt.Printf("personSalary[%s] = %d \n", key, value)
	// }

	// personSalary["nghiaht"] = 9500
	// fmt.Printf("map before deletion %v \n", personSalary)
	// delete(personSalary, "mike")
	// fmt.Printf("map after deletion %v \n", personSalary)
	// fmt.Printf("length is %v \n", len(personSalary))

	// newPersonSalary := personSalary
	// newPersonSalary["truongnd"] = 5000
	// fmt.Printf("Person salary changed %v \n", personSalary)

	map1 := map[string]int{
		"one": 1,
		"two": 2,
	}

	fmt.Printf("map1: %v \n", map1)

	map2 := map1

	fmt.Printf("map2: %v \n", map2)

	// if map1 == map2 {	// error

	// }
}
