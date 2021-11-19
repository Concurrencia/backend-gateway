package handlers

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

func createDefaultResponse(rw http.ResponseWriter, status int) Response {
	return Response{
		Status:         status,
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

func sendData(rw http.ResponseWriter, data interface{}, status int) {
	response := createDefaultResponse(rw, status)
	response.Data = data
	response.Message = "success"
	response.send()
}

func sendError(rw http.ResponseWriter, status int, message string) {
	response := createDefaultResponse(rw, status)
	response.Message = message
	response.send()
}
