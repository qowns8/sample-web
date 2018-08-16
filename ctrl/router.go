package ctrl

import (
	"github.com/gorilla/mux"
)

func route() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", GetRoot).Methods("GET")
	router.HandleFunc("/login", Login).Methods("GET")
	router.HandleFunc("/login", UserRegisteration).Methods("POST")
	return router
}
