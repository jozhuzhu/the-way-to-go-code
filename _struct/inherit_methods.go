package _struct

import "fmt"

type Base struct {
	id string
}

func (b *Base) Id() string {
	return b.id
}

func (b *Base) SetId(id string) {
	b.id = id
}

type Person struct {
	*Base
	FirstName, LastName string
}

type Employee struct {
	*Person
	salary float64
}

func EnrollEmployee() {
	b := new(Base)
	b.SetId("emp01")
	e := &Employee{
		&Person{
			b,
			"Joseph",
			"Jiang",
		},
		1000.00,
	}

	fmt.Println(e.Id())
}
