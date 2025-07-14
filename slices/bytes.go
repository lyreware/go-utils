package slices

// CloneBytes clones a slice of bytes, as slices.Clone would do, but does not
// allocate extra space.
func CloneBytes(origin []byte) []byte {
	if origin == nil {
		return nil
	}

	clone := make([]byte, len(origin))
	copy(clone, origin)

	return clone
}

// ConcatBytes concatenates a byte slices, as slices.Concat would do, but does
// not allocate extra space.
func ConcatBytes(slices ...[]byte) []byte {
	totalLen := 0
	for _, slice := range slices {
		totalLen += len(slice)
	}

	if totalLen == 0 {
		return nil
	}

	concat := make([]byte, 0, totalLen)
	for _, slice := range slices {
		concat = append(concat, slice...)
	}

	return concat
}
