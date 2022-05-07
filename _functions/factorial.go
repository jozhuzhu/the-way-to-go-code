package _functions

import (
	"math/big"
)

/*func main() {
	// normal
	var num int64 = 30
	ret := defer_trace.Factorial_1(new(big.Int).SetInt64(num))
	fmt.Printf("The factorial of %d is %s\n", num, ret.String())

	// integer overflow
	intret := defer_trace.Factorial_2(num)
	fmt.Printf("The factorial of %d is %d\n", num, intret)
}*/

func Factorial_1(n *big.Int) (ret *big.Int) {
	if n.Cmp(new(big.Int).SetInt64(1)) == 0 {
		return new(big.Int).SetInt64(1)
	}
	oldN := new(big.Int)
	oldN = oldN.Set(n)
	return n.Mul(oldN, Factorial_1(n.Sub(n, new(big.Int).SetInt64(1))))
}

// 30! 已经超过 int64 的数据范围
func Factorial_2(n int64) (ret int64) {
	if n == 1 {
		return 1
	}

	return Factorial_2(n-1) * n
}
