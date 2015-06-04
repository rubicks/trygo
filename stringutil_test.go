/* ~/go/src/github.com/rubicks/trygo/stringutil.go */

package stringutil

import (
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"foobar", "raboof"},
		{"deified", "deified"},
		{"racecar", "racecar"},
		{"repaper", "repaper"},
		{"reviver", "reviver"},
		{"rotavator", "rotavator"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
