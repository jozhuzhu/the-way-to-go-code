package _struct

import "fmt"

type Basics struct{}

func (Basics) Magic() {
	fmt.Println("base magic")
}

func (self Basics) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Basics
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

func DoMagic() {
	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()
}
