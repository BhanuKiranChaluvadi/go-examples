package main

/*
The errors.Is function allows you to check if a specific sentinel error value is anywhere inside a wrapped error.
The errors.As function allows you to get a reference to a certain type of error anywhere inside a wrapped error.
*/
import (
	"errors"
	"fmt"
)

var (
	errUhOh = errors.New("uh oh")
)

type ValueError struct {
	Value int
	Err   error
}

func newValueError(value int, err error) *ValueError {
	return &ValueError{
		Value: value,
		Err:   err,
	}
}

func (ve *ValueError) Error() string {
	return fmt.Sprintf("value error: %s", ve.Err)
}

func (ve *ValueError) Unwrap() error {
	return ve.Err
}

/*
func giveMeError() error {
	return errUhOh
}
*/

func validateValue(number int) error {
	if number == 1 {
		return newValueError(number, fmt.Errorf("that's odd"))
	} else if number == 2 {
		return newValueError(number, errUhOh)
	}
	return nil
}

func runValidation(number int) error {
	err := validateValue(number)
	if err != nil {
		return fmt.Errorf("run error: %w", err)
	}
	return nil
}

func main() {

	// errors.Is
	for num := 1; num <= 3; num++ {
		fmt.Printf("validating %d...", num)
		err := runValidation(num)
		// if errors.Is(err, errUhOh) {
		// 	fmt.Println("oh no!")
		// }

		// if err == errUhOh ||
		// errors.Unwrap(err) == errUhOh ||
		// errors.Unwrap(errors.Unwrap(err)) == errUhOh {
		// fmt.Println("oh no!")
		// }
		if errors.Is(err, errUhOh) {
			fmt.Println("oh no!")
		} else if err != nil {
			fmt.Println("there was an error:", err)
		} else {
			fmt.Println("valid!")
		}
	}

	// errors.As
	for num := 1; num <= 3; num++ {
		fmt.Printf("validating %d...", num)
		err := runValidation(num)

		var valueErr *ValueError
		if errors.Is(err, errUhOh) {
			fmt.Println("oh no!")
		} else if errors.As(err, &valueErr) {
			fmt.Printf("value error (%d): %v\n", valueErr.Value, valueErr.Err)
		} else if err != nil {
			fmt.Println("there was an error:", err)
		} else {
			fmt.Println("valid!")
		}

	}
}
