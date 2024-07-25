package utils

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DatabaseKey = "db"

func initDatabase(db *sql.DB) {
	transaction, err := db.Begin()
	if err != nil {
		log.Fatal("Unable to begin transation for database initialization", err)
		return
	}

	transaction.Exec(
		"CREATE TABLE IF NOT EXISTS `events` (" +
			"`id` int unsigned NOT NULL AUTO_INCREMENT," +
			"`name` varchar(256) NOT NULL," +
			"`number_of_registrations` int unsigned NOT NULL," +
			"PRIMARY KEY (`id`)" +
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;")

	err = transaction.Commit()
	if err != nil {
		log.Fatal("Error commiting transaction for database initialization", err)
		return
	}
}

func migrateDatabaseToV2(db *sql.DB) {
	transaction, err := db.Begin()
	if err != nil {
		log.Fatal("Unable to begin transation for database upgrade v2", err)
		return
	}

	transaction.Exec(
		"CREATE TABLE IF NOT EXISTS `users` (" +
			"`id` int(11) unsigned NOT NULL AUTO_INCREMENT," +
			"`created_by` int(11) unsigned DEFAULT NULL," +
			"`created_at` datetime DEFAULT NULL," +
			"`updated_by` int(11) unsigned DEFAULT NULL," +
			"`updated_at` datetime DEFAULT NULL," +
			"`deleted_by` int(11) unsigned DEFAULT NULL," +
			"`deleted_at` datetime DEFAULT NULL," +
			"`last_logged_in` datetime DEFAULT NULL," +
			"`email` varchar(254) NOT NULL," +
			"`first_name` varchar(64) NOT NULL," +
			"`last_name` varchar(64) NOT NULL," +
			"`username` varchar(64) NOT NULL," +
			"`password` varchar(64) NOT NULL," +
			"PRIMARY KEY (`id`)," +
			"UNIQUE KEY `username` (`username`)" +
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;")

	transaction.Exec(
		"CREATE TABLE IF NOT EXISTS `inventory` (" +
			"`id` int(11) unsigned NOT NULL AUTO_INCREMENT," +
			"`name` varchar(256) NOT NULL," +
			"`quantity` int(11) unsigned NOT NULL," +
			"PRIMARY KEY (`id`)" +
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;")

	transaction.Exec(
		"CREATE TABLE IF NOT EXISTS `forms` (" +
			"`id` int(11) unsigned NOT NULL AUTO_INCREMENT," +
			"`inventory_id` int(11) unsigned DEFAULT NULL," +
			"`created_by` int(11) unsigned NOT NULL," +
			"`created_at` datetime NOT NULL," +
			"`updated_by` int(11) unsigned DEFAULT NULL," +
			"`updated_at` datetime DEFAULT NULL," +
			"`deleted_by` int(11) unsigned DEFAULT NULL," +
			"`deleted_at` datetime DEFAULT NULL," +
			"`name` text NOT NULL," +
			"`publish_by` datetime DEFAULT NULL," +
			"PRIMARY KEY (`id`)" +
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;")

	transaction.Exec(
		"CREATE TABLE IF NOT EXISTS `form_fields` (" +
			"`id` int(11) unsigned NOT NULL AUTO_INCREMENT," +
			"`form_id` int(11) unsigned NOT NULL," +
			"`order` int(11) DEFAULT NULL," +
			"`required` tinyint(1) NOT NULL DEFAULT 0," +
			"`type` varchar(8) NOT NULL DEFAULT 'TEXT'," +
			"`metadata` text DEFAULT NULL," +
			"PRIMARY KEY (`id`)" +
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;")

	transaction.Exec(
		"CREATE TABLE IF NOT EXISTS `form_field_options` (" +
			"`id` int(11) unsigned NOT NULL AUTO_INCREMENT," +
			"`form_field_id` int(11) unsigned NOT NULL," +
			"`inventory_id` int(11) unsigned NOT NULL," +
			"`name` text NOT NULL," +
			"PRIMARY KEY (`id`)" +
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;")

	transaction.Exec(
		"CREATE TABLE IF NOT EXISTS `submissions` (" +
			"`id` int(11) unsigned NOT NULL AUTO_INCREMENT," +
			"`form_id` int(11) unsigned NOT NULL," +
			"`created_at` datetime NOT NULL," +
			"PRIMARY KEY (`id`)" +
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;")

	transaction.Exec(
		"CREATE TABLE IF NOT EXISTS `submission_fields` (" +
			"`id` int(11) unsigned NOT NULL AUTO_INCREMENT," +
			"`submission_id` int(11) unsigned NOT NULL," +
			"`form_field_id` int(11) unsigned NOT NULL," +
			"`value` text NOT NULL," +
			"PRIMARY KEY (`id`)" +
			") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;")

	err = transaction.Commit()
	if err != nil {
		log.Fatal("Error commiting transaction for database upgrade v2", err)
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

	initDatabase(db)
	migrateDatabaseToV2(db)

	return db
}
