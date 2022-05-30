package _struct

import "fmt"

type Anonymous struct {
	data float32
	int
	string
}

func PrintAnonymous() {
	anony := Anonymous{1.0, 2, "kidding me !"}

	fmt.Println(anony)
}
