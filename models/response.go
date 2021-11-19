package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status         int         `json:"status"`
	Data           interface{} `json:"data"`
	Message        string      `json:"message"`
	contentType    string
	responseWriter http.ResponseWriter
}

func createDefaultResponse(rw http.ResponseWriter) Response {
	return Response{
		Status:         http.StatusOK,
		responseWriter: rw,
		contentType:    "application/json",
	}
}

func (resp *Response) send() {
	resp.responseWriter.Header().Set("Content-Type", resp.contentType)
	resp.responseWriter.WriteHeader(resp.Status)

	output, _ := json.Marshal(&resp)
	fmt.Fprintln(resp.responseWriter, string(output))
}

func SendData(rw http.ResponseWriter, data interface{}) {
	response := createDefaultResponse(rw)
	response.Data = data
	response.send()
}

func (resp *Response) notFound() {
	resp.Status = http.StatusNotFound
	resp.Message = "Resource Not Found"
}

func SendNotFound(rw http.ResponseWriter) {
	response := createDefaultResponse(rw)
	response.notFound()
	response.send()
}

func (resp *Response) unprocessableEntity() {
	resp.Status = http.StatusUnprocessableEntity
	resp.Message = "Unprocessable Entity"
}

func SendUnprocessableEntity(rw http.ResponseWriter) {
	response := createDefaultResponse(rw)
	response.unprocessableEntity()
	response.send()
}
