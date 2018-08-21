package main

import (
	"log"
	"net/http"
	"github.com/qowns8/sample-web/router"
)

func main() {
	println("server start")
	router := router.Route()
	log.Fatal(http.ListenAndServe(":5000", router))
}
