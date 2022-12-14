// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
	"universal-robots/urcapctl/models"
)

// Schema cli for Error

// register flags to command
func registerModelErrorFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerErrorCode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerErrorMessage(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerErrorMoreInfo(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerErrorTarget(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerErrorCode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	codeDescription := `Error codes.This field contains a string succinctly identifying     the problem.`

	var codeFlagName string
	if cmdPrefix == "" {
		codeFlagName = "code"
	} else {
		codeFlagName = fmt.Sprintf("%v.code", cmdPrefix)
	}

	var codeFlagDefault string

	_ = cmd.PersistentFlags().String(codeFlagName, codeFlagDefault, codeDescription)

	return nil
}

func registerErrorMessage(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: message error unkown type is not supported by go-swagger cli yet

	return nil
}

func registerErrorMoreInfo(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	moreInfoDescription := `This field SHOULD contain a publicly-accessible URL where information about the error can be read in a web browser.`

	var moreInfoFlagName string
	if cmdPrefix == "" {
		moreInfoFlagName = "more_info"
	} else {
		moreInfoFlagName = fmt.Sprintf("%v.more_info", cmdPrefix)
	}

	var moreInfoFlagDefault string

	_ = cmd.PersistentFlags().String(moreInfoFlagName, moreInfoFlagDefault, moreInfoDescription)

	return nil
}

func registerErrorTarget(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var targetFlagName string
	if cmdPrefix == "" {
		targetFlagName = "target"
	} else {
		targetFlagName = fmt.Sprintf("%v.target", cmdPrefix)
	}

	if err := registerModelErrorTargetFlags(depth+1, targetFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelErrorFlags(depth int, m *models.Error, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, codeAdded := retrieveErrorCodeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || codeAdded

	err, messageAdded := retrieveErrorMessageFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || messageAdded

	err, moreInfoAdded := retrieveErrorMoreInfoFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || moreInfoAdded

	err, targetAdded := retrieveErrorTargetFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || targetAdded

	return nil, retAdded
}

func retrieveErrorCodeFlags(depth int, m *models.Error, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	codeFlagName := fmt.Sprintf("%v.code", cmdPrefix)
	if cmd.Flags().Changed(codeFlagName) {

		var codeFlagName string
		if cmdPrefix == "" {
			codeFlagName = "code"
		} else {
			codeFlagName = fmt.Sprintf("%v.code", cmdPrefix)
		}

		codeFlagValue, err := cmd.Flags().GetString(codeFlagName)
		if err != nil {
			return err, false
		}
		m.Code = codeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveErrorMessageFlags(depth int, m *models.Error, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	messageFlagName := fmt.Sprintf("%v.message", cmdPrefix)
	if cmd.Flags().Changed(messageFlagName) {
		// warning: message error unkown type is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveErrorMoreInfoFlags(depth int, m *models.Error, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	moreInfoFlagName := fmt.Sprintf("%v.more_info", cmdPrefix)
	if cmd.Flags().Changed(moreInfoFlagName) {

		var moreInfoFlagName string
		if cmdPrefix == "" {
			moreInfoFlagName = "more_info"
		} else {
			moreInfoFlagName = fmt.Sprintf("%v.more_info", cmdPrefix)
		}

		moreInfoFlagValue, err := cmd.Flags().GetString(moreInfoFlagName)
		if err != nil {
			return err, false
		}
		m.MoreInfo = moreInfoFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveErrorTargetFlags(depth int, m *models.Error, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	targetFlagName := fmt.Sprintf("%v.target", cmdPrefix)
	if cmd.Flags().Changed(targetFlagName) {
		// info: complex object target ErrorTarget is retrieved outside this Changed() block
	}
	targetFlagValue := m.Target
	if swag.IsZero(targetFlagValue) {
		targetFlagValue = &models.ErrorTarget{}
	}

	err, targetAdded := retrieveModelErrorTargetFlags(depth+1, targetFlagValue, targetFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || targetAdded
	if targetAdded {
		m.Target = targetFlagValue
	}

	return nil, retAdded
}
