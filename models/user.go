package models

import (
	"errors"
	"event-booking/database"
	"event-booking/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required" min=10`
}

func (u *User) Save() error {

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}
	stmt, err := database.DB.Prepare(INSERT_USER)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(u.Email, hashedPassword)

	return err
}

func (u *User) ValidateCredentials() error {

	row := database.DB.QueryRow(USER_BY_EMAIL, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("Invalid Credentials")
	}

	isPasswordMatch := utils.IsCorrectPassword(u.Password, retrievedPassword)

	if !isPasswordMatch {
		return errors.New("Invalid Credentials")
	}

	return nil
}
