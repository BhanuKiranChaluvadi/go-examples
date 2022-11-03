// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"universal-robots.com/urservice/models"
)

// Schema cli for ErrorTarget

// register flags to command
func registerModelErrorTargetFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerErrorTargetName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerErrorTargetType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerErrorTargetName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `This field MUST contain the name of the problematic field (with dot-syntax if necessary), query parameter, or header.`

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

func registerErrorTargetType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["field","parameter","header"]. This field MUST contain field, parameter, or header`

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	if err := cmd.RegisterFlagCompletionFunc(typeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["field","parameter","header"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelErrorTargetFlags(depth int, m *models.ErrorTarget, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, nameAdded := retrieveErrorTargetNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, typeAdded := retrieveErrorTargetTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveErrorTargetNameFlags(depth int, m *models.ErrorTarget, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Name = nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveErrorTargetTypeFlags(depth int, m *models.ErrorTarget, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	typeFlagName := fmt.Sprintf("%v.type", cmdPrefix)
	if cmd.Flags().Changed(typeFlagName) {

		var typeFlagName string
		if cmdPrefix == "" {
			typeFlagName = "type"
		} else {
			typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
		}

		typeFlagValue, err := cmd.Flags().GetString(typeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = typeFlagValue

		retAdded = true
	}

	return nil, retAdded
}