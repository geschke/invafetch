// Copyright 2022 Ralf Geschke. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"encoding/json"
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

func readProcessdataConfig() ([]golrackpi.ProcessData, error) {
	filename := "processdata.json"
	// todo: get filename from config variable, error handling and more
	f, err := os.ReadFile(filename)
	if err != nil {
		panic("file problem")
		// todo: error handling
	}
	fmt.Println(json.Valid(f))

	var processData []golrackpi.ProcessData
	err = json.Unmarshal(f, &processData)
	if err != nil {
		panic("json file error") // todo....

	}
	fmt.Println(processData)
	return processData, nil
}

func startCollect() {

	collectData, err := readProcessdataConfig()
	if err != nil {
		panic("problem!!!")
	}

	authData := golrackpi.AuthClient{
		Scheme:   authData.Scheme,
		Server:   authData.Server,
		Password: authData.Password,
	}

	daemon := pkg.CollectDaemon{AuthData: authData}
	daemon.Start(collectData)
}
