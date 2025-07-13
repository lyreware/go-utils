package convert

import (
	"testing"
)

var ToPtrTests = []struct {
	value any
}{
	{5},
	{-12.34},
	{"hello"},
}

func TestToPtr(t *testing.T) {
	t.Parallel()

	for _, test := range ToPtrTests {
		ptr := ToPtr(test.value)
		if ptr == nil || *ptr != test.value {
			t.Fatalf("ToPtr(%+v) returned %+v", test.value, ptr)
		}
	}
}
