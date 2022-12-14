// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"universal-robots.com/urservice/client/container"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationContainerGetContainerByIDCmd returns a cmd to handle operation getContainerById
func makeOperationContainerGetContainerByIDCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getContainerById",
		Short: `TODO`,
		RunE:  runOperationContainerGetContainerByID,
	}

	if err := registerOperationContainerGetContainerByIDParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationContainerGetContainerByID uses cmd flags to call endpoint api
func runOperationContainerGetContainerByID(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := container.NewGetContainerByIDParams()
	if err, _ := retrieveOperationContainerGetContainerByIDArtifactNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerGetContainerByIDUrcapIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerGetContainerByIDVendorIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationContainerGetContainerByIDResult(appCli.Container.GetContainerByID(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationContainerGetContainerByIDParamFlags registers all flags needed to fill params
func registerOperationContainerGetContainerByIDParamFlags(cmd *cobra.Command) error {
	if err := registerOperationContainerGetContainerByIDArtifactNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerGetContainerByIDUrcapIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerGetContainerByIDVendorIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationContainerGetContainerByIDArtifactNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	artifactNameDescription := `Required. artifact name of container defined in manifest`

	var artifactNameFlagName string
	if cmdPrefix == "" {
		artifactNameFlagName = "artifactName"
	} else {
		artifactNameFlagName = fmt.Sprintf("%v.artifactName", cmdPrefix)
	}

	var artifactNameFlagDefault string

	_ = cmd.PersistentFlags().String(artifactNameFlagName, artifactNameFlagDefault, artifactNameDescription)

	return nil
}
func registerOperationContainerGetContainerByIDUrcapIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationContainerGetContainerByIDVendorIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationContainerGetContainerByIDArtifactNameFlag(m *container.GetContainerByIDParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("artifactName") {

		var artifactNameFlagName string
		if cmdPrefix == "" {
			artifactNameFlagName = "artifactName"
		} else {
			artifactNameFlagName = fmt.Sprintf("%v.artifactName", cmdPrefix)
		}

		artifactNameFlagValue, err := cmd.Flags().GetString(artifactNameFlagName)
		if err != nil {
			return err, false
		}
		m.ArtifactName = artifactNameFlagValue

	}
	return nil, retAdded
}
func retrieveOperationContainerGetContainerByIDUrcapIDFlag(m *container.GetContainerByIDParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationContainerGetContainerByIDVendorIDFlag(m *container.GetContainerByIDParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationContainerGetContainerByIDResult parses request result and return the string content
func parseOperationContainerGetContainerByIDResult(resp0 *container.GetContainerByIDOK, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*container.GetContainerByIDDefault)
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
		resp0, ok := iResp0.(*container.GetContainerByIDOK)
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
