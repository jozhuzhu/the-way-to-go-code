package _interface

import (
	"fmt"
	"reflect"
	"testing"
)

type NotknownType struct {
	s1, s2, s3 string
}

func (n NotknownType) String() string {
	return n.s1 + " - " + n.s2 + " - " + n.s3
}

func (n NotknownType) Print(a int, b float64, c string) string {
	return fmt.Sprintf("%d-%f-%s", a, b, c)
}

// variable to investigate:
var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

func TestReflect(t *testing.T) {
	value := reflect.ValueOf(secret) // <main.NotknownType Value>
	typ := reflect.TypeOf(secret)    // main.NotknownType
	// alternative:
	// typ := value.Type()  // main.NotknownType
	fmt.Println(typ)
	knd := value.Kind() // struct
	fmt.Println(knd)

	// iterate through the fields of the struct:
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
		// error: panic: reflect.Value.SetString using value obtained using unexported field
		// value.Field(i).SetString("C#")
	}

	// call the first method, which is String():
	iv := reflect.ValueOf(1)
	fv := reflect.ValueOf(1.0)
	fmt.Println(fv.Type())
	sv := reflect.ValueOf("sdfkaf")
	inputs := []reflect.Value{iv, fv, sv}
	results := value.Method(0).Call(inputs)
	fmt.Println(results) // [Ada - Go - Oberon]
}
