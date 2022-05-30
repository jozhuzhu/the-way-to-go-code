package _struct

import "fmt"

const (
	MON = iota
	TUS
	WED
	THURS
	FRI
	SAT
	SUN
)

var days []string = []string{
	MON:   "Monday",
	TUS:   "Tuesday",
	WED:   "Wednesday",
	THURS: "Thursday",
	FRI:   "Friday",
	SAT:   "Saturday",
	SUN:   "Sunday",
}

type Day int

func (d Day) String() string {
	return days[d%7]
}

func WhichDay(d Day) {
	fmt.Println(d)
}
