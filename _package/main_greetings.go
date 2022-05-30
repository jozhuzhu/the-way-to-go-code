package _package

import (
	"fmt"
	"the-way-to-go-code/_package/greetings"
)

func Greet() {
	if greetings.IsAM() {
		fmt.Println(greetings.GoodDay())
	} else if greetings.IsEvening() {
		fmt.Println(greetings.GoodNight())
	}
}
