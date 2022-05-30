package _struct

import "fmt"

type employee struct {
	salary float64
}

func (e *employee) giveRaise(pecent float64) float64 {
	return e.salary * (1 + pecent/100)
}

func PrintSalary() {
	e := employee{1000}
	fmt.Println("before rising, salary is ", e.salary)
	fmt.Println("after rising, salary is ", e.giveRaise(50))
}
