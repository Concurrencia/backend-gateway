package main

import (
	"apigo/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func requestHandler() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/users/", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/api/users/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/api/users/", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	// run with fresh
	requestHandler()
}
