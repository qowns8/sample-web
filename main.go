package main

import (
	"log"
	"net/http"
)

func main() {
	println("server start")
	router := Route()
	log.Fatal(http.ListenAndServe(":5000", router))
}
