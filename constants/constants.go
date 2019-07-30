package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	// var a = math.Sqrt(4)
	// // const b = math.Sqrt(4)
	// const b = 2
	// fmt.Printf("a = %v, b = %v \n", a, b)

	// const hello = "Hello World"
	// fmt.Printf("type: %T, value: %v \n", hello, hello)

	// var hello1 = "Hello World"
	// fmt.Printf("type: %T, value: %v \n", hello1, hello1)

	// var defaultName = "Sam"
	// type myString string
	// var customName myString = "Sam"
	// customName = defaultName // error

	// const trueConst = true
	// type myBool bool
	// var defaultBool = trueConst
	// var customBool myBool = trueConst
	// defaultBool = customBool // error

	// const a = 5
	// var intVar = a
	// var int32Var int32 = a
	// var float64Var float64 = a
	// var complex64Var complex64 = a
	// fmt.Printf("intVar %T %v \nint32Var %T %v \nfloat64Var %T %v \ncomplex64Var %T %v \n", intVar, intVar, int32Var, int32Var, float64Var, float64Var, complex64Var, complex64Var)

	var a = 5.9 / 8
	fmt.Printf("a's type is %T, value is %v", a, a)
}
