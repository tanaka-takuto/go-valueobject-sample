package user

import "errors"

// NewErrCannotDeleteYourself creates an error that occurs when a user tries to delete themselves
func NewErrCannotDeleteYourself() error {
	return errors.New("cannot delete yourself")
}
