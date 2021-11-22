package handlers

import (
	"apigo/models"
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// GetUsers godoc
// @Summary Get all Users
// @Description Get details of all Users
// @Tags users
// @Produce  json
// @Success 200 {object} models.Users
// @Router /users [get]
func GetUsers(rw http.ResponseWriter, r *http.Request) {

	con, _ := net.Dial("tcp", "localhost:9001")
	defer con.Close()

	bufferIn := bufio.NewReader(con)
	msg, _ := bufferIn.ReadString('\n')
	msg = strings.TrimSpace(msg)

	users := models.Users{}
	json.Unmarshal([]byte(msg), &users)

	fmt.Println(users)
	sendData(rw, users, http.StatusOK)
}

// GetUser godoc
// @Summary Retrieves user based on given ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func GetUser(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userId := vars["id"]

	con, _ := net.Dial("tcp", "localhost:9002")
	defer con.Close()

	fmt.Fprintln(con, userId)

	bufferIn := bufio.NewReader(con)
	msg, _ := bufferIn.ReadString('\n')
	msg = strings.TrimSpace(msg)

	user := models.User{}
	json.Unmarshal([]byte(msg), &user)

	if user.ID == "" {
		sendError(rw, http.StatusNotFound, "User not found with Id")
		return
	}

	sendData(rw, user, http.StatusOK)

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

		if user.Email == "" || user.Password == "" {
			sendError(rw, http.StatusBadRequest, "Username or password can't be empty")
			return
		}

		con, _ := net.Dial("tcp", "localhost:9000")
		defer con.Close()

		byteInfo, _ := json.Marshal(user)
		fmt.Fprintln(con, string(byteInfo))

		bufferIn := bufio.NewReader(con)
		hashId, _ := bufferIn.ReadString('\n')
		hashId = strings.TrimSpace(hashId)

		user.ID = hashId

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

}

// DeleteUser godoc
// @Summary Deletes user based on given ID
// @Tags users
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [delete]
func DeleteUser(rw http.ResponseWriter, r *http.Request) {

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
	//user := models.User{}
	login_form := models.UserLogin{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&login_form); err != nil {
		sendError(rw, http.StatusUnprocessableEntity, "Unprocessable Entity")
		return
	}

	con, _ := net.Dial("tcp", "localhost:9003")
	defer con.Close()

	fmt.Fprintln(con, login_form.Email)
	fmt.Fprintln(con, login_form.Password)

	bufferIn := bufio.NewReader(con)
	msg, _ := bufferIn.ReadString('\n')
	msg = strings.TrimSpace(msg)

	user := models.User{}
	json.Unmarshal([]byte(msg), &user)

	if user.ID == "" {
		sendError(rw, http.StatusNotFound, "User Not Found")
	} else {
		sendData(rw, user, http.StatusOK)
	}
}
