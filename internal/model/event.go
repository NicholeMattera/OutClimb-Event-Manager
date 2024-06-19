package model

import (
	"database/sql"
)

type Event struct {
	Id                    int
	Name                  string
	NumberOfRegistrations int
}

func GetEvent(db *sql.DB, name string) (event Event, err error) {
	query := `
		SELECT
			id,
			name,
			number_of_registrations
		FROM events
		WHERE name = ?
		LIMIT 1
	`
	row := db.QueryRow(query, name)
	err = row.Scan(&event.Id, &event.Name, &event.NumberOfRegistrations)
	if err != nil {
		return
	}

	return
}

func CreateEvent(db *sql.DB, name string) (event Event, err error) {
	query := `
		INSERT INTO events (name, number_of_registrations)
		VALUES (?, 1)
	`
	result, err := db.Exec(query, name)
	if err != nil {
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		return
	}

	event.Id = int(id)
	event.Name = name
	event.NumberOfRegistrations = 1

	return
}

func UpdateEvent(db *sql.DB, event *Event) (err error) {
	query := `
		UPDATE events
		SET number_of_registrations = ?
		WHERE id = ?
	`
	_, err = db.Exec(query, event.NumberOfRegistrations, event.Id)
	if err != nil {
		return
	}

	return
}
