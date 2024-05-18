package domain

import (
	"fmt"

	cue "cuelang.org/go/cue/errors"
)

var (
	ErrorEventNotValid = ErrEventNotValid{} // event is not valid
)

type ErrEventNotValid struct {
	message string
}

func NewErrEventNotValid(err error) error {
	return ErrEventNotValid{
		message: cue.Details(err, nil),
	}
}

func (err ErrEventNotValid) Error() string {
	return fmt.Sprintf("event not valid: %s", err.message)
}

func (err ErrEventNotValid) Is(target error) bool {
	_, ok := target.(ErrEventNotValid)
	return ok
}
