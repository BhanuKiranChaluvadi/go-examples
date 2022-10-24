package main

import (
	"errors"
	"fmt"
	load "yamltutorial/loader"
	"yamltutorial/models"

	yamlErrors "yamltutorial/errors"

	openapiErrors "github.com/go-openapi/errors"
)

var (
	// manifestFile string = "simple-manifest.yaml"
	manifestFile string = "simple-error-manifest.yaml"
)

func main() {

	_, err := load.Load(manifestFile)

	// try and unwrap
	fmt.Println("Unwrapping validation error")
	var validationErrs *openapiErrors.CompositeError
	if errors.As(err, &validationErrs) {
		fmt.Println("validationErrs code: ", validationErrs.Code())
		// one item
		for _, err := range validationErrs.Errors {
			var compositeErrs *openapiErrors.CompositeError
			if errors.As(err, &compositeErrs) {
				fmt.Println("compositeErrs code: ", compositeErrs.Code())
			}
			// multiple items
			for _, err := range compositeErrs.Errors {
				var validationErr *openapiErrors.Validation
				if errors.As(err, &validationErr) {
					var code string
					if openapiErrors.RequiredFailCode == validationErr.Code() {
						code = yamlErrors.RequiredFailCode
					} else if openapiErrors.PatternFailCode == validationErr.Code() {
						code = yamlErrors.PatternFailCode
					}
					modelError := models.Error{
						Code:     code,
						Message:  validationErr.Error(),
						MoreInfo: "",
						Target: &models.ErrorTarget{
							Name: &validationErr.Name,
							Type: validationErr.In,
						},
					}
					fmt.Printf("%+v\n", modelError)
					byteValue, err := validationErr.MarshalJSON()
					if err != nil {
						fmt.Println("Marshal failed")
					}
					fmt.Println(string(byteValue))
				}
			}

		}
	}
}
