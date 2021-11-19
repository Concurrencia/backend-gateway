package handlers

import (
	"apigo/db"
	"apigo/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	users := models.Users{}
	db.Database.Find(&users)
	sendData(rw, users, http.StatusOK)
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	if user, err := getUserById(r); err != nil {
		sendError(rw, http.StatusNotFound, "User not found")
	} else {
		sendData(rw, user, http.StatusOK)
	}

}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	//Obtener registro
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		sendError(rw, http.StatusUnprocessableEntity, "Unprocessable Entity")
	} else {
		db.Database.Save(&user)
		sendData(rw, user, http.StatusCreated)
	}
}

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
		updated_user.Id = old_user.Id
		db.Database.Save(&updated_user)
		sendData(rw, updated_user, http.StatusOK)
	}

}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {

	if user, err := getUserById(r); err != nil {
		sendError(rw, http.StatusNotFound, "User not found")
	} else {
		db.Database.Delete(&user)
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
