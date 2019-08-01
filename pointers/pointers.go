package main

import (
	"fmt"
)

func change(val *int) {
	*val = 55
}

func modify(arr *[3]int) {
	(*arr)[0] = 90
}

func modifyBySlice(sls []int) {
	sls[0] = 90
}

func main() {
	// b := 255
	// var a = &b
	// fmt.Printf("Type of a is %T \n", a)
	// fmt.Printf("address of b is %v \n", a)

	// a := 25
	// var b *int
	// if b == nil {
	// 	fmt.Printf("b is %v \n", b)
	// 	b = &a
	// 	fmt.Printf("b after initialization is %v \n", b)
	// }

	// b := 255
	// a := &b
	// fmt.Printf("address of b is %v \n", a)
	// fmt.Printf("value of b is %v \n", *a)
	// *a++
	// fmt.Printf("new value of b is %v \n", b)

	// a := 58
	// fmt.Printf("value of a before function call is %v \n", a)
	// b := &a
	// change(b)
	// fmt.Printf("value of a after function call is %v \n", a)

	a := [3]int{89, 90, 91}
	// modify(&a)
	// fmt.Println(a)

	modifyBySlice(a[:])
	fmt.Println(a)
}
