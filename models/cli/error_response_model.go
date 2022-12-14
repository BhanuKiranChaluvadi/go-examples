// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/spf13/cobra"
	"universal-robots.com/urservice/models"
)

// Schema cli for ErrorResponse

// register flags to command
func registerModelErrorResponseFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerErrorResponseErrors(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerErrorResponseStatusCode(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerErrorResponseTrace(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerErrorResponseErrors(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: errors []*Error array type is not supported by go-swagger cli yet

	return nil
}

func registerErrorResponseStatusCode(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: primitive status_code uint32 is not supported by go-swagger cli yet

	return nil
}

func registerErrorResponseTrace(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	traceDescription := `This field SHOULD contain a lowercase UUID uniquely identifying the request.`

	var traceFlagName string
	if cmdPrefix == "" {
		traceFlagName = "trace"
	} else {
		traceFlagName = fmt.Sprintf("%v.trace", cmdPrefix)
	}

	var traceFlagDefault string

	_ = cmd.PersistentFlags().String(traceFlagName, traceFlagDefault, traceDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelErrorResponseFlags(depth int, m *models.ErrorResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, errorsAdded := retrieveErrorResponseErrorsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || errorsAdded

	err, statusCodeAdded := retrieveErrorResponseStatusCodeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statusCodeAdded

	err, traceAdded := retrieveErrorResponseTraceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || traceAdded

	return nil, retAdded
}

func retrieveErrorResponseErrorsFlags(depth int, m *models.ErrorResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	errorsFlagName := fmt.Sprintf("%v.errors", cmdPrefix)
	if cmd.Flags().Changed(errorsFlagName) {
		// warning: errors array type []*Error is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveErrorResponseStatusCodeFlags(depth int, m *models.ErrorResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statusCodeFlagName := fmt.Sprintf("%v.status_code", cmdPrefix)
	if cmd.Flags().Changed(statusCodeFlagName) {

		// warning: primitive status_code uint32 is not supported by go-swagger cli yet

		retAdded = true
	}

	return nil, retAdded
}

func retrieveErrorResponseTraceFlags(depth int, m *models.ErrorResponse, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	traceFlagName := fmt.Sprintf("%v.trace", cmdPrefix)
	if cmd.Flags().Changed(traceFlagName) {

		var traceFlagName string
		if cmdPrefix == "" {
			traceFlagName = "trace"
		} else {
			traceFlagName = fmt.Sprintf("%v.trace", cmdPrefix)
		}

		traceFlagValue, err := cmd.Flags().GetString(traceFlagName)
		if err != nil {
			return err, false
		}
		m.Trace = traceFlagValue

		retAdded = true
	}

	return nil, retAdded
}
