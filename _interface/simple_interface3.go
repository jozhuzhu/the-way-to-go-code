package _interface

import "fmt"

func gI(any interface{}) {
	switch any.(type) {
	case *Simple:
		fmt.Println("This is Simple instance")
	case *RSimple:
		fmt.Println("This is RSimple instance")
	}
}
