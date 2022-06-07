package person

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName, lastName string
}

type Persons []Person

func (p Persons) Len() int {
	return len(p)
}

func (p Persons) Less(i, j int) bool {
	return strings.ToUpper(p[i].firstName) < strings.ToUpper(p[j].firstName)
}

func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Persons) String() string {
	output := "["
	for i := 0; i < p.Len(); i++ {
		output = fmt.Sprintf("%s\nPerson:{firstname: %s, lastname: %s},", output, p[i].firstName, p[i].lastName)
	}
	output = fmt.Sprintf("%s%s", output, "\n]")
	return output
}
