package convert

// ToPtr converts passed value to pointer to it.
func ToPtr[T any](value T) *T {
	return &value
}
