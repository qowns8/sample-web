package router

import (
	"github.com/gorilla/mux"
	"github.com/qowns8/sample-web/middleware"
	"github.com/qowns8/sample-web/ctrls"
)

func Route() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/test", ctrls.GetRoot).Methods("GET")
	router.HandleFunc("/", middleware.BasicLogic(ctrls.GetRoot)).Methods("GET")
	router.HandleFunc("/login", ctrls.Login).Methods("POST")
	router.HandleFunc("/login", ctrls.UserRegisteration).Methods("PUT")
	router.HandleFunc("/canvas", middleware.BasicLogic(ctrls.GetMyCanvas)).Methods("GET")
	router.HandleFunc("/canvas", middleware.BasicLogic(ctrls.CreateCanvas)).Methods("POST")
	return router
}
