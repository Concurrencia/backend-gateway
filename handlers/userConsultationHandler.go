package handlers

import (
	"apigo/db"
	"apigo/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetAllConsultations godoc
// @Summary Get all Consultations
// @Tags consultations
// @Produce  json
// @Success 200 {object} models.Consultations
// @Router /consultations [get]
func GetAllConsultations(rw http.ResponseWriter, r *http.Request) {
	consultations := models.Consultations{}
	db.Database.Find(&consultations)
	sendData(rw, consultations, http.StatusOK)
}

// GetAllConsultationsByUserId godoc
// @Summary Get all Consultations from a User
// @Tags consultations
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Consultations
// @Param id path integer true "User ID"
// @Router /users/{id}/consultations [get]
func GetAllConsultationsByUserId(rw http.ResponseWriter, r *http.Request) {
	user, err := getUserById(r)
	if err != nil {
		sendError(rw, http.StatusNotFound, "User not found")
		return
	}

	consultations := models.Consultations{}
	db.Database.Find(&consultations, "user_id = ?", user.ID)
	sendData(rw, consultations, http.StatusOK)
}

// CreateConsultation godoc
// @Summary Creates a new Consultation
// @Tags consultations
// @Accept  json
// @Produce  json
// @Param id path integer true "User ID"
// @Param Create Consultation Dto body models.CreateConsultationDto true "Create Consultation"
// @Success 200 {object} models.Consultation
// @Router /users/{id}/consultations [post]
func CreateConsultation(rw http.ResponseWriter, r *http.Request) {

	user, err := getUserById(r)
	if err != nil {
		sendError(rw, http.StatusNotFound, "User not found")
		return
	}

	consultation := models.Consultation{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&consultation); err != nil {
		sendError(rw, http.StatusUnprocessableEntity, "Unprocessable Entity. "+err.Error())
	} else {
		consultation.UserID = user.ID
		db.Database.Save(&consultation)
		sendData(rw, consultation, http.StatusCreated)
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

	if consultation, err := getConsultationById(r); err != nil {
		sendError(rw, http.StatusNotFound, "Consultation not found")
	} else {
		db.Database.Delete(&consultation)
		sendData(rw, consultation, http.StatusOK)
	}
}

func getConsultationById(r *http.Request) (*models.Consultation, error) {

	vars := mux.Vars(r)
	consultationId, _ := strconv.Atoi(vars["id"])
	consultation := models.Consultation{}

	if err := db.Database.First(&consultation, consultationId); err.Error != nil {
		return nil, err.Error
	} else {
		return &consultation, nil
	}
}
