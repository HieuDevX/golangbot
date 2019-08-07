package main

import (
	"fmt"
)

// SalaryCalculator interface
type SalaryCalculator interface {
	DisplaySalary()
}

// LeaveCalculator interface
type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

// Employee struct
type Employee struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

// DisplaySalary method: Employee implements SalaryCalculator
func (e Employee) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

// CalculateLeavesLeft method: Employee implements LeaveCalculator
func (e Employee) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func main() {
	e := Employee{
		firstName:   "Naveen",
		lastName:    "Ramanathan",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	var s SalaryCalculator = e
	s.DisplaySalary()
	var l LeaveCalculator = e
	fmt.Println("\nLeaves left =", l.CalculateLeavesLeft())
}
