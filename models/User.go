package models

import (
	"database/sql"
	"perfume/packages/auth"
	"perfume/packages/database"
	"perfume/packages/exception"
)

// User struct
type User struct {
	*database.Model
}

// GetUsers is the method that get all users
func (user *User) GetUsers() *sql.Rows {
	defer user.Model.Init().Close()
	rows := user.Model.Query("select * from users")
	return rows
}

//GetUserByEmail is the method that get a user by email
func (user *User) GetUserByEmail(email string) (string, string, string) {
	user.Model.Init()
	defer user.Model.Close()
	var firstName, lastName, hashedPassword string
	rows := user.Model.Query("select first_name, last_name, password from users where email='" + email + "';")
	for rows.Next() {
		err := rows.Scan(&firstName, &lastName, &hashedPassword)
		recoder.Write(err)
	}
	return firstName, lastName, hashedPassword
}

// AddUser is the method that add user
func (user *User) AddUser(firstName, lastName, email, password string) {
	defer user.Model.Init().Close()
	encodePassword := auth.PasswordEncode(password)
	user.Model.Query("INSERT INTO users (`first_name`, `last_name`, `email`, `password`) VALUES (\"" + firstName + "\",\"" + lastName + "\",\"" + email + "\",\"" + encodePassword + "\");")
}

// InitUser to init user
func InitUser() User {
	return User{Model: database.InitModel()}
}
