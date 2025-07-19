package check

import (
	"testing"

	"github.com/crypto-tribe/go-utils/convert"
)

var isNilTests = []struct {
	name  string
	arg   any
	isNil bool
}{
	{"nil", nil, true},
	{"nil map", map[string]string(nil), true},
	{"map", map[string]string{"hello": "world"}, false},
	{"nil ptr", (*int)(nil), true},
	{"int", 5, false},
	{"string", "string", false},
	{"nil slice", []int(nil), true},
	{"slice", []int{1, 2, 3}, false},
	{"ptr", convert.ToPtr(123), false},
	{"nil chan", (chan int)(nil), true},
	{"chan", make(chan int), false},
	{"nil func", (func())(nil), true},
	{"func", func() {}, false},
}

func TestIsNil(t *testing.T) {
	t.Parallel()

	for _, test := range isNilTests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			isNil := IsNil(test.arg)
			if isNil != test.isNil {
				t.Fatalf(
					"IsNil(%v) expected %t but got %t",
					test.arg,
					test.isNil,
					isNil,
				)
			}
		})
	}
}
