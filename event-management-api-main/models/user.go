package models

import (
	"errors"
	"project/event-management-api/db"
	"project/event-management-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
	Role     string
}

func (u *User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`
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
	u.Role = "user"
	return nil
}

func (u *User) ValidateCredentials() error {
	query := "SELECT * FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string

	err := row.Scan(&u.ID, &u.Email, &retrievedPassword, &u.Role)

	if err != nil {
		return errors.New("user not found")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("invalid password")
	}

	return nil
}

func GetAllUsers() ([]User, error) {

	query := "SELECT * FROM users"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUserByID(id int64) (User, error) {

	query := "SELECT * FROM users WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Role)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (u *User) Delete() error {
	query := "DELETE FROM users WHERE id = ?"

	stmnt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmnt.Close()

	_, err = stmnt.Exec(u.ID)

	if err != nil {
		return err
	}

	return nil
}
