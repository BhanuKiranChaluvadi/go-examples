package load

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"syscall"
	"yamltutorial/models"

	yamlErrors "yamltutorial/errors"

	openapiErrors "github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"gopkg.in/yaml.v3"
)

func ManifestLoad(filePath string) (manifest *models.Manifest, err error) {
	// Open yaml File
	yamlFile, err := os.Open(filePath)
	if err != nil {
		var e *os.PathError
		// convert to the corresponsing error type
		if errors.As(err, &e) {
			if errors.Is(e, fs.ErrNotExist) {
				e := models.Error{
					Code:     yamlErrors.ErrInvalidFile,
					Message:  err,
					MoreInfo: "",
					Target: &models.ErrorTarget{
						Type:  e.Op,
						Value: e.Path,
					},
				}
			} else if errors.Is(err, fs.ErrInvalid) {

			} else {

			}
		}

	}

	if errors.Is(err, syscall.ENOSPC) {

	} else if errors.Is(err, syscall.RTM_NR_MSGTYPES) {

	}
	if err != nil {
		if errors.Is(err, ERREOF) {
			return nil, &models.CompositeError{
				Context: yamlErrors.ErrManifestOpen,
				Code:    0,
				Errors: []*models.Error{
					&models.Error{
						Code:     "",
						Message:  err,
						MoreInfo: "",
						Target:   &models.ErrorTarget{},
					},
				},
				Message: "",
			}
		} else {

		}

		return nil, &models.ContextError{
			Context: yamlErrors.ErrManifestRead,
			Err: &models.Error{
				Code:     "",
				Message:  err,
				MoreInfo: "",
				Target:   &models.ErrorTarget{},
			},
		}
	}
	defer yamlFile.Close()

	// read opened yamlFile as a byte array.
	byteValue, err := io.ReadAll(yamlFile)
	if err != nil {
		return nil, &models.ContextError{
			Context: yamlErrors.ErrManifestRead,
			Err:     err,
		}
	}

	// var manifest models.Manifest
	// we unmarshal our byteArray which contains our
	// yamlFile's content into 'manifest' model
	err = yaml.Unmarshal(byteValue, manifest)
	if err != nil {
		return nil, &models.ContextError{
			Context: yamlErrors.ErrFailedMarshal,
			Err:     err,
		}
	}

	// validate
	err = Validate(manifest)
	if err != nil {
		return nil, &models.ContextError{
			Context: yamlErrors.ErrManifestInvalid,
			Err:     err,
		}
	}
	return manifest, nil
}

func Validate(manifest *models.Manifest) (err error) {
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
	manifest.Validate()
	return nil
}

func formatErrors() (err error) {

	return &models.Error{
		Code:     "",
		Message:  nil,
		MoreInfo: "",
		Target:   &models.ErrorTarget{},
	}
}

func random() *models.Error {
	return &models.Error{}
}
