package gcs

import (
	"fmt"
)

type containerExistsError struct {
	ID string
}

func (e *containerExistsError) Error() string {
	return fmt.Sprintf("a container with the ID \"%s\" already exists", e.ID)
}

// NewContainerExistsError returns a *containerExistsError referring to the
// given ID.
func NewContainerExistsError(id string) *containerExistsError {
	return &containerExistsError{ID: id}
}

type containerDoesNotExistError struct {
	ID string
}

func (e *containerDoesNotExistError) Error() string {
	return fmt.Sprintf("a container with the ID \"%s\" does not exist", e.ID)
}

// NewContainerDoesNotExistError returns a *containerDoesNotExistError
// referring to the given ID.
func NewContainerDoesNotExistError(id string) *containerDoesNotExistError {
	return &containerDoesNotExistError{ID: id}
}

type processDoesNotExistError struct {
	Pid int
}

func (e *processDoesNotExistError) Error() string {
	return fmt.Sprintf("a process with the pid %d does not exist", e.Pid)
}

// NewProcessDoesNotExistError returns a *processDoesNotExistError referring to
// the given pid.
func NewProcessDoesNotExistError(pid int) *processDoesNotExistError {
	return &processDoesNotExistError{Pid: pid}
}
