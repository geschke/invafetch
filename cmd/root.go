// Copyright 2022 Ralf Geschke. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"

	"github.com/geschke/golrackpi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile    string
	dbHost     string
	dbName     string
	dbUser     string
	dbPassword string
	dbPort     string

	rootCmd = &cobra.Command{
		Use:   "invafetch",
		Short: "A tool for retrieving values from Kostal Plenticore inverters",
		//Long: ` `,
	}
)

var authData golrackpi.AuthClient

// init sets the global flags and their options.
func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ~/.env)")

	// todo: make authData variables usable by environment
	rootCmd.PersistentFlags().StringVarP(&authData.Password, "password", "p", "", "Password (required)")
	rootCmd.PersistentFlags().StringVarP(&authData.Server, "server", "s", "", "Server (e.g. inverter IP address) (required)")
	rootCmd.PersistentFlags().StringVarP(&authData.Scheme, "scheme", "m", "", "Scheme (http or https, default http)")
	rootCmd.MarkPersistentFlagRequired("password")
	rootCmd.MarkPersistentFlagRequired("server")

	rootCmd.PersistentFlags().StringVarP(&dbName, "dbname", "", "", "Database name")
	viper.BindPFlag("dbName", rootCmd.PersistentFlags().Lookup("dbname"))

	rootCmd.PersistentFlags().StringVarP(&dbHost, "dbhost", "", "", "Database host")
	viper.BindPFlag("dbHost", rootCmd.PersistentFlags().Lookup("dbhost"))

	rootCmd.PersistentFlags().StringVarP(&dbUser, "dbuser", "", "", "Database user")
	viper.BindPFlag("dbUser", rootCmd.PersistentFlags().Lookup("dbuser"))

	rootCmd.PersistentFlags().StringVarP(&dbPassword, "dbpassword", "", "", "Database password")
	viper.BindPFlag("dbPassword", rootCmd.PersistentFlags().Lookup("dbpassword"))

	rootCmd.PersistentFlags().StringVarP(&dbPort, "dbport", "", "", "Database port (default: 3306)")
	viper.BindPFlag("dbPort", rootCmd.PersistentFlags().Lookup("dbport"))

	viper.SetDefault("dbPort", "3306")

}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {

		// Find home directory.
		//home, err := os.UserHomeDir()
		//cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(".")
		viper.AddConfigPath("./config")
		viper.AddConfigPath("/config")

	}
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())

	}

}

// Exec is the entrypoint of the Cobra CLI library.
func Exec() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
