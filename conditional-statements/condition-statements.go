package main

import (
	"fmt"
)

func main() {
	// num := 10
	// if num%2 == 0 {
	// 	fmt.Println("The number is even")
	// } else {
	// 	fmt.Println("The number is odd")
	// }

	// if num := 10; num%2 == 0 {
	// 	fmt.Println("The number is even")
	// } else {
	// 	fmt.Println("The number is odd")
	// }

	num := 99
	if num <= 50 {
		fmt.Println("number is less than or equal to 50")
	} else if num >= 51 && num <= 100 {
		fmt.Println("number is between 51 and 100")
	} else {
		fmt.Println("number is greater than 100")
	}
}
