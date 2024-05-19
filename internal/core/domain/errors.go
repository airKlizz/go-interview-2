package domain

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrorEventNotValid      = ErrEventNotValid{}                    // event is not valid
	ErrorDeviceNotSupported = errors.New("device not supported")    // device not supported
	ErrorDeviceNotFound     = errors.New("device target not found") // device target not found
	ErrorEventFailed        = errors.New("failed to handle event")  // failed to handle event
)

type ErrEventNotValid struct {
	reasons []string
}

func NewErrEventNotValid(reasons []string) error {
	return ErrEventNotValid{
		reasons: reasons,
	}
}

func (err ErrEventNotValid) Error() string {
	return fmt.Sprintf("event not valid: %s", strings.Join(err.reasons, ", "))
}

func (err ErrEventNotValid) Is(target error) bool {
	_, ok := target.(ErrEventNotValid)
	return ok
}
