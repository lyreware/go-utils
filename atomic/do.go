package atomic

import (
	"errors"
)

// DoFunc is a function that operates on a dirty value.
type DoFunc[T any] func(dirty *T) error

// Do passes the dirty value to the function and, if the function returned without
// errors, replaces the target with the dirty version. This is useful for use in public
// methods to avoid corrupting the state with errors.
func Do[T any](target *T, dirty T, fn DoFunc[T]) error {
	err := fn(&dirty)
	if err != nil {
		return errors.Join(ErrExternal, err)
	}

	if target == nil {
		return ErrTargetIsNil
	}

	*target = dirty

	return nil
}
