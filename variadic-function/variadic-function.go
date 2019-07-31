package main

import (
	"fmt"
)

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T \n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Printf("%v found at index %v in %v \n", num, i, nums)
			found = true
		}
	}
	if !found {
		fmt.Printf("%v not found in %v \n", num, nums)
	}
	// nums = append(nums, 13)
	// fmt.Printf("new nums is %v \n", nums)
}

func change(s ...string) {
	s[0] = "Go"
}

func main() {
	// find(89, 89, 90, 95)
	// find(45, 56, 57, 45, 90, 109)
	// find(78, 38, 56, 98)
	// find(87)

	// nums := []int{13, 3, 20, 6}
	// // find(13, nums) // error
	// find(13, nums...)

	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}

// func find(num int, nums []int) {
// 	fmt.Printf("type of nums is %T \n", nums)
// 	found := false
// 	for i, v := range nums {
// 		if v == num {
// 			fmt.Printf("%v found at index %v in %v \n", num, i, nums)
// 			found = true
// 		}
// 	}
// 	if !found {
// 		fmt.Printf("%v not found in %v \n", num, nums)
// 	}
// }

// func main() {
// 	find(89, []int{89, 90, 95})
// 	find(45, []int{56, 67, 45, 90, 109})
// 	find(78, []int{38, 56, 98})
// 	find(87, []int{})
// }
