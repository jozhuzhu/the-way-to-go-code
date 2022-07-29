package _goruntine

import "fmt"

//type Any interface {}

//type reuseFunc func() Any
//
//func BuildLazyEvaluator(reuse reuseFunc) reuseFunc {
//	result := make(chan Any)
//	loopFunc := func() {
//		for {
//			result <- reuse()
//		}
//	}
//
//	returnFunc := func() Any {
//		return <-result
//	}
//
//	go loopFunc()
//	return returnFunc
//}
//
//func TestLazyEvaluator() {
//	x, y := 0, 1
//	fibonacci := func() Any {
//		x, y = y, x + y
//		return x
//	}
//	f := BuildLazyEvaluator(fibonacci)
//	for i := 0; i < 10; i++ {
//		fmt.Println(f().(int))
//	}
//}

// ===========================================================================================================
type Any interface{}
type EvalFunc func(Any) (Any, Any)

func TestLazyEvaluator() {
	fibFunc := func(state Any) (Any, Any) {
		os := state.([]uint64)
		v1 := os[0]
		v2 := os[1]
		ns := []uint64{v2, v1 + v2}
		return v1, ns
	}
	fib := BuildLazyUInt64Evaluator(fibFunc, []uint64{0, 1})

	for i := 0; i < 10; i++ {
		fmt.Printf("Fib nr %v: %v\n", i, fib())
	}
}

func BuildLazyEvaluator(evalFunc EvalFunc, initState Any) func() Any {
	retValChan := make(chan Any)
	loopFunc := func() {
		var actState Any = initState
		var retVal Any
		for {
			retVal, actState = evalFunc(actState)
			retValChan <- retVal
		}
	}
	retFunc := func() Any {
		return <-retValChan
	}
	go loopFunc()
	return retFunc
}

func BuildLazyUInt64Evaluator(evalFunc EvalFunc, initState Any) func() uint64 {
	ef := BuildLazyEvaluator(evalFunc, initState)
	return func() uint64 {
		return ef().(uint64)
	}
}
