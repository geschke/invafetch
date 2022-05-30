// Copyright 2022 Ralf Geschke. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/geschke/golrackpi"

	"github.com/spf13/cobra"
)

func init() {

	rootCmd.AddCommand(infoCmd)
	infoCmd.AddCommand(infoVersionCmd)

}

var infoCmd = &cobra.Command{
	Use: "info",

	Short: "Returns miscellaneous information",
	//Long:  ``,
	Run: func(cmd *cobra.Command,
		args []string) {
		handleInfo()
	},
}

var infoVersionCmd = &cobra.Command{
	Use: "version",

	Short: "Returns information about the inverter API",
	//Long:  ``,

	Run: func(cmd *cobra.Command,
		args []string) {
		infoVersion()
	},
}

// infoVersion prints information about the API (i.e. hostname, api version...)
func infoVersion() {
	lib := golrackpi.NewWithParameter(golrackpi.AuthClient{
		Scheme:   authData.Scheme,
		Server:   authData.Server,
		Password: authData.Password,
	})

	_, err := lib.Login()

	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}
	defer lib.Logout()

	info, err := lib.Version()
	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}

	for k, v := range info {
		fmt.Printf("%s: %v\n", k, v)
	}

}

// Handle info-related commands
func handleInfo() {
	fmt.Println("\nUnknown or missing command.\nRun invafetch info --help to show available commands.")
}
