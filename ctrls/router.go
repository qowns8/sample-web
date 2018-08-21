package ctrls

import (
	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", BasicLogic(GetRoot)).Methods("GET")
	router.HandleFunc("/login", Login).Methods("POST")
	router.HandleFunc("/login", UserRegisteration).Methods("PUT")
	router.HandleFunc("/canvas", BasicLogic(GetMyCanvas)).Methods("GET")
	router.HandleFunc("/canvas", BasicLogic(CreateCanvas)).Methods("POST")
	return router
}
