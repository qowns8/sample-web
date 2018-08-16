package main

import (
	"log"
	"net/http"
	"github.com/sample-web/ctrl"
)

func main() {
	router := route()
	log.Fatal(http.ListenAndServe(":5000", router))
}
