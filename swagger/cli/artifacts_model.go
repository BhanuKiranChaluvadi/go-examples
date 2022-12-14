// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"urcapCli/models"

	"github.com/spf13/cobra"
)

// Schema cli for Artifacts

// register flags to command
func registerModelArtifactsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerArtifactsContainers(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerArtifactsPolyscopeBundles(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerArtifactsWebArchives(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerArtifactsContainers(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: containers []*Container array type is not supported by go-swagger cli yet

	return nil
}

func registerArtifactsPolyscopeBundles(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: polyscopeBundles []*PolyscopeBundle array type is not supported by go-swagger cli yet

	return nil
}

func registerArtifactsWebArchives(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: webArchives []*WebArchive array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelArtifactsFlags(depth int, m *models.Artifacts, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, containersAdded := retrieveArtifactsContainersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || containersAdded

	err, polyscopeBundlesAdded := retrieveArtifactsPolyscopeBundlesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || polyscopeBundlesAdded

	err, webArchivesAdded := retrieveArtifactsWebArchivesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || webArchivesAdded

	return nil, retAdded
}

func retrieveArtifactsContainersFlags(depth int, m *models.Artifacts, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	containersFlagName := fmt.Sprintf("%v.containers", cmdPrefix)
	if cmd.Flags().Changed(containersFlagName) {
		// warning: containers array type []*Container is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveArtifactsPolyscopeBundlesFlags(depth int, m *models.Artifacts, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	polyscopeBundlesFlagName := fmt.Sprintf("%v.polyscopeBundles", cmdPrefix)
	if cmd.Flags().Changed(polyscopeBundlesFlagName) {
		// warning: polyscopeBundles array type []*PolyscopeBundle is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveArtifactsWebArchivesFlags(depth int, m *models.Artifacts, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	webArchivesFlagName := fmt.Sprintf("%v.webArchives", cmdPrefix)
	if cmd.Flags().Changed(webArchivesFlagName) {
		// warning: webArchives array type []*WebArchive is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
