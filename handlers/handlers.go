package handlers

import (
	"apigo/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(response http.ResponseWriter, request *http.Request) {
	if users, err := models.ListUsers(); err != nil {
		models.SendNotFound(response)
	} else {
		models.SendData(response, users)
	}
}

func GetUser(response http.ResponseWriter, request *http.Request) {

	if user, err := getUserByRequestId(request); err != nil {
		models.SendNotFound(response)
	} else {
		models.SendData(response, user)
	}
}

func CreateUser(response http.ResponseWriter, request *http.Request) {

	newUser := models.User{}
	//Recuperar el cuerpo del request
	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&newUser); err != nil {
		models.SendUnprocessableEntity(response)
	} else {
		newUser.Save()
		models.SendData(response, newUser)
	}

}

func UpdateUser(response http.ResponseWriter, request *http.Request) {

	updatedUser := models.User{}
	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&updatedUser); err != nil {
		models.SendUnprocessableEntity(response)
		return
	}

	if user, err := getUserByRequestId(request); err != nil {
		models.SendNotFound(response)
	} else {
		updatedUser.Id = user.Id
		updatedUser.Save()
		models.SendData(response, updatedUser)
	}

}

func DeleteUser(response http.ResponseWriter, request *http.Request) {

	if user, err := getUserByRequestId(request); err != nil {
		models.SendNotFound(response)
	} else {
		user.Delete()
		models.SendData(response, user)
	}

}

func getUserByRequestId(request *http.Request) (*models.User, error) {
	vars := mux.Vars(request)
	userId, _ := strconv.Atoi(vars["id"])

	if user, err := models.GetUserById(userId); err != nil {
		return nil, err
	} else {
		return user, nil
	}

}
