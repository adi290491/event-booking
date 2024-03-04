package models

import (
	"event-booking/database"
	"log"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `"binding:"required"`
	Description string    `"binding:"required"`
	Location    string    `"binding:"required"`
	DateTime    time.Time `"binding:"required"`
	UserID      int64
}

var events = []Event{}

func (e *Event) Save() error {

	stmt, err := database.DB.Prepare(INSERT_EVENT)

	if err != nil {
		log.Fatal("Could not create insert statement:", err.Error())
	}

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	defer stmt.Close()

	if err != nil {
		log.Fatal("Error while inserting record:", err.Error())
	}

	return err
}

func GetAllEvents() ([]Event, error) {
	select_query := "SELECT * FROM events"
	rows, err := database.DB.Query(select_query)

	if err != nil {
		log.Fatal("Error while fetching records:", err.Error())
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

func GetEventById(event_id int64) (*Event, error) {

	row := database.DB.QueryRow(EVENT_BY_ID, event_id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return &event, err
	}

	return &event, nil
}

func (event Event) Update() error {

	stmt, err := database.DB.Prepare(UPDATE_EVENT)

	if err != nil {
		log.Fatal("Could not create update query:", err.Error())
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID, event.Name, event.Description, event.Location, event.DateTime)

	return err
}

func (event Event) Delete() error {
	stmt, err := database.DB.Prepare(DELETE_EVENT)

	if err != nil {
		log.Fatal("Could not create delete query:", err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID)

	return err
}

func (event Event) Register(userId int64) error {
	stmt, err := database.DB.Prepare(INSERT_INTO_REGISTRATION)

	if err != nil {
		log.Fatal("Could not create insert registration query:", err.Error())
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID, userId)

	return err
}

func (event Event) Cancel(userId int64) error {
	stmt, err := database.DB.Prepare(CANCEL_EVENT_REGISTRATION)

	if err != nil {
		log.Fatal("Could not create insert registration query:", err.Error())
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID, userId)

	return err

}
