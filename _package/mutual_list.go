package _package

import (
	"container/list"
	"fmt"
)

func PrintMutualList() {
	l := list.New()

	l.PushBack(101)
	l.PushBack(102)
	l.PushBack(103)

	//length := l.Len()
	//
	//for i := 0; i < length; i++ {
	//	e := l.Front()
	//	l.Remove(e)
	//
	//	fmt.Println(e.Value)
	//}
	for e := l.Front(); e != nil; e = e.Next() {
		// fmt.Println(e)
		fmt.Println(e.Value)
	}
}
