// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"go-swagger-cli/client/todos"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationTodosFindTodosCmd returns a cmd to handle operation findTodos
func makeOperationTodosFindTodosCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "findTodos",
		Short: ``,
		RunE:  runOperationTodosFindTodos,
	}

	if err := registerOperationTodosFindTodosParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationTodosFindTodos uses cmd flags to call endpoint api
func runOperationTodosFindTodos(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := todos.NewFindTodosParams()
	if err, _ := retrieveOperationTodosFindTodosLimitFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationTodosFindTodosSinceFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationTodosFindTodosResult(appCli.Todos.FindTodos(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationTodosFindTodosParamFlags registers all flags needed to fill params
func registerOperationTodosFindTodosParamFlags(cmd *cobra.Command) error {
	if err := registerOperationTodosFindTodosLimitParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationTodosFindTodosSinceParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationTodosFindTodosLimitParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	limitDescription := ``

	var limitFlagName string
	if cmdPrefix == "" {
		limitFlagName = "limit"
	} else {
		limitFlagName = fmt.Sprintf("%v.limit", cmdPrefix)
	}

	var limitFlagDefault int32 = 20

	_ = cmd.PersistentFlags().Int32(limitFlagName, limitFlagDefault, limitDescription)

	return nil
}
func registerOperationTodosFindTodosSinceParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	sinceDescription := ``

	var sinceFlagName string
	if cmdPrefix == "" {
		sinceFlagName = "since"
	} else {
		sinceFlagName = fmt.Sprintf("%v.since", cmdPrefix)
	}

	var sinceFlagDefault int64

	_ = cmd.PersistentFlags().Int64(sinceFlagName, sinceFlagDefault, sinceDescription)

	return nil
}

func retrieveOperationTodosFindTodosLimitFlag(m *todos.FindTodosParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("limit") {

		var limitFlagName string
		if cmdPrefix == "" {
			limitFlagName = "limit"
		} else {
			limitFlagName = fmt.Sprintf("%v.limit", cmdPrefix)
		}

		limitFlagValue, err := cmd.Flags().GetInt32(limitFlagName)
		if err != nil {
			return err, false
		}
		m.Limit = &limitFlagValue

	}
	return nil, retAdded
}
func retrieveOperationTodosFindTodosSinceFlag(m *todos.FindTodosParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("since") {

		var sinceFlagName string
		if cmdPrefix == "" {
			sinceFlagName = "since"
		} else {
			sinceFlagName = fmt.Sprintf("%v.since", cmdPrefix)
		}

		sinceFlagValue, err := cmd.Flags().GetInt64(sinceFlagName)
		if err != nil {
			return err, false
		}
		m.Since = &sinceFlagValue

	}
	return nil, retAdded
}

// parseOperationTodosFindTodosResult parses request result and return the string content
func parseOperationTodosFindTodosResult(resp0 *todos.FindTodosOK, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*todos.FindTodosDefault)
		if ok {
			if !swag.IsZero(respD) && !swag.IsZero(respD.Payload) {
				msgStr, err := json.Marshal(respD.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*todos.FindTodosOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
