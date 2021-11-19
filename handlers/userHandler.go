package handlers

import (
	"apigo/db"
	"apigo/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetUsers godoc
// @Summary Get all Users
// @Description Get details of all Users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Users
// @Router /users [get]
func GetUsers(rw http.ResponseWriter, r *http.Request) {
	users := models.Users{}
	db.Database.Find(&users)
	sendData(rw, users, http.StatusOK)
}

// GetUser godoc
// @Summary Retrieves user based on given ID
// @Tags users
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func GetUser(rw http.ResponseWriter, r *http.Request) {
	if user, err := getUserById(r); err != nil {
		sendError(rw, http.StatusNotFound, "User not found")
	} else {
		sendData(rw, user, http.StatusOK)
	}

}

// CreateUser godoc
// @Summary Creates a new User
// @Tags users
// @Accept  json
// @Produce  json
// @Param RequestUserDto body models.RequestUserDto true "Create User"
// @Success 200 {object} models.User
// @Router /users/register [post]
func CreateUser(rw http.ResponseWriter, r *http.Request) {
	//Obtener registro
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		sendError(rw, http.StatusUnprocessableEntity, "Unprocessable Entity. "+err.Error())
	} else {
		db.Database.Save(&user)
		fmt.Println(user)
		sendData(rw, user, http.StatusCreated)
	}
}

// UpdateUser godoc
// @Summary Update a  user based on given ID
// @Description Update a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path integer true "User ID"
// @Param user body models.RequestUserDto true "Update user"
// @Success 200 {object} models.User
// @Router /users/{id} [put]
func UpdateUser(rw http.ResponseWriter, r *http.Request) {

	updated_user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&updated_user); err != nil {
		sendError(rw, http.StatusUnprocessableEntity, "Unprocessable Entity")
		return
	}

	if old_user, err := getUserById(r); err != nil {
		sendError(rw, http.StatusNotFound, "User not found")
	} else {
		updated_user.ID = old_user.ID
		db.Database.Save(&updated_user)
		sendData(rw, updated_user, http.StatusOK)
	}

}

// DeleteUser godoc
// @Summary Deletes user based on given ID
// @Tags users
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [delete]
func DeleteUser(rw http.ResponseWriter, r *http.Request) {

	if user, err := getUserById(r); err != nil {
		sendError(rw, http.StatusNotFound, "User not found")
	} else {
		db.Database.Delete(&user)
		sendData(rw, user, http.StatusOK)
	}
}

// Login godoc
// @Summary Login
// @Description Login
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.UserLogin true "Login"
// @Success 200 {object} models.User
// @Router /users/login [post]
func Login(rw http.ResponseWriter, r *http.Request) {
	user := models.User{}
	login_form := models.UserLogin{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&login_form); err != nil {
		sendError(rw, http.StatusUnprocessableEntity, "Unprocessable Entity")
		return
	}

	if err := db.Database.First(&user, "email = ? AND password = ?", login_form.Email, login_form.Password); err.Error != nil {
		sendError(rw, http.StatusNotFound, "Wrong email or password")
	} else {
		sendData(rw, user, http.StatusOK)
	}

}

func getUserById(r *http.Request) (*models.User, error) {
	//Obtener ID
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	user := models.User{}

	if err := db.Database.First(&user, userId); err.Error != nil {
		return nil, err.Error
	} else {
		return &user, nil
	}
}
