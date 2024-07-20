package models

import (
	"demo/restserver/db"
	"fmt"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
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
	result, err := statement.Exec(u.Email, u.Password)
	if err != nil {
		fmt.Println("User - Save() - Failed to execute statement")
		return err
	}
	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}
