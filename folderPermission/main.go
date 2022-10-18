package main

/*
 * https://chmod-calculator.com/
 * http://www.unixmantra.com/2013/04/unix-linux-file-permissions.html
 */
import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	// base_path string = "/root/.installed_urcaps"
	// base_path         string = "/home/bkch/.random"
	base_path         string = "/ur/bin/contributions"
	installation_path string = "universal-robots/gripper"
)

func main() {
	fmt.Println("Creating folder structure")
	// err := os.MkdirAll(filepath.Join(base_path, installation_path), 0766)
	err := os.MkdirAll(filepath.Join(base_path, installation_path), 0755)
	if err != nil {
		fmt.Println(err.Error())
	}
}
