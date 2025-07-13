package atomic_test

import (
	"errors"
	"testing"

	"github.com/lyreware/go-utils/atomic"
)

func TestDoFuncError(t *testing.T) {
	t.Parallel()

	value := 123

	err := atomic.Do(&value, value, func(x *int) error {
		*x *= 123

		return errors.New("runtime error")
	})
	if err == nil || err.Error() != "runtime error" || value != 123 {
		t.Fatalf("Do returned %d, %v", value, err)
	}
}

func TestDoSuccess(t *testing.T) {
	t.Parallel()

	value := 123

	err := atomic.Do(&value, value, func(x *int) error {
		*x *= 2

		return nil
	})
	if err != nil || value != 246 {
		t.Fatalf("Do returned %d, %v", value, err)
	}
}

func TestDoTargetIsNil(t *testing.T) {
	t.Parallel()

	err := atomic.Do(nil, 5, func(_ *int) error {
		return nil
	})
	if err == nil || err.Error() != "target is nil" {
		t.Fatalf(
			"Do(nil, 5) expected \"target is nil\" error but got %v",
			err,
		)
	}
}
