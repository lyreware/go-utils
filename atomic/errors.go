package atomic

import (
	"errors"
)

var (
	// ErrTargetIsNil is an error when atomic target is nil.
	ErrTargetIsNil = errors.New("target is nil")

	// ErrExternal is an external error during atomic operation.
	ErrExternal = errors.New("external")
)
