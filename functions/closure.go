package functions

import "fmt"

func F() {
	fv := func() {
		fmt.Println("Hello, World!")
	}

	fmt.Printf("The type of fv is %T\n", fv)
}
