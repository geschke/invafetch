// Copyright 2022 Ralf Geschke. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/geschke/golrackpi"
	"github.com/geschke/invafetch/internal/dbconn"
	"github.com/geschke/invafetch/internal/invdb"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {

	rootCmd.AddCommand(infoCmd)
	infoCmd.AddCommand(infoVersionCmd)
	infoCmd.AddCommand(infoDbCmd)

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

	Short: "Returns information about the API",
	//Long:  ``,

	Run: func(cmd *cobra.Command,
		args []string) {
		infoVersion()
	},
}

var infoDbCmd = &cobra.Command{
	Use: "db",

	Short: "Returns information about the database (just testing)",
	//Long:  ``,

	Run: func(cmd *cobra.Command,
		args []string) {
		infoDb()
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

func GetDbConfig() dbconn.DatabaseConfiguration {
	// could something go wrong here?
	fmt.Println(viper.Get("dbName"))
	fmt.Println(viper.Get("dbHost"))
	fmt.Println(viper.Get("dbUser"))
	fmt.Println(viper.Get("dbPassword"))
	fmt.Println(viper.Get("dbPort"))
	var config dbconn.DatabaseConfiguration
	config.DBHost = viper.GetString("dbHost")
	config.DBName = viper.GetString("dbName")
	config.DBPassword = viper.GetString("dbPassword")
	config.DBUser = viper.GetString("dbUser")
	config.DBPort = viper.GetString("dbPort")
	return config
}

func infoDb() {
	fmt.Println("Test")
	config := dbconn.ConnectDB(GetDbConfig())

	repository := invdb.NewRepository(config)
	repository.GetProcessdata()
}

// Handle info-related commands
func handleInfo() {
	fmt.Println("\nUnknown or missing command.\nRun invafetch info --help to show available commands.")
}
