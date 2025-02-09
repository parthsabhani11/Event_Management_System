package models

import "project/event-management-api/db"

func (event Event) Register(userId int64) error {

	query := "INSERT INTO registrations (event_id, user_id) VALUES (?, ?)"

	stmnt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmnt.Close()

	_, err = stmnt.Exec(event.ID, userId)

	return err
}

func (event Event) Unregister(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"

	stmnt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmnt.Close()

	_, err = stmnt.Exec(event.ID, userId)

	return err
}
