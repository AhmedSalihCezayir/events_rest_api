package models

import (
	"time"

	"example.com/events-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, user_id, date_time) 
	VALUES (?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query) // we could ofc use db.exec(query, params) but prepare is more efficient
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.UserID, e.DateTime)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query) // change -> exec, fetch -> query
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, date_time = ? 
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	if err != nil {
		return err
	}

	return nil
}

func (e Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID)
	return err
}

func (e Event) Register(userID int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userID)
	return err
}

func CancelRegistration(eventID, userID int64) error {
	query := "DELETE FROM registrations WHERE event_id = ?  AND user_id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil
	}
	defer stmt.Close()

	_, err = stmt.Exec(eventID, userID)
	return err
}

func GetEventAttendees(eventID int64) ([]User, error) {
	query := `
	SELECT users.id, users.email, users.password
	FROM users 
	INNER JOIN registrations 
	ON users.id = registrations.user_id
	WHERE registrations.event_id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(eventID)
	if err != nil {
		return nil, err
	}

	var attendees []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		attendees = append(attendees, user)
	}
	return attendees, nil
}