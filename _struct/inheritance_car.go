package _struct

import "fmt"

type Engine interface {
	Start()
	Stop()
}

type miniEngine struct{}

func (m *miniEngine) Start() {
	fmt.Println("miniEngine has started")
}

func (m *miniEngine) Stop() {
	fmt.Println("miniEngine has stopped")
}

type Car struct {
	wheelCount int
	Engine
}

func (c *Car) numbersOfWheels() int {
	return c.wheelCount
}

type Mercedes struct {
	*Car
}

func (m *Mercedes) sayHiToMerkel() {
	fmt.Println("Hi, Merkel")
}

func GoForFun() {
	car := &Car{
		4,
		&miniEngine{},
	}

	fmt.Printf("My Car has %d wheels\n", car.numbersOfWheels())
	car.Start()
	car.Stop()

	mercedes := &Mercedes{
		car,
	}

	mercedes.sayHiToMerkel()
	mercedes.Start()
	mercedes.Stop()
}
