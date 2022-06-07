package _interface

import (
	"fmt"
	float642 "the-way-to-go-code/_interface/float64"
)

func Floats() {
	var f float642.Float64Array
	f = float642.NewFloat64Array()

	if !float642.IsSorted(f) {
		panic("failed")
	}
	fmt.Println(f)

	f = float642.Fill()
	float642.Sort(f)

	fmt.Println(f)
}
