// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/spf13/cobra"
	"universal-robots.com/urservice/models"
)

// Schema cli for WebArchive

// register flags to command
func registerModelWebArchiveFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerWebArchiveFolder(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebArchiveMasterIndex(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebArchiveName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerWebArchiveFolder(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	folderDescription := `Required. Path of the folder with all web related pages. This path is relative to the zipped urcap top folder. 
`

	var folderFlagName string
	if cmdPrefix == "" {
		folderFlagName = "folder"
	} else {
		folderFlagName = fmt.Sprintf("%v.folder", cmdPrefix)
	}

	var folderFlagDefault string

	_ = cmd.PersistentFlags().String(folderFlagName, folderFlagDefault, folderDescription)

	return nil
}

func registerWebArchiveMasterIndex(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	masterIndexDescription := `Path to the main index pages for the front to load.`

	var masterIndexFlagName string
	if cmdPrefix == "" {
		masterIndexFlagName = "masterIndex"
	} else {
		masterIndexFlagName = fmt.Sprintf("%v.masterIndex", cmdPrefix)
	}

	var masterIndexFlagDefault string

	_ = cmd.PersistentFlags().String(masterIndexFlagName, masterIndexFlagDefault, masterIndexDescription)

	return nil
}

func registerWebArchiveName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `Required. The name of the web archive
`

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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelWebArchiveFlags(depth int, m *models.WebArchive, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, folderAdded := retrieveWebArchiveFolderFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || folderAdded

	err, masterIndexAdded := retrieveWebArchiveMasterIndexFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || masterIndexAdded

	err, nameAdded := retrieveWebArchiveNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	return nil, retAdded
}

func retrieveWebArchiveFolderFlags(depth int, m *models.WebArchive, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	folderFlagName := fmt.Sprintf("%v.folder", cmdPrefix)
	if cmd.Flags().Changed(folderFlagName) {

		var folderFlagName string
		if cmdPrefix == "" {
			folderFlagName = "folder"
		} else {
			folderFlagName = fmt.Sprintf("%v.folder", cmdPrefix)
		}

		folderFlagValue, err := cmd.Flags().GetString(folderFlagName)
		if err != nil {
			return err, false
		}
		m.Folder = &folderFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebArchiveMasterIndexFlags(depth int, m *models.WebArchive, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	masterIndexFlagName := fmt.Sprintf("%v.masterIndex", cmdPrefix)
	if cmd.Flags().Changed(masterIndexFlagName) {

		var masterIndexFlagName string
		if cmdPrefix == "" {
			masterIndexFlagName = "masterIndex"
		} else {
			masterIndexFlagName = fmt.Sprintf("%v.masterIndex", cmdPrefix)
		}

		masterIndexFlagValue, err := cmd.Flags().GetString(masterIndexFlagName)
		if err != nil {
			return err, false
		}
		m.MasterIndex = masterIndexFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebArchiveNameFlags(depth int, m *models.WebArchive, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.Name = &nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}
