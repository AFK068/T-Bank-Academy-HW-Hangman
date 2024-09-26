package apperrors

import "errors"

var wrappedErr interface{ Unwrap() error }

func UnwrapError(err error) error {
	if errors.As(err, &wrappedErr) {
		for unwrappedErr := errors.Unwrap(err); unwrappedErr != nil; unwrappedErr = errors.Unwrap(err) {
			err = unwrappedErr
		}
	}

	return err
}
