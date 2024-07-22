package models

import (
	"demo/restserver/db"
	"demo/restserver/utils"
	"errors"
	"fmt"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
  INSERT INTO Users(email, password)
  VALUES (?, ?)
  `
	statement, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println("User - Save() - Failed to prepare statement")
		return err
	}
	defer statement.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		fmt.Println("User - Save() - Failed to hash password")
		return err
	}
	result, err := statement.Exec(u.Email, hashedPassword)
	if err != nil {
		fmt.Println("User - Save() - Failed to execute statement")
		return err
	}
	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}

func (u *User) ValidateCredentials() error {
	query := `
  SELECT id, password
  FROM Users
  WHERE email = ?
  `
	row := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return errors.New("Invalid credentials")
	}
	passwordIsValid := utils.CheckPassword(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("Invalid credentials")
	}
	return nil
}
