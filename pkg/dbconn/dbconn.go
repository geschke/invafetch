// Copyright 2022 Ralf Geschke. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dbconn

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"
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
func ConnectDB(dbConfig DatabaseConfiguration, timeout time.Duration) (*sql.DB, error) {
	var dsn string
	//var ctx context.Context
	dbname := dbConfig.DBName
	dbhost := dbConfig.DBHost
	dbuser := dbConfig.DBUser
	dbpassword := dbConfig.DBPassword

	dbport := dbConfig.DBPort
	if len(dbport) < 1 {
		dbport = "3306"
	}

	if len(dbname) >= 1 && len(dbhost) >= 1 && len(dbuser) >= 1 && len(dbpassword) >= 1 {
		dsn = dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?timeout=10s"
	} else {

		return nil, errors.New("no database connect parameter found, exiting. Please use parameter or environment variables to define database connection")

	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return db, errors.New("error by opening database:" + err.Error())

	}

	if timeout == 0 {
		timeout = 10
	}
	// Open doesn't open a connection. Validate DSN data:
	// works, but requires a running database
	/*err = db.Ping()
	if err != nil {
		return db, errors.New("error by connecting database:" + err.Error())
	}*/

	// does not work: if the database isn't currently running, if returns an error without waiting to timeout, so it's the same as with db.Ping()
	/*	ctx, cancelfunc := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelfunc()
		err = db.PingContext(ctx)
		if err != nil {

			return db, errors.New("error by connecting database:" + err.Error())
		}
		log.Println("Connected to DB successfully")

	*/

	ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
	defer cancel()

	connectionEstablished := false

	for active := true; active; {
		select {
		case <-time.After(1 * time.Second):
			log.Println("try to connect database...")
			err = db.PingContext(ctx)
			if err == nil {
				active = false
				connectionEstablished = true
			}

		case <-ctx.Done():
			log.Println(ctx.Err()) // timeout exceeding, connecting was not successful
			active = false

		}
	}
	if !connectionEstablished {
		return db, errors.New("error by connecting database")
	}

	return db, nil
}

// CloseDB closes the database
func CloseDB(db *sql.DB) error {

	err := db.Close()
	if err != nil {
		return errors.New("error by closing database: " + err.Error())
	}
	return nil
}
