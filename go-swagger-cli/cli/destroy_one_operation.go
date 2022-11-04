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

// makeOperationTodosDestroyOneCmd returns a cmd to handle operation destroyOne
func makeOperationTodosDestroyOneCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "destroyOne",
		Short: ``,
		RunE:  runOperationTodosDestroyOne,
	}

	if err := registerOperationTodosDestroyOneParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationTodosDestroyOne uses cmd flags to call endpoint api
func runOperationTodosDestroyOne(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := todos.NewDestroyOneParams()
	if err, _ := retrieveOperationTodosDestroyOneIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationTodosDestroyOneResult(appCli.Todos.DestroyOne(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationTodosDestroyOneParamFlags registers all flags needed to fill params
func registerOperationTodosDestroyOneParamFlags(cmd *cobra.Command) error {
	if err := registerOperationTodosDestroyOneIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationTodosDestroyOneIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. `

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault int64

	_ = cmd.PersistentFlags().Int64(idFlagName, idFlagDefault, idDescription)

	return nil
}

func retrieveOperationTodosDestroyOneIDFlag(m *todos.DestroyOneParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("id") {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetInt64(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

	}
	return nil, retAdded
}

// parseOperationTodosDestroyOneResult parses request result and return the string content
func parseOperationTodosDestroyOneResult(resp0 *todos.DestroyOneNoContent, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*todos.DestroyOneDefault)
		if ok {
			if !swag.IsZero(respD) && !swag.IsZero(respD.Payload) {
				msgStr, err := json.Marshal(respD.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning destroyOneNoContent is not supported

		return "", respErr
	}

	// warning: non schema response destroyOneNoContent is not supported by go-swagger cli yet.

	return "", nil
}
