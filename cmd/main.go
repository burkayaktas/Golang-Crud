package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tutorials/go/crud/domain/db"
	"github.com/tutorials/go/crud/handlers"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/users", h.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", h.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/users", h.AddUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", h.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", h.DeleteUser).Methods(http.MethodDelete)

	log.Println("API is Running!")
	http.ListenAndServe(":8080", router)
}
