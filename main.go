package main

import (
	"the-way-to-go-code/_struct"
)

func main() {
	for i := 0; i < 10; i++ {
		_struct.WhichDay(_struct.Day(i))
	}
}
