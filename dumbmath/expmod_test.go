// github.com/rubicks/trygo/dumbmath/expmod_test.go

package dumbmath

import (
	"testing"
)

func TestExpmod(t *testing.T) {
	type bem [3]int
	cases := []struct {
		in   bem
		want int
	}{
		// prime powers of two
		{bem{2, 2, 100}, 4},
		{bem{2, 3, 100}, 8},
		{bem{2, 5, 100}, 32},
		{bem{2, 7, 100}, 28},
		{bem{2, 11, 100}, 48},
		{bem{2, 13, 100}, 92},
		// carmichael
		{bem{3, 560, 561}, 375},
		{bem{11, 560, 561}, 154},
		{bem{17, 560, 561}, 34},
	}
	for _, c := range cases {
		got := Expmod(c.in[0], c.in[1], c.in[2])
		if got != c.want {
			t.Errorf("Expmod(%#v, %#v, %#v) == %#v, want %#v",
				c.in[0], c.in[1], c.in[2], got, c.want)
		}
	}

	primes := []int{3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47}
	for _, p := range primes {
		got := Expmod(2, p-1, p)
		if got != 1 {
			t.Errorf("Expmod(2,%#v,%#v) == %#v, want %#v", p-1, p, got, 1)
		}
	}
}
