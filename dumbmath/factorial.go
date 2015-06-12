// github.com/rubicks/trygo/dumbmath/factorial.go

package dumbmath

import (
	"math/big"
)

// func zero() *big.Int {
// 	return big.NewInt(0)
// }

// func one() *big.Int {
// 	return big.NewInt(1)
// }

// func copy(n *big.Int) *big.Int {
// 	return zero().Set(n)
// }

// func add(a, b *big.Int) *big.Int {
// 	return zero().Add(a, b)
// }

// func incr(n *big.Int) *big.Int {
// 	return add(copy(n), one())
// }

// Factorial returns the factorial of its argument, calculated via repeated
// multiplication
func Factorial(n *big.Int) *big.Int {
	one := big.NewInt(1)
	ret := big.NewInt(1)
	for ; 0 < n.Sign(); n.Sub(n, one) {
		ret.Mul(ret, n)
	}
	return ret
}
