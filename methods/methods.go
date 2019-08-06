package main

import (
	"fmt"
)

// type Employee struct {
// 	name     string
// 	salary   int
// 	currency string
// }

/*
 * phương thức displaySalary có receiver là Employee
 */
// func (e Employee) displaySalary() {
// 	fmt.Printf("Salary of %s is %s%d \n", e.name, e.currency, e.salary)
// }

/*
 phương thức displaySalary() được chuyển thành hàm displaySalary với Employee là tham số truyền vào
*/
// func displaySalary(e Employee) {
// 	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
// }

// // Rectangle struct
// type Rectangle struct {
// 	length int
// 	width  int
// }

// // Circle struct
// type Circle struct {
// 	radius float64
// }

// // Area of Rectangle
// func (r Rectangle) Area() int {
// 	return r.length * r.width
// }

// // Area of Circle
// func (c Circle) Area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// // Employee struct
// type Employee struct {
// 	name string
// 	age  int
// }

/*
Phương thức với vật nhận là giá trị
*/
// func (e Employee) changeName(newName string) {
// 	e.name = newName
// }

// /*
// Phương thức với vật nhận là con trỏ
// */
// func (e *Employee) changeAge(newAge int) {
// 	e.age = newAge
// }

// type address struct {
// 	city  string
// 	state string
// }

// func (a address) fullAddress() {
// 	fmt.Printf("Full address: %s, %s", a.city, a.state)
// }

// type person struct {
// 	firstName string
// 	lastName  string
// 	address
// }

type rectangle struct {
	length int
	width  int
}

func area(r rectangle) {
	fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

func (r rectangle) area() {
	fmt.Printf("Area Method result: %d\n", (r.length * r.width))
}

func perimeter(r *rectangle) {
	fmt.Println("perimeter function output:", 2*(r.length+r.width))

}

func (r *rectangle) perimeter() {
	fmt.Println("perimeter method output:", 2*(r.length+r.width))
}

func main() {
	// emp1 := Employee{
	// 	name:     "Hieu Hoang",
	// 	salary:   5000,
	// 	currency: "$",
	// }

	// emp1.displaySalary() // gọi phương thức displaySalary() của kiểu Employ
	// displaySalary(emp1)

	// r := Rectangle{
	// 	length: 10,
	// 	width:  5,
	// }
	// fmt.Printf("Area of rectangle %d\n", r.Area())
	// c := Circle{
	// 	radius: 12,
	// }
	// fmt.Printf("Area of circle %f", c.Area())

	// e := Employee{
	// 	name: "Mark Andrew",
	// 	age:  50,
	// }
	// fmt.Printf("Employee name before change: %s", e.name)
	// e.changeName("Michael Andrew")
	// fmt.Printf("\nEmployee name after change: %s", e.name)

	// fmt.Printf("\n\nEmployee age before change: %d", e.age)
	// // (&e).changeAge(51)
	// e.changeAge(51) // cách viết tắt pointer, vẫn có kết quả như trên
	// fmt.Printf("\nEmployee age after change: %d", e.age)

	// p := person{
	// 	firstName: "Elon",
	// 	lastName:  "Musk",
	// 	address: address{
	// 		city:  "Los Angeles",
	// 		state: "California",
	// 	},
	// }

	// // explicit is: p.address.fullAddress()
	// p.fullAddress() //truy cập phương thức fullAddress của struct address

	r := rectangle{
		length: 10,
		width:  5,
	}
	area(r)
	r.area()

	p := &r
	/*
	 lỗi biên dịch, cannot use p (type *rectangle) as type rectangle
	 in argument to area
	*/
	// area(p)

	// p.area() //dùng con trỏ p gọi đến phương thức area() có vật nhận là giá trị

	perimeter(p)
	p.perimeter()

	/*
		cannot use r (type rectangle) as type *rectangle in argument to perimeter
	*/
	//perimeter(r)

	r.perimeter() //dùng giá trị r gọi đến phương thức perimeter() có vật nhận là con trỏ
}
