// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/spf13/cobra"
	"universal-robots.com/urservice/models"
)

// Schema cli for Metadata

// register flags to command
func registerModelMetadataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerMetadataCopyright(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMetadataDescription(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMetadataLicense(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMetadataUrcapID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMetadataUrcapName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMetadataVendorID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMetadataVendorName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMetadataVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerMetadataCopyright(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	copyrightDescription := `Specifies copy rights for the urcap`

	var copyrightFlagName string
	if cmdPrefix == "" {
		copyrightFlagName = "copyright"
	} else {
		copyrightFlagName = fmt.Sprintf("%v.copyright", cmdPrefix)
	}

	var copyrightFlagDefault string

	_ = cmd.PersistentFlags().String(copyrightFlagName, copyrightFlagDefault, copyrightDescription)

	return nil
}

func registerMetadataDescription(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	descriptionDescription := `Short description of urcap`

	var descriptionFlagName string
	if cmdPrefix == "" {
		descriptionFlagName = "description"
	} else {
		descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
	}

	var descriptionFlagDefault string

	_ = cmd.PersistentFlags().String(descriptionFlagName, descriptionFlagDefault, descriptionDescription)

	return nil
}

func registerMetadataLicense(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	licenseDescription := `License of various software/hardware used while developing the urcap`

	var licenseFlagName string
	if cmdPrefix == "" {
		licenseFlagName = "license"
	} else {
		licenseFlagName = fmt.Sprintf("%v.license", cmdPrefix)
	}

	var licenseFlagDefault string

	_ = cmd.PersistentFlags().String(licenseFlagName, licenseFlagDefault, licenseDescription)

	return nil
}

func registerMetadataUrcapID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	urcapIdDescription := `Required. Urcap id for this application. The ID must be unique for each urcap application developed by a company (vendorID).
`

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

func registerMetadataUrcapName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	urcapNameDescription := `Required. Urcap name of this application. This Will be displayed on user interface.
`

	var urcapNameFlagName string
	if cmdPrefix == "" {
		urcapNameFlagName = "urcapName"
	} else {
		urcapNameFlagName = fmt.Sprintf("%v.urcapName", cmdPrefix)
	}

	var urcapNameFlagDefault string

	_ = cmd.PersistentFlags().String(urcapNameFlagName, urcapNameFlagDefault, urcapNameDescription)

	return nil
}

func registerMetadataVendorID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	vendorIdDescription := `Required. Company urcap developer ID. The ID is shared among all urcap applications developed by company.
`

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

func registerMetadataVendorName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	vendorNameDescription := `Required. Urcap name of this application. This Will be displayed on user interface.
`

	var vendorNameFlagName string
	if cmdPrefix == "" {
		vendorNameFlagName = "vendorName"
	} else {
		vendorNameFlagName = fmt.Sprintf("%v.vendorName", cmdPrefix)
	}

	var vendorNameFlagDefault string

	_ = cmd.PersistentFlags().String(vendorNameFlagName, vendorNameFlagDefault, vendorNameDescription)

	return nil
}

func registerMetadataVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	versionDescription := `Required. Urcap version (major.minor.patch)
`

	var versionFlagName string
	if cmdPrefix == "" {
		versionFlagName = "version"
	} else {
		versionFlagName = fmt.Sprintf("%v.version", cmdPrefix)
	}

	var versionFlagDefault string

	_ = cmd.PersistentFlags().String(versionFlagName, versionFlagDefault, versionDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelMetadataFlags(depth int, m *models.Metadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, copyrightAdded := retrieveMetadataCopyrightFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || copyrightAdded

	err, descriptionAdded := retrieveMetadataDescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, licenseAdded := retrieveMetadataLicenseFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || licenseAdded

	err, urcapIdAdded := retrieveMetadataUrcapIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || urcapIdAdded

	err, urcapNameAdded := retrieveMetadataUrcapNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || urcapNameAdded

	err, vendorIdAdded := retrieveMetadataVendorIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || vendorIdAdded

	err, vendorNameAdded := retrieveMetadataVendorNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || vendorNameAdded

	err, versionAdded := retrieveMetadataVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || versionAdded

	return nil, retAdded
}

func retrieveMetadataCopyrightFlags(depth int, m *models.Metadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	copyrightFlagName := fmt.Sprintf("%v.copyright", cmdPrefix)
	if cmd.Flags().Changed(copyrightFlagName) {

		var copyrightFlagName string
		if cmdPrefix == "" {
			copyrightFlagName = "copyright"
		} else {
			copyrightFlagName = fmt.Sprintf("%v.copyright", cmdPrefix)
		}

		copyrightFlagValue, err := cmd.Flags().GetString(copyrightFlagName)
		if err != nil {
			return err, false
		}
		m.Copyright = copyrightFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMetadataDescriptionFlags(depth int, m *models.Metadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	descriptionFlagName := fmt.Sprintf("%v.description", cmdPrefix)
	if cmd.Flags().Changed(descriptionFlagName) {

		var descriptionFlagName string
		if cmdPrefix == "" {
			descriptionFlagName = "description"
		} else {
			descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
		}

		descriptionFlagValue, err := cmd.Flags().GetString(descriptionFlagName)
		if err != nil {
			return err, false
		}
		m.Description = descriptionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMetadataLicenseFlags(depth int, m *models.Metadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	licenseFlagName := fmt.Sprintf("%v.license", cmdPrefix)
	if cmd.Flags().Changed(licenseFlagName) {

		var licenseFlagName string
		if cmdPrefix == "" {
			licenseFlagName = "license"
		} else {
			licenseFlagName = fmt.Sprintf("%v.license", cmdPrefix)
		}

		licenseFlagValue, err := cmd.Flags().GetString(licenseFlagName)
		if err != nil {
			return err, false
		}
		m.License = licenseFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMetadataUrcapIDFlags(depth int, m *models.Metadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	urcapIdFlagName := fmt.Sprintf("%v.urcapID", cmdPrefix)
	if cmd.Flags().Changed(urcapIdFlagName) {

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
		m.UrcapID = &urcapIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMetadataUrcapNameFlags(depth int, m *models.Metadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	urcapNameFlagName := fmt.Sprintf("%v.urcapName", cmdPrefix)
	if cmd.Flags().Changed(urcapNameFlagName) {

		var urcapNameFlagName string
		if cmdPrefix == "" {
			urcapNameFlagName = "urcapName"
		} else {
			urcapNameFlagName = fmt.Sprintf("%v.urcapName", cmdPrefix)
		}

		urcapNameFlagValue, err := cmd.Flags().GetString(urcapNameFlagName)
		if err != nil {
			return err, false
		}
		m.UrcapName = &urcapNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMetadataVendorIDFlags(depth int, m *models.Metadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	vendorIdFlagName := fmt.Sprintf("%v.vendorID", cmdPrefix)
	if cmd.Flags().Changed(vendorIdFlagName) {

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
		m.VendorID = &vendorIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMetadataVendorNameFlags(depth int, m *models.Metadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	vendorNameFlagName := fmt.Sprintf("%v.vendorName", cmdPrefix)
	if cmd.Flags().Changed(vendorNameFlagName) {

		var vendorNameFlagName string
		if cmdPrefix == "" {
			vendorNameFlagName = "vendorName"
		} else {
			vendorNameFlagName = fmt.Sprintf("%v.vendorName", cmdPrefix)
		}

		vendorNameFlagValue, err := cmd.Flags().GetString(vendorNameFlagName)
		if err != nil {
			return err, false
		}
		m.VendorName = &vendorNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMetadataVersionFlags(depth int, m *models.Metadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	versionFlagName := fmt.Sprintf("%v.version", cmdPrefix)
	if cmd.Flags().Changed(versionFlagName) {

		var versionFlagName string
		if cmdPrefix == "" {
			versionFlagName = "version"
		} else {
			versionFlagName = fmt.Sprintf("%v.version", cmdPrefix)
		}

		versionFlagValue, err := cmd.Flags().GetString(versionFlagName)
		if err != nil {
			return err, false
		}
		m.Version = &versionFlagValue

		retAdded = true
	}

	return nil, retAdded
}