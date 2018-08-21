package main

import (
	"log"
	"net/http"
	"github.com/qowns8/sample-web/ctrls"
)

func main() {
	println("server start")
	router := ctrls.Route()
	log.Fatal(http.ListenAndServe(":5000", router))
}
