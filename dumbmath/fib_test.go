/* github.com/rubicks/trygo/dumbmath/fib_test.go */

package dumbmath

import (
	"math/big"
	"testing"
)

func TestFib(t *testing.T) {
	cases := []struct {
		in, want *big.Int
	}{
		{big.NewInt(0), big.NewInt(0)},
		{big.NewInt(1), big.NewInt(1)},
		{big.NewInt(2), big.NewInt(1)},
		{big.NewInt(3), big.NewInt(2)},
		{big.NewInt(4), big.NewInt(3)},
		{big.NewInt(5), big.NewInt(5)},
		{big.NewInt(6), big.NewInt(8)},
		{big.NewInt(7), big.NewInt(13)},
		{big.NewInt(8), big.NewInt(21)},
		{big.NewInt(9), big.NewInt(34)},
		{big.NewInt(10), big.NewInt(55)},
		{big.NewInt(11), big.NewInt(89)},
		{big.NewInt(13), big.NewInt(233)},
		{big.NewInt(17), big.NewInt(1597)},
		{big.NewInt(19), big.NewInt(4181)},
	}
	for _, c := range cases {
		got := Fib(c.in)
		if got.Cmp(c.want) != 0 {
			t.Errorf("Fib(%#v) == %#v, want %#v", c.in, got, c.want)
		}
	}
}
