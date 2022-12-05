// https://github.com/spf13/cobra/blob/main/user_guide.md
package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	cmdDeploy = &cobra.Command{
		Use:   "Deploy [robot ip] [urcapx file]",
		Short: "Deploy urcap to the robot",
		Long: `install urcap(urcapx) file on the robot or ursim.
Requires ip address and path to urcapx file.`,
		Args: cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}
)
