// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Error struct {
	Message string `json:"message,omitempty"`
	Err     error  `json:"err,omitempty"`
}

func (e *Error) Error() string {
	byteArray, _ := json.Marshal(e)
	return string(byteArray)
}

func generateError() (err error) {
	defer func() {
		if err != nil {
			err = &Error{
				Message: "I am the message",
				Err:     err,
			}
		}
	}()
	err = errors.New("i am the error")
	return err
}

func main() {

	err := generateError()
	// fmt.Printf("%+v", err)
	fmt.Printf("%s", err)
}
