// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"swagger-example/client/container"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationContainerStopContainerByIDCmd returns a cmd to handle operation stopContainerById
func makeOperationContainerStopContainerByIDCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "stopContainerById",
		Short: `TODO`,
		RunE:  runOperationContainerStopContainerByID,
	}

	if err := registerOperationContainerStopContainerByIDParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationContainerStopContainerByID uses cmd flags to call endpoint api
func runOperationContainerStopContainerByID(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := container.NewStopContainerByIDParams()
	if err, _ := retrieveOperationContainerStopContainerByIDArtifactNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerStopContainerByIDUrcapIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationContainerStopContainerByIDVendorIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationContainerStopContainerByIDResult(appCli.Container.StopContainerByID(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationContainerStopContainerByIDParamFlags registers all flags needed to fill params
func registerOperationContainerStopContainerByIDParamFlags(cmd *cobra.Command) error {
	if err := registerOperationContainerStopContainerByIDArtifactNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerStopContainerByIDUrcapIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationContainerStopContainerByIDVendorIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationContainerStopContainerByIDArtifactNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationContainerStopContainerByIDUrcapIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationContainerStopContainerByIDVendorIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationContainerStopContainerByIDArtifactNameFlag(m *container.StopContainerByIDParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationContainerStopContainerByIDUrcapIDFlag(m *container.StopContainerByIDParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationContainerStopContainerByIDVendorIDFlag(m *container.StopContainerByIDParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationContainerStopContainerByIDResult parses request result and return the string content
func parseOperationContainerStopContainerByIDResult(resp0 *container.StopContainerByIDOK, respErr error) (string, error) {
	if respErr != nil {

		var iRespD interface{} = respErr
		respD, ok := iRespD.(*container.StopContainerByIDDefault)
		if ok {
			if !swag.IsZero(respD) && !swag.IsZero(respD.Payload) {
				msgStr, err := json.Marshal(respD.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning stopContainerByIdOK is not supported

		return "", respErr
	}

	// warning: non schema response stopContainerByIdOK is not supported by go-swagger cli yet.

	return "", nil
}
