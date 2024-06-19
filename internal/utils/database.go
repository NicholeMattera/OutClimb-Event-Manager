package utils

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DatabaseKey = "db"

func initAndMigrate(db *sql.DB) {
	// Database Initialization
	transaction, err := db.Begin()
	if err != nil {
		log.Fatal("Unable to begin transation for database initialization", err)
		return
	}

	transaction.Exec(`CREATE TABLE IF NOT EXISTS events (
		id int unsigned NOT NULL AUTO_INCREMENT,
		name varchar(256) NOT NULL,
		number_of_registrations int unsigned NOT NULL,
		PRIMARY KEY (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`)

	err = transaction.Commit()
	if err != nil {
		log.Fatal("Error commiting transaction for database initialization", err)
		return
	}
}

func Database() *sql.DB {
	username, usernameExists := os.LookupEnv("OUTCLIMB_DB_USER")
	if !usernameExists {
		log.Fatal("Error: No database username provided")
		return nil
	}

	password, passwordExists := os.LookupEnv("OUTCLIMB_DB_PASSWORD")
	if !passwordExists {
		log.Fatal("Error: No database password provided")
		return nil
	}

	host, hostExists := os.LookupEnv("OUTCLIMB_DB_HOST")
	if !hostExists {
		log.Fatal("Error: No database hostname provided")
		return nil
	}

	name, nameExists := os.LookupEnv("OUTCLIMB_DB_NAME")
	if !nameExists {
		log.Fatal("Error: No database name provided")
		return nil
	}

	db, err := sql.Open("mysql", username+":"+password+"@("+host+")/"+name+"?parseTime=true")
	if err != nil {
		log.Fatal("Error: Unable to connect to MySQL server", err)
		return nil
	}

	initAndMigrate(db)

	return db
}
