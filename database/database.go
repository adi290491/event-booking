package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitConnection() {

	var err error
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		user, password, host, port, dbname)
	DB, err = sql.Open("postgres", connStr)

	checkError("Error while connecting to database:", err)

	err = DB.Ping()

	checkError("Failed to receive packets from database", err)

	fmt.Println("Successfully Connected!")

	createUsersTable()
	createEventsTable()
	createUserEventRegistrationTable()
}

func createUsersTable() {
	_, err := DB.Exec(create_users_table)

	checkError("Error creating Users table", err)

	fmt.Println("Users Table created successfully!")
}

func createEventsTable() {

	_, err := DB.Exec(create_events_table)

	checkError("Error creating events table", err)

	fmt.Println("Events Table created successfully!")
}

func createUserEventRegistrationTable() {
	_, err := DB.Exec(create_users_event_registration_table)

	checkError("Error creating User Events Registration table", err)

	fmt.Println("User Events Registration Table created successfully!")
}

func checkError(message string, err error) {
	if err != nil {
		err = fmt.Errorf(message, err)
		log.Fatal(err)
		panic(err)
	}
}
