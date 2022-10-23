package database

import (
	"database/sql"
	"errors"
	"log"
	"os"
)

func SetUpDatabase() {
	_, err := os.Stat("sqlite-database.db")
	if !errors.Is(err, os.ErrNotExist) {
		return
	}

	file, err := os.Create("sqlite-database.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	db, err := sql.Open("sqlite3", "../sqlite-database.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	query := `
	CREATE TABLE IF NOT EXISTS sequences(
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"quantity_valid_sequence" integer,
		"quantity_invalid_sequence" integer,
		"rate_valid_sequence" real
	);
	`

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err.Error())
	}
}
