package main

import (
	"net/http"
	"encoding/json"
	"log"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func AccessMiddleware() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("access_token")
			isVaild := user.TokenCheck(token)

			if isVaild || token == "" {
				reqJson := makeErrorRequestJson(405, "access token invailed")
				req, _ := json.Marshal(reqJson)
				w.Write(req)
			} else {
				f(w, r)
			}
		}
	}
}
func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			defer func() {
			}()
			log.Println("visit : ", r.URL.Path)
			f(w, r)
		}
	}
}

func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func BasicLogic(f http.HandlerFunc) http.HandlerFunc {
	return Chain(f, AccessMiddleware(), Logging())
}