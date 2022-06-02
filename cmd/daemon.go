// Copyright 2022 Ralf Geschke. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/geschke/golrackpi"
	"github.com/geschke/invafetch/pkg"
	"github.com/spf13/cobra"
)

func init() {

	rootCmd.AddCommand(startCmd)

}

var startCmd = &cobra.Command{
	Use: "start",

	Short: "Start collecting and storing values from inverter",
	//Long:  ``,

	Run: func(cmd *cobra.Command,
		args []string) {
		startCollect()
	},
}

// readProcessdataConfig loads the processdata definition from the file "processdata.json". This structure is used when requesting processdata
// from the inverter, so the requested data will be stored into the database.
func readProcessdataConfig() ([]golrackpi.ProcessData, error) {
	filename := "processdata.json" // filename is currently fixed, the processdata definition comes with invafetch package and contains nearly all processdata parameters

	var processData []golrackpi.ProcessData

	f, err := os.ReadFile(filename)
	if err != nil {
		return processData, err
	}
	if !json.Valid(f) {
		return processData, errors.New("file content has invalid JSON structure")
	}

	err = json.Unmarshal(f, &processData)
	if err != nil {
		return processData, err
	}

	return processData, nil
}

// startCollect starts the goroutines which are requesting processdata from the inverter
func startCollect() {

	collectData, err := readProcessdataConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, "An error occurred when reading processdata.json definition:", err)
		os.Exit(1)
	}

	authData := golrackpi.AuthClient{
		Scheme:   authData.Scheme,
		Server:   authData.Server,
		Password: authData.Password,
	}

	daemon := pkg.CollectDaemon{AuthData: authData, DbConfig: dbConfig}
	daemon.Start(collectData, timeNewLoginMinutes, timeRequestDurationSeconds)
}
