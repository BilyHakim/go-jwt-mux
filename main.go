package main

import (
	"github.com/BilyHakim/go-jwt-mux/controllers/authcontroller"
	"github.com/BilyHakim/go-jwt-mux/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
