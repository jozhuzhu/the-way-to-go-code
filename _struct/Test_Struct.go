package _struct

import (
	"fmt"
	structPack "the-way-to-go-code/_struct/struct_pack"
)

func main() {
	struct1 := new(structPack.ExpStruct)
	struct1.Mi1 = 10
	struct1.Mf1 = 16.

	fmt.Printf("Mi1 = %d\n", struct1.Mi1)
	fmt.Printf("Mf1 = %f\n", struct1.Mf1)
}
