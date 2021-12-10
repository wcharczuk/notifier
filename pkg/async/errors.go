package async

import (
	"fmt"
	"strings"
)

// Errors is a channel for errors
type Errors chan error

// First returns the first (non-nil) error.
func (e Errors) First() error {
	if errorCount := len(e); errorCount > 0 {
		var err error
		for x := 0; x < errorCount; x++ {
			err = <-e
			if err != nil {
				return err
			}
		}
		return nil
	}
	return nil
}

// All returns all the non-nil errors in the channel
// as a multi-error.
func (e Errors) All() error {
	if errorCount := len(e); errorCount > 0 {
		var errors []error
		for x := 0; x < errorCount; x++ {
			err := <-e
			if err != nil {
				errors = append(errors, err)
			}
		}
		if len(errors) > 0 {
			return MultiError(errors)
		}
		return nil
	}
	return nil
}

// MultiError is an array of errors.
type MultiError []error

// Error implements error.
func (me MultiError) Error() string {
	if len(me) == 0 {
		return ""
	}
	if len(me) == 1 {
		return me[0].Error()
	}
	lines := []string{
		fmt.Sprintf("%d errors occurred", len(me)),
	}
	for _, err := range me {
		lines = append(lines, fmt.Sprintf("\t%+v", err))
	}
	return strings.Join(lines, "\n")
}
