package models

type Consultation struct {
	ID               int64  `json:"id"`
	LoanAmount       string `json:"loanAmount"`
	CreditHistory    string `json:"creditHistory"`
	PropertyAreaNum  string `json:"propertyAreaNum"`
	CantMultas       string `json:"cantMultas"`
	NivelGravedadNum string `json:"nivelGravedadNum"`
	Result           string `json:"result"`
	UserID           string `json:"userId"`
}

type CreateConsultationDto struct {
	LoanAmount       string `json:"loanAmount"`
	CreditHistory    string `json:"creditHistory"`
	PropertyAreaNum  string `json:"propertyAreaNum"`
	CantMultas       string `json:"cantMultas"`
	NivelGravedadNum string `json:"nivelGravedadNum"`
	Result           string `json:"result"`
}

// Consultations
type Consultations []Consultation
