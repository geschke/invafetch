// Copyright 2022 Ralf Geschke. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dbconn

import (
	"database/sql"
	"errors"
)

// DatabaseConfiguration holds connection parameters for a database
type DatabaseConfiguration struct {
	DBName     string `mapstructure:"DBNAME"`
	DBHost     string `mapstructure:"DBHOST"`
	DBUser     string `mapstructure:"DBUSER"`
	DBPassword string `mapstructure:"DBPASSWORD"`
	DBPort     string `mapstructure:"DBPORT"`
}

// ConnectDB returns a database connection
func ConnectDB(dbConfig DatabaseConfiguration) (*sql.DB, error) {
	var dsn string
	dbname := dbConfig.DBName
	dbhost := dbConfig.DBHost
	dbuser := dbConfig.DBUser
	dbpassword := dbConfig.DBPassword

	dbport := dbConfig.DBPort
	if len(dbport) < 1 {
		dbport = "3306"
	}

	if len(dbname) >= 1 && len(dbhost) >= 1 && len(dbuser) >= 1 && len(dbpassword) >= 1 {
		dsn = dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname

	} else {

		return nil, errors.New("no database connect parameter found, exiting. Please use parameter or environment variables to define database connection")

	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return db, errors.New("Error by connecting database:" + err.Error())

	}
	return db, nil
}

// CloseDB closes the database
func CloseDB(db *sql.DB) error {

	err := db.Close()
	if err != nil {
		return errors.New("Error by closing database: " + err.Error())
	}
	return nil
}
