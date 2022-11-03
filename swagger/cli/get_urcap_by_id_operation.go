// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"swagger-example/client/urcap"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationUrcapGetUrcapByIDCmd returns a cmd to handle operation getUrcapById
func makeOperationUrcapGetUrcapByIDCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getUrcapById",
		Short: `TODO`,
		RunE:  runOperationUrcapGetUrcapByID,
	}

	if err := registerOperationUrcapGetUrcapByIDParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationUrcapGetUrcapByID uses cmd flags to call endpoint api
func runOperationUrcapGetUrcapByID(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := urcap.NewGetUrcapByIDParams()
	if err, _ := retrieveOperationUrcapGetUrcapByIDUrcapIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationUrcapGetUrcapByIDVendorIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationUrcapGetUrcapByIDResult(appCli.Urcap.GetUrcapByID(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationUrcapGetUrcapByIDParamFlags registers all flags needed to fill params
func registerOperationUrcapGetUrcapByIDParamFlags(cmd *cobra.Command) error {
	if err := registerOperationUrcapGetUrcapByIDUrcapIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationUrcapGetUrcapByIDVendorIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationUrcapGetUrcapByIDUrcapIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	urcapIdDescription := `Required. urcapID of urcap defined in manifest`

	var urcapIdFlagName string
	if cmdPrefix == "" {
		urcapIdFlagName = "urcapID"
	} else {
		urcapIdFlagName = fmt.Sprintf("%v.urcapID", cmdPrefix)
	}

	var urcapIdFlagDefault string

	_ = cmd.PersistentFlags().String(urcapIdFlagName, urcapIdFlagDefault, urcapIdDescription)

	return nil
}
func registerOperationUrcapGetUrcapByIDVendorIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	vendorIdDescription := `Required. vendorID of urcap defined in manifest`

	var vendorIdFlagName string
	if cmdPrefix == "" {
		vendorIdFlagName = "vendorID"
	} else {
		vendorIdFlagName = fmt.Sprintf("%v.vendorID", cmdPrefix)
	}

	var vendorIdFlagDefault string

	_ = cmd.PersistentFlags().String(vendorIdFlagName, vendorIdFlagDefault, vendorIdDescription)

	return nil
}

func retrieveOperationUrcapGetUrcapByIDUrcapIDFlag(m *urcap.GetUrcapByIDParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("urcapID") {

		var urcapIdFlagName string
		if cmdPrefix == "" {
			urcapIdFlagName = "urcapID"
		} else {
			urcapIdFlagName = fmt.Sprintf("%v.urcapID", cmdPrefix)
		}

		urcapIdFlagValue, err := cmd.Flags().GetString(urcapIdFlagName)
		if err != nil {
			return err, false
		}
		m.UrcapID = urcapIdFlagValue

	}
	return nil, retAdded
}
func retrieveOperationUrcapGetUrcapByIDVendorIDFlag(m *urcap.GetUrcapByIDParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("vendorID") {

		var vendorIdFlagName string
		if cmdPrefix == "" {
			vendorIdFlagName = "vendorID"
		} else {
			vendorIdFlagName = fmt.Sprintf("%v.vendorID", cmdPrefix)
		}

		vendorIdFlagValue, err := cmd.Flags().GetString(vendorIdFlagName)
		if err != nil {
			return err, false
		}
		m.VendorID = vendorIdFlagValue

	}
	return nil, retAdded
}

// parseOperationUrcapGetUrcapByIDResult parses request result and return the string content
func parseOperationUrcapGetUrcapByIDResult(resp0 *urcap.GetUrcapByIDOK, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*urcap.GetUrcapByIDDefault)
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
		resp0, ok := iResp0.(*urcap.GetUrcapByIDOK)
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
