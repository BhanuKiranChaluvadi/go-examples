package models

import "fmt"

func (e *Error) Error() string {
	return fmt.Sprintf("error: %s", e.Message)
}

func (e *Error) Unwrap() error {
	return e.Message
}
