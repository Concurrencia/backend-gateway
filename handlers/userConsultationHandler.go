package handlers

import (
	"apigo/algorithm"
	"apigo/models"
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// GetAllConsultations godoc
// @Summary Get all Consultations
// @Tags consultations
// @Produce  json
// @Success 200 {object} models.Consultations
// @Router /consultations [get]
func GetAllConsultations(rw http.ResponseWriter, r *http.Request) {
	con, _ := net.Dial("tcp", "localhost:9010")
	defer con.Close()

	bufferIn := bufio.NewReader(con)
	msg, _ := bufferIn.ReadString('\n')
	msg = strings.TrimSpace(msg)

	consults := models.Consultations{}
	json.Unmarshal([]byte(msg), &consults)

	fmt.Println(consults)
	sendData(rw, consults, http.StatusOK)
}

// GetAllConsultationsByUserId godoc
// @Summary Get all Consultations from a User
// @Tags consultations
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Consultations
// @Param id path string true "User ID"
// @Router /users/{id}/consultations [get]
func GetAllConsultationsByUserId(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	con, _ := net.Dial("tcp", "localhost:9012")
	defer con.Close()

	fmt.Fprintln(con, userId)

	bufferIn := bufio.NewReader(con)
	msg, _ := bufferIn.ReadString('\n')
	msg = strings.TrimSpace(msg)
	if msg == "nil" {
		sendError(rw, http.StatusNotFound, "User not found")
		return
	}

	msg, _ = bufferIn.ReadString('\n')
	msg = strings.TrimSpace(msg)

	consults := models.Consultations{}
	json.Unmarshal([]byte(msg), &consults)

	fmt.Println(consults)
	sendData(rw, consults, http.StatusOK)
}

// CreateConsultation godoc
// @Summary Creates a new Consultation
// @Tags consultations
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param Create Consultation Dto body models.CreateConsultationDto true "Create Consultation"
// @Success 200 {object} models.Consultation
// @Router /users/{id}/consultations [post]
func CreateConsultation(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	con, _ := net.Dial("tcp", "localhost:9011")
	defer con.Close()

	fmt.Fprintln(con, userId)

	consult := models.Consultation{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&consult); err != nil {
		sendError(rw, http.StatusUnprocessableEntity, "Unprocessable Entity. "+err.Error())
	} else {
		fmt.Println("ENTRO")
		consult.Result = algorithm.RandomForestPredict(consult.LoanAmount, consult.CreditHistory, consult.PropertyAreaNum, consult.CantMultas, consult.NivelGravedadNum)
		byteInfo, _ := json.Marshal(consult)
		fmt.Fprintln(con, string(byteInfo))

		bufferIn := bufio.NewReader(con)
		msg, _ := bufferIn.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "nil" {
			sendError(rw, http.StatusNotFound, "User not found")
			return
		}
		msg, _ = bufferIn.ReadString('\n')
		msg = strings.TrimSpace(msg)

		newConsult := models.Consultation{}
		json.Unmarshal([]byte(msg), &newConsult)
		sendData(rw, newConsult, http.StatusCreated)
	}

}

// DeleteConsultation godoc
// @Summary Creates a new Consultation
// @Tags consultations
// @Accept  json
// @Produce  json
// @Param id path integer true "User ID"
// @Success 200 {object} models.Consultation
// @Router /consultations/{id} [delete]
func DeleteConsultation(rw http.ResponseWriter, r *http.Request) {

}
