package _interface

type Miner interface {
	ElemAt(i int) interface{}
	Less(i, j int) bool
	Len() int
	Swap(i, j int)
}

func Min(data Miner) interface{} {
	min := data.ElemAt(0)
	for i := 1; i < data.Len(); i++ {
		if data.Less(i, i-1) {
			min = data.ElemAt(i)
		} else {
			data.Swap(i, i-1)
		}
	}

	return min
}

type IntArray []int

func (I IntArray) Len() int {
	return len(I)
}
func (I IntArray) Less(i, j int) bool {
	return I[i] < I[j]
}
func (I IntArray) Swap(i, j int) {
	I[i], I[j] = I[j], I[i]
}
func (I IntArray) ElemAt(i int) interface{} {
	return I[i]
}

type StringArray []string

func (s StringArray) Len() int {
	return len(s)
}
func (s StringArray) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s StringArray) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s StringArray) ElemAt(i int) interface{} {
	return s[i]
}
