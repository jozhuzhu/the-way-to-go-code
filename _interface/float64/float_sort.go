package float64

import (
	"fmt"
	"math/rand"
	"time"
)

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type Float64Array []float64

func (f Float64Array) Len() int {
	return len(f)
}

func (f Float64Array) Less(i, j int) bool {
	return f[i] < f[j]
}

func (f Float64Array) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f Float64Array) List() string {
	output := "["

	for i := 0; i < f.Len(); i++ {
		output = fmt.Sprintf("%s%2.3f, ", output, f[i])
	}

	output = fmt.Sprintf("%s%s", output[:len(output)-2], "]")

	return output
}

func (f Float64Array) String() string {
	return f.List()
}

func Fill() Float64Array {
	var f Float64Array
	// 没有考虑到的点
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < 10; i++ {
		f = append(f, rand.Float64())
	}

	return f
}

func NewFloat64Array() Float64Array {
	var f Float64Array
	for i := 0; i < 25; i++ {
		f = append(f, float64(i))
	}

	return f
}

func Sort(data Sorter) {
	for i := 1; i < data.Len(); i++ {
		for j := 0; j < data.Len()-1; j++ {
			if data.Less(j+1, j) {
				data.Swap(j, j+1)
			}
		}
	}
}

func IsSorted(data Sorter) bool {
	for i := 0; i < data.Len()-1; i++ {
		if data.Less(i+1, i) {
			return false
		}
	}

	return true
}
