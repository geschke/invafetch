// Copyright 2022 Ralf Geschke. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"

	"github.com/geschke/golrackpi"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "invafetch",
	Short: "A tool for retrieving values from Kostal Plenticore inverters",
	//Long: ` `,
}

var authData golrackpi.AuthClient

var (
	outputCSV       bool   = false
	outputFile      string = ""
	outputTimestamp bool   = false
	outputAppend    bool   = false
	outputNoHeaders bool   = false
)

// init sets the global flags and their options.
func init() {
	rootCmd.PersistentFlags().StringVarP(&authData.Password, "password", "p", "", "Password (required)")
	rootCmd.PersistentFlags().StringVarP(&authData.Server, "server", "s", "", "Server (e.g. inverter IP address) (required)")
	rootCmd.PersistentFlags().StringVarP(&authData.Scheme, "scheme", "m", "", "Scheme (http or https, default http)")
	rootCmd.MarkPersistentFlagRequired("password")
	rootCmd.MarkPersistentFlagRequired("server")

}

// Exec is the entrypoint of the Cobra CLI library.
func Exec() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
