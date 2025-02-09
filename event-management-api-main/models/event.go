package models

import (
	"project/event-management-api/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
}

func (e *Event) Save() error {
	query := `INSERT INTO events (name, description, location, dateTime, userId) 
	VALUES (?, ?, ?, ?, ?)` // To prevent sql injection use ? instead of %s

	stmnt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmnt.Close()
	result, err := stmnt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)

	if err != nil {
		return err
	}

	id, _ := result.LastInsertId()
	e.ID = id

	return nil
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
		var event Event
		if err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id) //Fetch only one row

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) Update() error {
	query := `UPDATE events 
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?`

	stmnt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmnt.Close()

	_, err = stmnt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)

	if err != nil {
		return err
	}

	return nil
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"

	stmnt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmnt.Close()

	_, err = stmnt.Exec(event.ID)

	if err != nil {
		return err
	}

	return nil
}
