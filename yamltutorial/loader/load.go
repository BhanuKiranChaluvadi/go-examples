package load

import (
	"errors"
	"fmt"
	"io"
	"os"
	"yamltutorial/models"

	yamlErrors "yamltutorial/errors"

	openapiErrors "github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"gopkg.in/yaml.v3"
)

func Load(filePath string) (models.Manifest, error) {
	// Open yaml File
	yamlFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Open error:", err)
	}
	defer yamlFile.Close()

	// read opened yamlFile as a byte array.
	byteValue, err := io.ReadAll(yamlFile)
	if err != nil {
		fmt.Println("Read error:", err)
	}

	var manifest models.Manifest
	// we unmarshal our byteArray which contains our
	// yamlFile's content into 'manifest' model
	err = yaml.Unmarshal(byteValue, &manifest)
	if err != nil {
		fmt.Println("yaml unmarshal error:", err)
	}

	format := strfmt.Default
	err = manifest.Validate(format)
	if err != nil {
		fmt.Println("manifest validation error:", err)
	}

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
