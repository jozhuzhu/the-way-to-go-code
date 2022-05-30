package _package

import (
	"fmt"
	"unsafe"
)

func CheckInt() {
	k := 1.0
	size := unsafe.Sizeof(k)

	fmt.Println(size)
}
