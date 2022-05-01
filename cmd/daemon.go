// Copyright 2022 Ralf Geschke. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/geschke/golrackpi"
	"github.com/geschke/invafetch/internal"

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

func startCollect() {
	/*	lib := golrackpi.NewWithParameter(golrackpi.AuthClient{
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
	*/

	authData := golrackpi.AuthClient{
		Scheme:   authData.Scheme,
		Server:   authData.Server,
		Password: authData.Password,
	}

	daemon := internal.CollectDaemon{AuthData: authData}
	daemon.Start()
}
