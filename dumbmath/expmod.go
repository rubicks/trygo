// github.com/rubicks/trygo/dumbmath/expmod.go

package dumbmath

func Expmod(b, e, m int) int {
	ret := 1
	for 0 < e {
		if 1 == e%2 {
			ret *= b
			ret %= m
		}
		b *= b
		b %= m
		e >>= 1
	}
	return ret
}
