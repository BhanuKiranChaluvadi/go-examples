package main

import (
	"fmt"
	"io"
	"os"
	"yamltutorial/models"

	"github.com/go-openapi/strfmt"
	"gopkg.in/yaml.v3"
)

func main() {
	// Open yaml File
	yamlFile, err := os.Open("simple-manifest.yaml")
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
}
