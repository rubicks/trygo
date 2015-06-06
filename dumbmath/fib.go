/* github.com/rubicks/trygo/dumbmath/fib.go */

package dumbmath

import (
	"math/big"
)

func Decr(n *big.Int) *big.Int {
	return n.Sub(n, big.NewInt(1))
}

func Fib(n *big.Int) *big.Int {
	ret := big.NewInt(0)
	nex := big.NewInt(1)
	for ; 0 < n.Sign(); n = Decr(n) {
		ret, nex = nex, new(big.Int).Add(ret, nex)
	}
	return ret
}
