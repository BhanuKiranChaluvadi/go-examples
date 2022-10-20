package models

import (
	"fmt"
	"strings"
)

func (e *Error) Error() string {
	return fmt.Sprintf("error: %s", e.Message)
}

func (e *Error) Unwrap() error {
	return e.Message
}

func (c *CompositeError) Error() string {
	if len(c.Errors) > 0 {
		msgs := []string{c.Message + ":"}
		for _, e := range c.Errors {
			msgs = append(msgs, e.Error())
		}
		return strings.Join(msgs, "\n")
	}
	return c.message
}

func (c *CompositeError) Unwrap() error {
	return nil
}
