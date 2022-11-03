// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"swagger-example/models"

	"github.com/spf13/cobra"
)

// Schema cli for PolyscopeBundle

// register flags to command
func registerModelPolyscopeBundleFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPolyscopeBundleBundle(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPolyscopeBundleName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPolyscopeBundleBundle(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	bundleDescription := `Required. The polyscope bundle path. Path to the jar file relative to the zipped urcap top folder.`

	var bundleFlagName string
	if cmdPrefix == "" {
		bundleFlagName = "bundle"
	} else {
		bundleFlagName = fmt.Sprintf("%v.bundle", cmdPrefix)
	}

	var bundleFlagDefault string

	_ = cmd.PersistentFlags().String(bundleFlagName, bundleFlagDefault, bundleDescription)

	return nil
}

func registerPolyscopeBundleName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Required. The name of the polyscope bundle. 
`

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPolyscopeBundleFlags(depth int, m *models.PolyscopeBundle, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, bundleAdded := retrievePolyscopeBundleBundleFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || bundleAdded

	err, nameAdded := retrievePolyscopeBundleNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	return nil, retAdded
}

func retrievePolyscopeBundleBundleFlags(depth int, m *models.PolyscopeBundle, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	bundleFlagName := fmt.Sprintf("%v.bundle", cmdPrefix)
	if cmd.Flags().Changed(bundleFlagName) {

		var bundleFlagName string
		if cmdPrefix == "" {
			bundleFlagName = "bundle"
		} else {
			bundleFlagName = fmt.Sprintf("%v.bundle", cmdPrefix)
		}

		bundleFlagValue, err := cmd.Flags().GetString(bundleFlagName)
		if err != nil {
			return err, false
		}
		m.Bundle = &bundleFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePolyscopeBundleNameFlags(depth int, m *models.PolyscopeBundle, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = &nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}