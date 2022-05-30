// Copyright 2022 Ralf Geschke. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/geschke/golrackpi"
	"github.com/geschke/invafetch/pkg/dbconn"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile  string
	dbConfig dbconn.DatabaseConfiguration

	rootCmd = &cobra.Command{
		Use:   "invafetch",
		Short: "A tool for retrieving values from Kostal Plenticore inverters",
		//Long: ` `,

	}
)

var authData golrackpi.AuthClient

// init sets the global flags and their options.
func init() {
	cobra.OnInitialize(initConfig, initAuthData)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ~/.env)")

	rootCmd.PersistentFlags().StringVarP(&authData.Password, "password", "p", "", "Password (required)")
	viper.BindPFlag("inv_password", rootCmd.PersistentFlags().Lookup("password"))

	rootCmd.PersistentFlags().StringVarP(&authData.Server, "server", "s", "", "Server (e.g. inverter IP address) (required)")
	viper.BindPFlag("inv_server", rootCmd.PersistentFlags().Lookup("server"))

	rootCmd.PersistentFlags().StringVarP(&authData.Scheme, "scheme", "m", "", "Scheme (http or https, default http)")
	viper.BindPFlag("inv_scheme", rootCmd.PersistentFlags().Lookup("scheme"))

	//rootCmd.MarkPersistentFlagRequired("password")
	//rootCmd.MarkPersistentFlagRequired("server")

	rootCmd.PersistentFlags().StringVarP(&dbConfig.DBName, "dbname", "", "", "Database name")
	viper.BindPFlag("dbName", rootCmd.PersistentFlags().Lookup("dbname"))

	rootCmd.PersistentFlags().StringVarP(&dbConfig.DBHost, "dbhost", "", "", "Database host")
	viper.BindPFlag("dbHost", rootCmd.PersistentFlags().Lookup("dbhost"))

	rootCmd.PersistentFlags().StringVarP(&dbConfig.DBUser, "dbuser", "", "", "Database user")
	viper.BindPFlag("dbUser", rootCmd.PersistentFlags().Lookup("dbuser"))

	rootCmd.PersistentFlags().StringVarP(&dbConfig.DBPassword, "dbpassword", "", "", "Database password")
	viper.BindPFlag("dbPassword", rootCmd.PersistentFlags().Lookup("dbpassword"))

	rootCmd.PersistentFlags().StringVarP(&dbConfig.DBPort, "dbport", "", "", "Database port (default: 3306)")
	viper.BindPFlag("dbPort", rootCmd.PersistentFlags().Lookup("dbport"))

}

func initAuthData() {
	error := false
	if !viper.IsSet("INV_SERVER") {
		fmt.Fprintln(os.Stderr, errors.New("server parameter / INV_SERVER variable missing.\nPlease use --server options or add INV_SERVER to the config file or to ENV variables"))
		error = true
	}
	if !viper.IsSet("INV_PASSWORD") {
		fmt.Fprintln(os.Stderr, errors.New("password parameter / INV_PASSWORD variable missing.\nPlease use --password options or add INV_PASSWORD to the config file or to ENV variables"))
		error = true
	}
	if error {
		os.Exit(1)
	}

	authData.Server = viper.Get("INV_SERVER").(string)
	authData.Password = viper.Get("INV_PASSWORD").(string)
	authData.Scheme = viper.Get("INV_SCHEME").(string)

}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the --config flag.
		viper.SetConfigFile(cfgFile)
	} else {

		viper.AddConfigPath(".")
		viper.AddConfigPath("./config")
		viper.AddConfigPath("/config")

	}
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	//viper.AutomaticEnv()

	viper.SetDefault("dbPort", "3306")
	viper.SetDefault("inv_scheme", "http")

	viper.BindEnv("DBHOST")
	viper.BindEnv("DBNAME")
	viper.BindEnv("DBUSER")
	viper.BindEnv("DBPASSWORD")
	viper.BindEnv("DBPORT")
	viper.BindEnv("INV_SERVER")
	viper.BindEnv("INV_PASSWORD")
	viper.BindEnv("INV_SCHEME")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// fail silently

		} else {
			fmt.Fprintln(os.Stderr, errors.New("config file was found but another error occurred: "+err.Error()))
			os.Exit(1)
		}
	}

	err := viper.Unmarshal(&dbConfig)
	if err != nil {
		fmt.Fprintln(os.Stderr, errors.New("an error occurred: "+err.Error()))
		os.Exit(1)
	}

}

// Exec is the entrypoint of the Cobra CLI library.
func Exec() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
