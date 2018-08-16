package main

import (
	"log"
	"net/http"
)

func main() {
	router := route()
	log.Fatal(http.ListenAndServe(":5000", router))
}
