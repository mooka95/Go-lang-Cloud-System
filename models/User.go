package models

import (
	"CloudSystem/database"
	"CloudSystem/queries"
	"CloudSystem/utils"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	Id         int64
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	City       string `json:"city" binding:"required"`
	Street     string `json:"street" binding:"required"`
	Country    string `json:"country" binding:"required"`
	FirstName  string
	LastName   string
	Identifier string
}

func NewUser(email, password, firstName, lastName string) *User {
	return &User{
		Email:     email,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
	}

}

func (user *User) AddUser(currentConnection *sql.Tx) (*User, error) {
	sqlStatement := `INSERT INTO users (email, password,first_name, last_name, identifier) VALUES ($1, $2,$3,$4, $5) RETURNING identifier,id`
	stmt, err := currentConnection.Prepare(sqlStatement)
	fmt.Println("err from here")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	// hash password
	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	// var id int64
	// var identifier string
	err = stmt.QueryRow(user.Email, hashPassword, user.FirstName, user.LastName, uuid.New()).Scan(&user.Identifier, &user.Id)

	if err != nil {
		return nil, err
	}

	return user, nil
}
func (user *User) AddUserAddress(currentConnection *sql.Tx) (*string, error) {
	stmt, err := currentConnection.Prepare(queries.QueryAddressMap["insertAddress"])

	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	var id string
	err = stmt.QueryRow(user.City, user.Street, user.Country, uuid.New(), user.Id).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &id, nil
}
func GetUserId(identifier string) (string, error) {
	var id string
	err := database.DB.QueryRow(queries.QueryUserMap["getUserIdFromIdentifier"], identifier).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}
func GetUserByEmail(email string) (*User, error) {
	row := database.DB.QueryRow(queries.QueryUserMap["getUserByEmail"], email)
	var user User
	err := row.Scan(&user.Identifier, &user.Password, &user.Id)

	if err != nil {
		return nil, errors.New("email or password not valid")
	}
	return &user, nil
}

func (user *User) ValidatePassword(userInputPassword string) error {

	passwordIsValid := utils.CheckPasswordHash(userInputPassword, user.Password)

	if !passwordIsValid {
		return errors.New("email or password invalid")
	}

	return nil
}
