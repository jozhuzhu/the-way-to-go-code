package _struct

import "fmt"

type Celsius float64

func (c *Celsius) String() string {
	return fmt.Sprintf("%2.5f℃", *c)
}

func (c *Celsius) GetInfo() string {
	return fmt.Sprintf("%2.5f℃", *c)
}

func PrintCelsius() {
	var c Celsius = 36.6

	fmt.Println(c.GetInfo())
	fmt.Println(c)
	fmt.Println(&c)
}
