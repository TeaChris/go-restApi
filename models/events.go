package models

import (
	"rest-api/db"
	"time"
)

type Event struct {
ID			int64	
Name		string 		`binding:"required"`
Description string 		`binding:"required"`
Location 	string		`binding:"required"`
DateTime 	time.Time	`binding:"required"`
UserId 		int			
}

var events = []Event{}

func (e Event) Save() error {
	// later: add it to a database
	query := `
	INSERT INTO events (name, description, location, date_time, user_id) 
	VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)

	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()

	e.ID = id

	return err

}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var e Event
		err :=rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId)

		if err != nil {
			return nil, err
		}

		events = append(events, e)
	}

	return events, nil
}