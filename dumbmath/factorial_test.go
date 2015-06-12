/* github.com/rubicks/trygo/dumbmath/factorial_test.go */

package dumbmath

import (
	"math/big"
	"testing"
)

func TestFactorial(t *testing.T) {
	cases := []struct {
		in, want *big.Int
	}{
		{big.NewInt(0), big.NewInt(1)},
		{big.NewInt(1), big.NewInt(1)},
		{big.NewInt(2), big.NewInt(2)},
		{big.NewInt(3), big.NewInt(6)},
		{big.NewInt(4), big.NewInt(24)},
	}
	for _, c := range cases {
		got := Factorial(c.in)
		if got.Cmp(c.want) != 0 {
			t.Errorf("Factorial(%#v) == %#v, want %#v", c.in, got, c.want)
		}
	}
}
