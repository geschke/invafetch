package dbconn

import (
	"database/sql"
	"fmt"
	"os"
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
func ConnectDB(dbConfig DatabaseConfiguration) (db *sql.DB) {
	var dsn string
	//if len(dsn) < 1 {
	dbname := dbConfig.DBName         // os.Getenv("DBNAME")
	dbhost := dbConfig.DBHost         //os.Getenv("DBHOST")
	dbuser := dbConfig.DBUser         // os.Getenv("DBUSER")
	dbpassword := dbConfig.DBPassword //os.Getenv("DBPASSWORD")

	dbport := dbConfig.DBPort // os.Getenv("DBPORT")
	if len(dbport) < 1 {
		dbport = "3306"
	}
	if len(dbname) >= 1 && len(dbhost) >= 1 && len(dbuser) >= 1 && len(dbpassword) >= 1 {
		dsn = dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname

	} else {
		fmt.Println("No database connect parameter found, exiting. Please use --dsn or environment variables to define database connection.")
		os.Exit(1)
	}

	//}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error by connecting database.")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return db
}
