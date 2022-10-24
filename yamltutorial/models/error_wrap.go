package models

import (
	"fmt"
)

func (e *Error) Error() string {
	return fmt.Sprintf("error: %s", e.Message)
}

func (e *Error) Unwrap() error {
	return e.Message
}

/*
 * context error
 */
func (e *ContextError) Error() string {
	return fmt.Sprintf("error: %s", e.Err)
}

func (e *ContextError) Unwrap() error {
	return e.Err
}

func (e *ContextError) GetContext() string {
	return e.Context
}

/*
 * composite error
 */
func (e *CompositeError) Error() string {

}
