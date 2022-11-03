// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"swagger-example/models"

	"github.com/spf13/cobra"
)

// Schema cli for ErrorCode

// register flags to command
func registerModelErrorCodeFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerErrorCodeMajor(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerErrorCodeMinor(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerErrorCodeMajor(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	majorDescription := `API context`

	var majorFlagName string
	if cmdPrefix == "" {
		majorFlagName = "major"
	} else {
		majorFlagName = fmt.Sprintf("%v.major", cmdPrefix)
	}

	var majorFlagDefault int64

	_ = cmd.PersistentFlags().Int64(majorFlagName, majorFlagDefault, majorDescription)

	return nil
}

func registerErrorCodeMinor(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	minorDescription := `error context`

	var minorFlagName string
	if cmdPrefix == "" {
		minorFlagName = "minor"
	} else {
		minorFlagName = fmt.Sprintf("%v.minor", cmdPrefix)
	}

	var minorFlagDefault int64

	_ = cmd.PersistentFlags().Int64(minorFlagName, minorFlagDefault, minorDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelErrorCodeFlags(depth int, m *models.ErrorCode, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, majorAdded := retrieveErrorCodeMajorFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || majorAdded

	err, minorAdded := retrieveErrorCodeMinorFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || minorAdded

	return nil, retAdded
}

func retrieveErrorCodeMajorFlags(depth int, m *models.ErrorCode, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	majorFlagName := fmt.Sprintf("%v.major", cmdPrefix)
	if cmd.Flags().Changed(majorFlagName) {

		var majorFlagName string
		if cmdPrefix == "" {
			majorFlagName = "major"
		} else {
			majorFlagName = fmt.Sprintf("%v.major", cmdPrefix)
		}

		majorFlagValue, err := cmd.Flags().GetInt64(majorFlagName)
		if err != nil {
			return err, false
		}
		m.Major = majorFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveErrorCodeMinorFlags(depth int, m *models.ErrorCode, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	minorFlagName := fmt.Sprintf("%v.minor", cmdPrefix)
	if cmd.Flags().Changed(minorFlagName) {

		var minorFlagName string
		if cmdPrefix == "" {
			minorFlagName = "minor"
		} else {
			minorFlagName = fmt.Sprintf("%v.minor", cmdPrefix)
		}

		minorFlagValue, err := cmd.Flags().GetInt64(minorFlagName)
		if err != nil {
			return err, false
		}
		m.Minor = minorFlagValue

		retAdded = true
	}

	return nil, retAdded
}
