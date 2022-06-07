package person

import (
	"fmt"
	float642 "the-way-to-go-code/_interface/float64"
)

func SortPersons() {
	p1 := Person{
		"Yao",
		"xinlianag",
	}

	p2 := Person{
		"wang",
		"zhiyuan",
	}

	p3 := Person{
		"xia",
		"longfei",
	}

	persons := Persons{p3, p2, p1}
	fmt.Println(persons)

	float642.Sort(persons)
	fmt.Println(persons)
}
