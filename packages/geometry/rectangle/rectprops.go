package rectangle

import (
	"fmt"
	"math"
)

/*
 * init func
 */
func init() {
	fmt.Println("rectangle package initialized")
}

// Area func
func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

// Diagonal func
func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}
