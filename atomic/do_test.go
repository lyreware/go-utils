package atomic

import (
	"errors"
	"testing"
)

func TestDoFuncError(t *testing.T) {
	t.Parallel()

	value := 123
	externalErr := errors.New("external error")

	err := Do(&value, value, func(dirty *int) error {
		*dirty = 456

		return externalErr
	})
	if !errors.Is(err, ErrExternal) || !errors.Is(err, externalErr) || value != 123 {
		t.Fatalf("Do returned %d, %v", value, err)
	}
}

func TestDoSuccess(t *testing.T) {
	t.Parallel()

	value := 123

	err := Do(&value, value, func(x *int) error {
		*x *= 2

		return nil
	})
	if err != nil || value != 246 {
		t.Fatalf("Do returned %d, %v", value, err)
	}
}

func TestDoTargetIsNil(t *testing.T) {
	t.Parallel()

	err := Do(nil, 5, func(_ *int) error {
		return nil
	})
	if !errors.Is(err, ErrTargetIsNil) {
		t.Fatalf("Do(nil, 5) expected nil target error but got %v", err)
	}
}
