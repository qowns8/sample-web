package main

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
)

var user = User{}

type LoginFormData struct {
	email string `json: "email"`
	pwd string `json: "pwd"`
}

type CreateLoginFormDate struct {
	loginFormData LoginFormData `json: "loginForm"`
	name string `json: "name"`
}

type loginResult struct {
	result bool `json: "result"`
	token string `json: "token"`
}

func UserRegisteration(w http.ResponseWriter, r *http.Request) {
	createLoginFormData := CreateLoginFormDate{}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, createLoginFormData)
	result, err := user.Create(createLoginFormData)

	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		js, _ := json.Marshal(result)
		w.Write(js)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {

	loginFormData := LoginFormData{email:"", pwd:""}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &loginFormData)

	loginUser := user.GetUser(loginFormData)

	isvalid := isVaildPassword(loginUser, loginFormData.pwd)

	if isvalid {
		result, _ := json.Marshal(loginResult{
			result:isvalid,
			token: loginUser.access_token,
		})
		w.Write(result)
	} else {
		result, _ := json.Marshal(loginResult{
			result:isvalid,
			token: "",
		})
		w.Write(result)
	}
}