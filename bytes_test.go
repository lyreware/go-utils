package utils

import (
	"slices"
	"testing"
)

func TestCloneByteSlice(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		slice []byte
	}{
		{"nil", nil},
		{"empty", []byte{}},
		{"make with zero length", make([]byte, 0, 10)},
		{"make with non zero length", make([]byte, 4, 10)},
		{"init fixed slice", []byte{1, 2, 3, 4}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			clone := CloneByteSlice(test.slice)
			if !slices.Equal(clone, test.slice) {
				t.Fatalf("CloneByteSlice(%v) returned %v", test.slice, clone)
			}

			if (test.slice == nil && clone != nil) || (test.slice != nil && clone == nil) {
				t.Fatalf("CloneByteSlice(%v): got %v", test.slice, clone)
			}

			if cap(clone) > len(test.slice) {
				t.Fatalf("CloneByteSlice(%v) expected not greater than %d but got %d", test.slice, len(test.slice), cap(clone))
			}

			if len(clone) > 0 && &test.slice[0] == &clone[0] {
				t.Fatalf("CloneByteSlice(%v) returned the same slice %v", test.slice, &clone[0])
			}
		})
	}
}

func TestConcatByteSlices(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		slices [][]byte
		result []byte
	}{
		{
			"many slices",
			[][]byte{[]byte{1, 2, 3}, []byte{4, 5, 6}, []byte{7, 8, 9}},
			[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{"nil", nil, nil},
		{"many nil slices", [][]byte{nil, nil, nil, nil}, nil},
		{"empty slice of slices", [][]byte{}, nil},
		{"slice of slices", [][]byte{[]byte{}, []byte{}, []byte{}}, nil},
		{
			"slices and empty slices",
			[][]byte{[]byte{}, []byte{1}, []byte{2, 3}, []byte{}},
			[]byte{1, 2, 3},
		},
		{
			"slices and nil slices",
			[][]byte{nil, []byte{1, 2}, nil, []byte{3}},
			[]byte{1, 2, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			result := ConcatByteSlices(test.slices...)
			if !slices.Equal(result, test.result) {
				t.Fatalf("ConcatByteSlices(%v): expected %v but got %v", test.slices, test.result, result)
			}

			if cap(result) > len(test.result) {
				t.Fatalf(
					"ConcatByteSlices(%v): expected capacity not greater than %d but got %d",
					test.slices,
					cap(test.result),
					cap(result),
				)
			}
		})
	}
}
