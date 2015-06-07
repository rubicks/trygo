/* github.com/rubicks/trygo/dumbmath/sqrt_test.go */

package dumbmath

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	cases := []struct {
		in, want float64
	}{
		{1.0, 1.0},
		{4.0, 2.0},
		{9.0, 3.0},
		{16.0, 4.0},
		{25.0, 5.0},
	}
	for _, c := range cases {
		got := Sqrt(c.in)
		if 0.001 < math.Abs(got-c.want) {
			t.Errorf("Sqrt(%#v) == %#v, want %#v", c.in, got, c.want)
		}
	}
}
