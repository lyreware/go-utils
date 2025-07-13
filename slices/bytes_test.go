package slices

import (
	"slices"
	"testing"
)

var cloneBytesTests = []struct {
	name   string
	origin []byte
}{
	{"nil", nil},
	{"empty", []byte{}},
	{"make with zero length", make([]byte, 0, 10)},
	{"make with non zero length", make([]byte, 4, 10)},
	{"init fixed slice", []byte{1, 2, 3, 4}},
}

func TestCloneBytes(t *testing.T) {
	t.Parallel()

	for _, test := range cloneBytesTests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			clone := CloneBytes(test.origin)
			if !slices.Equal(clone, test.origin) {
				t.Fatalf("CloneBytes(%v) returned %v", test.origin, clone)
			}

			if (test.origin == nil && clone != nil) ||
				(test.origin != nil && clone == nil) {
				t.Fatalf("CloneBytes(%v): got %v", test.origin, clone)
			}

			if cap(clone) > len(test.origin) {
				t.Fatalf(
					"CloneBytes(%v) expected not greater than %d but got %d",
					test.origin,
					len(test.origin),
					cap(clone),
				)
			}

			if len(clone) > 0 && &test.origin[0] == &clone[0] {
				t.Fatalf(
					"CloneBytes(%v) returned the same slice %v",
					test.origin,
					&clone[0],
				)
			}
		})
	}
}

var concatBytesTests = []struct {
	name   string
	slices [][]byte
	concat []byte
}{
	{
		"many slices",
		[][]byte{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
	{"nil", nil, nil},
	{
		"many nil slices",
		[][]byte{nil, nil, nil, nil},
		nil,
	},
	{"empty slice of slices", [][]byte{}, nil},
	{
		"slice of slices",
		[][]byte{{}, {}, {}},
		nil,
	},
	{
		"slices and empty slices",
		[][]byte{{}, {1}, {2, 3}, {}},
		[]byte{1, 2, 3},
	},
	{
		"slices and nil slices",
		[][]byte{nil, {1, 2}, nil, {3}},
		[]byte{1, 2, 3},
	},
}

func TestConcatBytes(t *testing.T) {
	t.Parallel()

	for _, test := range concatBytesTests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			concat := ConcatBytes(test.slices...)
			if !slices.Equal(concat, test.concat) {
				t.Fatalf(
					"ConcatBytes(%v): expected %v but got %v",
					test.slices,
					test.concat,
					concat,
				)
			}

			if cap(concat) > len(test.concat) {
				t.Fatalf(
					"ConcatBytes(%v): expected capacity not greater than %d but got %d",
					test.slices,
					cap(test.concat),
					cap(concat),
				)
			}
		})
	}
}
