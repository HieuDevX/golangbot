package main

import "fmt"

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Printf("inside function %v \n", num)
}

func printArray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func main() {
	// var a [3]int
	// fmt.Println(a)
	// a[0] = 12
	// a[1] = 78
	// a[2] = 50
	// fmt.Println(a)

	// a := [3]int{12, 78, 50}
	// fmt.Println(a)

	// a := [3]int{12}
	// fmt.Println(a)

	// a := [...]int{12, 78, 50}
	// fmt.Println(a)

	// a := [3]int{5, 78, 8}
	// var b [5]int
	// b = a // error

	// a := [...]string{"USA", "China", "India", "Germany", "France"}
	// b := a
	// b[0] = "Singapore"
	// fmt.Printf("a is %v \n", a)
	// fmt.Printf("b is %v \n", b)

	// num := [...]int{5, 6, 7, 8, 9}
	// fmt.Printf("before passing to function %v \n", num)
	// changeLocal(num)
	// fmt.Printf("after passing to function %v \n", num)

	// a := [...]float64{67.7, 89.8, 21, 78}
	// fmt.Printf("length of a is %v \n", len(a))
	// for i := 0; i < len(a); i++ {
	// 	fmt.Printf("%d element of a is %.2f \n", i, a[i])
	// }

	// a := [...]float64{67.7, 89.8, 21, 78}
	// sum := float64(0)
	// for i, v := range a { // range return both the index and value
	// 	fmt.Printf("%d element of a is %0.2f \n", i, v)
	// 	sum += v
	// }
	// for _, v := range a { // range return both the index and value
	// 	fmt.Printf("element of a is %0.2f \n", v)
	// 	sum += v
	// }
	// for i := range a { // range return both the index and value
	// 	fmt.Printf("%d element of a \n", i)
	// 	sum += float64(i)
	// }
	// fmt.Printf("Sum of all index element of a is %v \n", sum)

	/*
	 * Multidimensional arrays
	 */

	a := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, // this comma is necessary. The compiler will complain if you omit this comma
	}
	printArray(a)

	var b [3][2]string
	b[0][0] = "apple"
	b[0][1] = "samsum"
	b[1][0] = "microsoft"
	b[1][1] = "google"
	b[2][0] = "AT&T"
	b[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printArray(b)
}
