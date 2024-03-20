package database

import (
	"database/sql"
	"event-booking/configurations"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitConnection(cfg *configurations.Config) {
	var err error
	// Create a connection pool to the database
	user := cfg.Database.Username
	password := cfg.Database.Password
	host := cfg.Database.Host
	port, _ := strconv.ParseInt(cfg.Database.Port, 10, 64)
	dbname := cfg.Database.DbName

	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		user, password, host, port, dbname)
	fmt.Println(connStr)
	//, port, dbname)
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
