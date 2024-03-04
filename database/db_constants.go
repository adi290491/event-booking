package database

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "S@10dulkar"
	dbname   = "event_booking"

	create_events_table = `
	CREATE TABLE IF NOT EXISTS events (
		event_id       SERIAL PRIMARY KEY,
		event_name      VARCHAR(128) NOT NULL,
		description    VARCHAR(255) NOT NULL,
		location      VARCHAR(50) NOT NULL,
		date_time	timestamp NOT NULL,
		user_id		INTEGER,
		CONSTRAINT fk_user_id
			FOREIGN KEY(user_id)
				REFERENCES users(id)
	  )`

	create_users_table = `
	CREATE TABLE IF NOT EXISTS users (
		id       SERIAL PRIMARY KEY,
		email      VARCHAR(128) NOT NULL UNIQUE,
		password    VARCHAR(255) NOT NULL
	  )
	`

	create_users_event_registration_table = `
	CREATE TABLE IF NOT EXISTS user_event_registrations (
		registration_id SERIAL PRIMARY KEY,
		event_id INTEGER,
		user_id INTEGER,
		CONSTRAINT fk_event_id
			FOREIGN KEY(event_id)
				REFERENCES events(event_id)
				ON DELETE SET NULL,
		CONSTRAINT fk_u_id
			FOREIGN KEY(user_id)
				REFERENCES users(id)
				ON DELETE SET NULL
	)
	`
)
