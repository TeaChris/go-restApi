package models

import (
	"errors"
	"rest-api/db"
	"rest-api/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"-"`
}


func (u *User) Save() error {
	query := `
	INSERT INTO users (email, password) VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = userId
	return nil
}

func (u User) ValidateCredentials() error {
    query := "SELECT password FROM users WHERE email = ?"
    row := db.DB.QueryRow(query, u.Email)

    var retrievedPassword string
    err := row.Scan(&retrievedPassword)
    if err != nil {
        return err
    }


    passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
    if !passwordIsValid {
        return errors.New("invalid credentials")
    }

    return nil
}
