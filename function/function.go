package main

import "fmt"

func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

func rectProps1(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func rectProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return
}

func main() {
	// price, no := 90, 6
	// totalPrice := calculateBill(price, no)
	// fmt.Printf("Total price is %v \n", totalPrice)

	// area, perimeter := rectProps(10.8, 5.6)
	// area, perimeter := rectProps1(10.8, 5.6)
	// fmt.Printf("Area is %f, Perimeter is %f \n", area, perimeter)

	area, _ := rectProps(10.8, 5.6)
	fmt.Printf("Area is %f \n", area)
}
