package utils

import "testing"

func TestIsNil(t *testing.T) {
	t.Parallel()

	intValue := 123

	tests := []struct {
		name   string
		arg    any
		result bool
	}{
		{"nil", nil, true},
		{"nil map", map[string]string(nil), true},
		{"map", map[string]string{"hello": "world"}, false},
		{"nil ptr", (*int)(nil), true},
		{"int", 5, false},
		{"string", "string", false},
		{"nil slice", []int(nil), true},
		{"slice", []int{1, 2, 3}, false},
		{"ptr", &intValue, false},
		{"nil chan", (chan int)(nil), true},
		{"chan", make(chan int), false},
		{"nil func", (func())(nil), true},
		{"func", TestIsNil, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if result := IsNil(test.arg); result != test.result {
				t.Fatalf("IsNil(%v) expected %t but got %t", test.arg, test.result, result)
			}
		})
	}
}
