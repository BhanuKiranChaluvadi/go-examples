// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"swagger-example/client"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// debug flag indicating that cli should output debug logs
var debug bool

// config file location
var configFile string

// dry run flag
var dryRun bool

// name of the executable
var exeName string = filepath.Base(os.Args[0])

// logDebugf writes debug log to stdout
func logDebugf(format string, v ...interface{}) {
	if !debug {
		return
	}
	log.Printf(format, v...)
}

// depth of recursion to construct model flags
var maxDepth int = 5

// makeClient constructs a client object
func makeClient(cmd *cobra.Command, args []string) (*client.UrcapAPI, error) {
	hostname := viper.GetString("hostname")
	viper.SetDefault("base_path", client.DefaultBasePath)
	basePath := viper.GetString("base_path")
	scheme := viper.GetString("scheme")

	r := httptransport.New(hostname, basePath, []string{scheme})
	r.SetDebug(debug)
	// set custom producer and consumer to use the default ones

	r.Consumers["application/json"] = runtime.JSONConsumer()

	// warning: consumes multipart/form-data is not supported by go-swagger cli yet

	r.Producers["application/json"] = runtime.JSONProducer()

	appCli := client.New(r, strfmt.Default)
	logDebugf("Server url: %v://%v", scheme, hostname)
	return appCli, nil
}

// MakeRootCmd returns the root cmd
func MakeRootCmd() (*cobra.Command, error) {
	cobra.OnInitialize(initViperConfigs)

	// Use executable name as the command name
	rootCmd := &cobra.Command{
		Use: exeName,
	}

	// register basic flags
	rootCmd.PersistentFlags().String("hostname", client.DefaultHost, "hostname of the service")
	viper.BindPFlag("hostname", rootCmd.PersistentFlags().Lookup("hostname"))
	rootCmd.PersistentFlags().String("scheme", client.DefaultSchemes[0], fmt.Sprintf("Choose from: %v", client.DefaultSchemes))
	viper.BindPFlag("scheme", rootCmd.PersistentFlags().Lookup("scheme"))
	rootCmd.PersistentFlags().String("base-path", client.DefaultBasePath, fmt.Sprintf("For example: %v", client.DefaultBasePath))
	viper.BindPFlag("base_path", rootCmd.PersistentFlags().Lookup("base-path"))

	// configure debug flag
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "output debug logs")
	// configure config location
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file path")
	// configure dry run flag
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "do not send the request to server")

	// register security flags
	// add all operation groups
	operationGroupArtifactCmd, err := makeOperationGroupArtifactCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupArtifactCmd)

	operationGroupContainerCmd, err := makeOperationGroupContainerCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupContainerCmd)

	operationGroupUrcapCmd, err := makeOperationGroupUrcapCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupUrcapCmd)

	// add cobra completion
	rootCmd.AddCommand(makeGenCompletionCmd())

	return rootCmd, nil
}

// initViperConfigs initialize viper config using config file in '$HOME/.config/<cli name>/config.<json|yaml...>'
// currently hostname, scheme and auth tokens can be specified in this config file.
func initViperConfigs() {
	if configFile != "" {
		// use user specified config file location
		viper.SetConfigFile(configFile)
	} else {
		// look for default config
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(path.Join(home, ".config", exeName))
		viper.SetConfigName("config")
	}

	if err := viper.ReadInConfig(); err != nil {
		logDebugf("Error: loading config file: %v", err)
		return
	}
	logDebugf("Using config file: %v", viper.ConfigFileUsed())
}

func makeOperationGroupArtifactCmd() (*cobra.Command, error) {
	operationGroupArtifactCmd := &cobra.Command{
		Use:  "artifact",
		Long: ``,
	}

	operationGetArtifactActualNameByPathCmd, err := makeOperationArtifactGetArtifactActualNameByPathCmd()
	if err != nil {
		return nil, err
	}
	operationGroupArtifactCmd.AddCommand(operationGetArtifactActualNameByPathCmd)

	operationGetArtifactMetadataByPathCmd, err := makeOperationArtifactGetArtifactMetadataByPathCmd()
	if err != nil {
		return nil, err
	}
	operationGroupArtifactCmd.AddCommand(operationGetArtifactMetadataByPathCmd)

	return operationGroupArtifactCmd, nil
}
func makeOperationGroupContainerCmd() (*cobra.Command, error) {
	operationGroupContainerCmd := &cobra.Command{
		Use:  "container",
		Long: ``,
	}

	operationGetContainerByIDCmd, err := makeOperationContainerGetContainerByIDCmd()
	if err != nil {
		return nil, err
	}
	operationGroupContainerCmd.AddCommand(operationGetContainerByIDCmd)

	operationGetContainerPortMappingByIDCmd, err := makeOperationContainerGetContainerPortMappingByIDCmd()
	if err != nil {
		return nil, err
	}
	operationGroupContainerCmd.AddCommand(operationGetContainerPortMappingByIDCmd)

	operationStartContainerByIDCmd, err := makeOperationContainerStartContainerByIDCmd()
	if err != nil {
		return nil, err
	}
	operationGroupContainerCmd.AddCommand(operationStartContainerByIDCmd)

	operationStatusContainerByIDCmd, err := makeOperationContainerStatusContainerByIDCmd()
	if err != nil {
		return nil, err
	}
	operationGroupContainerCmd.AddCommand(operationStatusContainerByIDCmd)

	operationStopAllContainersCmd, err := makeOperationContainerStopAllContainersCmd()
	if err != nil {
		return nil, err
	}
	operationGroupContainerCmd.AddCommand(operationStopAllContainersCmd)

	operationStopContainerByIDCmd, err := makeOperationContainerStopContainerByIDCmd()
	if err != nil {
		return nil, err
	}
	operationGroupContainerCmd.AddCommand(operationStopContainerByIDCmd)

	operationValidateContainerByIDCmd, err := makeOperationContainerValidateContainerByIDCmd()
	if err != nil {
		return nil, err
	}
	operationGroupContainerCmd.AddCommand(operationValidateContainerByIDCmd)

	return operationGroupContainerCmd, nil
}
func makeOperationGroupUrcapCmd() (*cobra.Command, error) {
	operationGroupUrcapCmd := &cobra.Command{
		Use:  "urcap",
		Long: ``,
	}

	operationAddUrcapCmd, err := makeOperationUrcapAddUrcapCmd()
	if err != nil {
		return nil, err
	}
	operationGroupUrcapCmd.AddCommand(operationAddUrcapCmd)

	operationDeleteUrcapByIDCmd, err := makeOperationUrcapDeleteUrcapByIDCmd()
	if err != nil {
		return nil, err
	}
	operationGroupUrcapCmd.AddCommand(operationDeleteUrcapByIDCmd)

	operationGetUrcapByIDCmd, err := makeOperationUrcapGetUrcapByIDCmd()
	if err != nil {
		return nil, err
	}
	operationGroupUrcapCmd.AddCommand(operationGetUrcapByIDCmd)

	operationGetUrcapsCmd, err := makeOperationUrcapGetUrcapsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupUrcapCmd.AddCommand(operationGetUrcapsCmd)

	operationUpdateUrcapByIDCmd, err := makeOperationUrcapUpdateUrcapByIDCmd()
	if err != nil {
		return nil, err
	}
	operationGroupUrcapCmd.AddCommand(operationUpdateUrcapByIDCmd)

	return operationGroupUrcapCmd, nil
}
