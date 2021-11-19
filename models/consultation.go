package models

import "apigo/db"

type Consultation struct {
	ID     int64  `json:"id"`
	Dato1  string `json:"dato1"`
	Dato2  string `json:"dato2"`
	Result string `json:"result"`
	UserID int64  `json:"userId"`
}

type CreateConsultationDto struct {
	Dato1  string `json:"dato1"`
	Dato2  string `json:"dato2"`
	Result string `json:"result"`
}

// Consultations
type Consultations []Consultation

func MigrarConsultations() {
	db.Database.AutoMigrate(Consultation{})
}
