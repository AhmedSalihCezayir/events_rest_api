package models

import (
	"errors"

	"example.com/events-api/db"
	"example.com/events-api/utils"
)

type User struct {
	ID       int64 `json:",omitempty"`
	Nickname string
	Email    string `binding:"required"`
	Password string `binding:"required" json:",omitempty"`
	IsAdmin  bool
}

func (u *User) Save() error {
	query := "INSERT INTO users (nickname, email, password, is_admin) VALUES (?, ?, ?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Nickname, u.Email, hashedPass, u.IsAdmin)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.ID = id

	return err
}

func (u User) Delete() error {
	query := "DELETE FROM users WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.ID)
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password, is_admin FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword, &u.IsAdmin)
	if err != nil {
		return errors.New("Invalid Credentials!")
	}

	isValidPass := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !isValidPass {
		return errors.New("Invalid Credentials!")
	}

	return nil
}

func FindUserById(userID int64) (*User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	row := db.DB.QueryRow(query, userID)

	var user User
	err := row.Scan(&user.ID, &user.Nickname, &user.Email, &user.Password, &user.IsAdmin)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetAllUsers() ([]User, error) {
	query := "SELECT id, nickname, email, is_admin FROM users"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Nickname, &user.Email, &user.IsAdmin)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
