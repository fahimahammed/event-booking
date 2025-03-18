package models

import (
	"errors"
	"event-booking/db"
	"event-booking/utils"
	"fmt"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	// TODO: Save user to database
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	fmt.Println("Hash Pass: ", hashedPassword)

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = int64(id)
	return nil
}

func (u User) Login() error {
	query := "SELECT password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrivedPassword string

	err := row.Scan(&retrivedPassword)

	if err != nil {
		return err
	}

	fmt.Println(retrivedPassword, u.Password)
	passwordValid := utils.CheckPasswordHash(u.Password, retrivedPassword)

	if !passwordValid {
		return errors.New("invalid password")
	}

	return nil
}
