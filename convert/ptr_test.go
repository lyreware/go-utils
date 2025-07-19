package convert

import (
	"testing"
)

var toPtrTests = []struct {
	value any
}{
	{5},
	{-12.34},
	{"hello"},
}

func TestToPtr(t *testing.T) {
	t.Parallel()

	for _, test := range toPtrTests {
		ptr := ToPtr(test.value)
		if ptr == nil || *ptr != test.value {
			t.Fatalf("ToPtr(%+v) returned %+v", test.value, ptr)
		}
	}
}
