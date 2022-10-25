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
	return ""
}

func PrintRandomAPIError() *APIError {
	return &APIError{
		Trace:      "9daee671-916a-4678-850b-10b911f0236d",
		StatusCode: 500,
		TranslationCode: &TranslationCode{
			Major: 1001,
			Minor: 2005,
		},
		APIContext:  "installation_failed",
		Description: "There is a miss-match of pattern",
		Resolve:     "Please download urcapx file again and try installing again or contact urcap developer",
		MoreInfo:    "https://docs.api.example.com/v2/users/create_user#first_name",
		Errors: []*Error{
			{
				Code:     "invalid_manifest",
				Message:  fmt.Errorf("the urcapID field is required in manifest"),
				MoreInfo: "https://docs.api.example.com/v2/users/create_user#first_name",
				Target: &ErrorTarget{
					Name:  "field",
					Value: "urcapID",
				},
			},
		},
	}
}

/*
&{9daee671-916a-4678-850b-10b911f0236d 500 0xc000314f60 installation_failed There is a miss-match of pattern Please download urcapx file again and try installing again or contact urcap developer https://docs.api.example.com/v2/users/create_user#first_name [error: the urcapID field is required in manifest]}
*/
