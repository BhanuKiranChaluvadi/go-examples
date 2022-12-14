// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"universal-robots.com/urservice/client/urcap"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationUrcapGetUrcapsCmd returns a cmd to handle operation getUrcaps
func makeOperationUrcapGetUrcapsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "getUrcaps",
		Short: `This operation provides a view of installed urcaps's data in JSON.
`,
		RunE: runOperationUrcapGetUrcaps,
	}

	if err := registerOperationUrcapGetUrcapsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationUrcapGetUrcaps uses cmd flags to call endpoint api
func runOperationUrcapGetUrcaps(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := urcap.NewGetUrcapsParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationUrcapGetUrcapsResult(appCli.Urcap.GetUrcaps(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationUrcapGetUrcapsParamFlags registers all flags needed to fill params
func registerOperationUrcapGetUrcapsParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationUrcapGetUrcapsResult parses request result and return the string content
func parseOperationUrcapGetUrcapsResult(resp0 *urcap.GetUrcapsOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*urcap.GetUrcapsOK)
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
