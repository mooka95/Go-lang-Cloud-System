package models

import (
	"CloudSystem/database"
	"CloudSystem/utils"

	"github.com/google/uuid"
)

type User struct {
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	FirstName  string `json:"firstName" binding:"required"`
	LastName   string `json:"lastName" binding:"required"`
	identifier string
}

func NewUser(email, password, firstName, lastName string) *User {
	return &User{
		Email:     email,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
	}

}
func (user *User) SetIdentifier(identifier string) {
	user.identifier = identifier
}
func (user *User) AddUser() (*string, error) {
	sqlStatement := `INSERT INTO users (email, password,first_name, last_name, identifier) VALUES ($1, $2,$3,$4, $5) RETURNING identifier`
	stmt, err := database.DB.Prepare(sqlStatement)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	// hash password
	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	var id string
	err = stmt.QueryRow(user.Email, hashPassword, user.FirstName, user.LastName, uuid.New()).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &id, nil
}
