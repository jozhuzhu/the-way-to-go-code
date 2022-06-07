package _interface

import "fmt"

func mapFunc_1(mf func(obj) obj, list ...obj) []obj {
	results := make([]obj, len(list))

	for i, v := range list {
		results[i] = mf(v)
	}

	return results
}

func TestMap_1() {
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
	ints = mapFunc_1(mf, ints...)
	fmt.Println(ints)

	fmt.Println(strings)
	strings = mapFunc_1(mf, strings...)
	fmt.Println(strings)
}
