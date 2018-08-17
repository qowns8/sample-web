package main

import (
	"github.com/gorilla/mux"
)


func Route() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", BasicLogic(GetRoot)).Methods("GET")
	router.HandleFunc("/login", Login).Methods("POST")
	router.HandleFunc("/login", UserRegisteration).Methods("PUT")
	return router
}
