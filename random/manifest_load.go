package cli

import (
	"io"
	"os"
	"universal-robots/urcapctl/models"

	"github.com/go-openapi/strfmt"
	"gopkg.in/yaml.v2"
)

func manifestLoad(filepath string) (manifest *models.Manifest, err error) {

	// Open yaml File
	yamlFile, err := os.Open(filepath)
	if err != nil {
		return
	}

	// read yamlFile as a byte array.
	byteValue, err := io.ReadAll(yamlFile)
	if err != nil {
		return
	}

	// unmarshal our byteArray into 'manifest' model
	err = yaml.Unmarshal(byteValue, manifest)
	if err != nil {
		return
	}

	err = manifest.Validate(strfmt.Default)
	if err != nil {
		return
	}

	return
}
