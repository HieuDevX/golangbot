package main

import "fmt"

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) // copies neededCountries to conuntriesCpy
	return countriesCpy
}

func main() {
	// a := [5]int{76, 77, 78, 79, 80}
	// var b = a[1:4] // creates a slice from a[1] to a[3]
	// fmt.Println(b)

	// c := []int{6, 7, 8} // creates and array and returns a slice reference
	// fmt.Println(c)

	// darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	// dslice := darr[2:5]
	// fmt.Printf("array before %v \n", darr)
	// for i := range dslice {
	// 	dslice[i]++
	// }
	// fmt.Printf("array after %v \n", darr)

	// numa := [3]int{78, 79, 80}
	// nums1 := numa[:]
	// nums2 := numa[:]
	// fmt.Printf("array before change 1 %v \n", numa)
	// nums1[0] = 100
	// fmt.Printf("array after modification to slice nums1 %v \n", numa)
	// nums2[1] = 101
	// fmt.Printf("array after modification to slice nums2 %v \n", numa)

	// fruitArray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	// fruitSlice := fruitArray[1:3]
	// fmt.Printf("length of slice %d capcity %d \n", len(fruitSlice), cap(fruitSlice))
	// fruitSlice = fruitSlice[:cap(fruitSlice)] // re-slicing fruitslice till its capacity
	// fmt.Printf("After re-slicing length is %d and capacity is %d \n", len(fruitSlice), cap(fruitSlice))

	// i := make([]int, 5, 5)
	// fmt.Println(i)

	/*
	 * append
	 */

	// cars := []string{"Ferrari", "Honda", "Ford"}
	// fmt.Printf("cars: %v has old length %d and capacity %d \n", cars, len(cars), cap(cars))
	// cars = append(cars, "Toyota")
	// fmt.Printf("cars: %v has new length %d and capacity %d \n", cars, len(cars), cap(cars))

	// var names []string
	// if names == nil {
	// 	fmt.Println("slice is nil going to append")
	// 	names = append(names, "John", "Sebastian", "Vinay")
	// 	fmt.Printf("names contents: %v \n", names)
	// }

	// veggies := []string{"potatoes", "tomatoes", "brinjal"}
	// fruits := []string{"oranges", "apples"}
	// food := append(veggies, fruits...)
	// fmt.Printf("food: %v \n", food)

	// nos := []int{8, 7, 6}
	// fmt.Printf("slice before function call %v \n", nos)
	// subtactOne(nos)
	// fmt.Printf("slice after function call %v \n", nos)

	/*
	 * Multidimensional slices
	 */

	// pls := [][]string{
	// 	{"C", "C++"},
	// 	{"Javascript"},
	// 	{"Go", "Rust"},
	// }

	// for _, v1 := range pls {
	// 	for _, v2 := range v1 {
	// 		fmt.Printf("%s ", v2)
	// 	}
	// 	fmt.Printf("\n")
	// }

	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}
