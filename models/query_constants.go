package models

const (
	INSERT_EVENT = `
	INSERT INTO events(event_name, description, location, date_time, user_id)
	VALUES ($1, $2, $3, $4, $5) RETURNING event_id`

	ALL_EVENTS = `SELECT * FROM events`

	EVENT_BY_ID = `SELECT * FROM events WHERE event_id = $1`

	UPDATE_EVENT = `UPDATE events
	SET
	event_name = $2,
	description = $3,
	location = $4,
	date_time = $5
	WHERE event_id = $1`

	DELETE_EVENT = `
	DELETE FROM events
	WHERE event_id = $1
	`

	INSERT_USER = `
	INSERT INTO users (email, password)
	VALUES ($1, $2)
	`

	USER_BY_EMAIL = `
	SELECT id, password FROM users
	WHERE email = $1
	`

	INSERT_INTO_REGISTRATION = `
	INSERT INTO user_event_registrations (event_id, user_id)
	VALUES ($1, $2) RETURNING registration_id
	`

	CANCEL_EVENT_REGISTRATION = `
	DELETE FROM user_event_registrations
	WHERE event_id = $1 AND user_id = $2
	`
)
