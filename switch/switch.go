package main

import (
	"fmt"
)

func number() int {
	num := 15 * 5
	return num
}

func main() {
	// finger := 8
	// switch finger {
	// case 1:
	// 	fmt.Println("Ngon cai")
	// case 2:
	// 	fmt.Println("Ngon tro")
	// case 3:
	// 	fmt.Println("Ngon giua")
	// case 4:
	// 	fmt.Println("Ngon ap ut")
	// case 5:
	// 	fmt.Println("Ngon ut")
	// default:
	// 	fmt.Println("So ngon tay khong chinh xac")
	// }

	// letter := "e"
	// switch letter {
	// case "a", "e", "i", "o", "u":
	// 	fmt.Println("Nguyen am")
	// default:
	// 	fmt.Println("Khong phai la mot nguyen am")
	// }

	// num := 75
	// switch {
	// case num >= 0 && num <= 50:
	// 	fmt.Println("num lon hon 0 va nho hon 50")
	// case num >= 51 && num <= 100:
	// 	fmt.Println("num lon hon 51 va nho hon 100")
	// case num >= 101:
	// 	fmt.Println("num lon hon 100")
	// }

	switch num := number(); {
	case num < 50:
		fmt.Printf("%d thi nho hon 50 \n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d thi nho hon 100 \n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d thi nho hon 200 \n", num)
	}
}
