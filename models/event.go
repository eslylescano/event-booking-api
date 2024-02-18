package models

import (
	"time"

	"example.com/mod/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserdId     int
}

var events []Event = []Event{}

func (e *Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)`
	smt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	result, err := smt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserdId)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err

}

func GetAllEvents() []Event {
	return events
}
