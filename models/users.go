package models

import (
	"apigo/db"
	"errors"
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"passowrd"`
	Email    string `json:"email"`
}

type Users []User

const UserSchema string = `CREATE TABLE users (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(100) NOT NULL,
	email VARCHAR(50),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)`

// Contruir usuario
func NewUser(username, password, email string) *User {
	user := &User{Username: username, Password: password, Email: email}
	return user
}

//Crear usuario e insertar
func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.insert()
	return user
}

// Isertar Registro

func (user *User) insert() {
	query := "INSERT users SET username=?, password=?, email=?"
	result, _ := db.Exec(query, user.Username, user.Password, user.Email)
	user.Id, _ = result.LastInsertId()
}

// Listar todo el registro
func ListUsers() (Users, error) {
	query := "SELECT id, username, password, email FROM users"
	users := Users{}
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		users = append(users, user)
	}

	return users, nil
}

func GetUserById(id int) (*User, error) {
	user := User{}
	query := "SELECT id, username, password, email FROM users where id=?"
	rows, err := db.Query(query, id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	}

	if user.Username == "" {
		return nil, errors.New("User not found")
	}

	return &user, nil
}

func (user *User) update() {
	query := "UPDATE users SET username=?, password=?, email=? WHERE id=?"
	db.Exec(query, user.Username, user.Password, user.Email, user.Id)
}

//Guardar o editar registro
func (user *User) Save() {
	if user.Id == 0 {
		user.insert()
	} else {
		user.update()
	}
}

func (user *User) Delete() {
	query := "DELETE FROM users WHERE id=?"
	db.Exec(query, user.Id)
}
