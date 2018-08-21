package ctrls

import (
	"net/http"
)

type ErrorRequest struct {
	Result string `json "result"`
	Code int `json: "code"`
	Message string `json: "message"`
}


func GetRoot (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("visit /"))
}