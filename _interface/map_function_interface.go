package _interface

import "fmt"

type obj interface{}

func mapFunc(mf func(obj) obj, data []obj) []obj {
	//for i, v := range data {
	//	data[i] = mf(v)
	//}
	//
	//return data

	result := make([]obj, len(data))

	for ix, v := range data {
		result[ix] = mf(v)
	}

	return result
}

func TestMap() {
	mf := func(i obj) obj {
		//switch t := o.(type) {
		//case int:
		//	return 2 * t
		//
		//case string:
		//	return t + t
		//}
		//return nil
		switch i.(type) {
		case int:
			return i.(int) * 2
		case string:
			return i.(string) + i.(string)
		}
		return i
	}

	ints := []obj{1, 4, 8, 3, 0, 7, 2}
	strings := []obj{"abc", "def", "love", "丹琦"}

	fmt.Println(ints)
	ints = mapFunc(mf, ints)
	fmt.Println(ints)

	fmt.Println(strings)
	strings = mapFunc(mf, strings)
	fmt.Println(strings)
}
