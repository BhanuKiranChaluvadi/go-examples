package main

import (
	"log"
	"os"
)

const (
	INSTALLATION_PATH_ENV = "installationPath"
)

var INSTALLATION_PATH string

func init() {
	INSTALLATION_PATH, ok := os.LookupEnv(INSTALLATION_PATH_ENV)
	if !ok {
		log.Printf("%s environment variable not set\n", INSTALLATION_PATH_ENV)
	}
	log.Printf("%s=%s\n", INSTALLATION_PATH_ENV, INSTALLATION_PATH)

}
