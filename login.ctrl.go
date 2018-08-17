package main

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"log"
)

var user = User{}

type LoginFormData struct {
	Email string `json: "email"`
	Pwd string `json: "pwd"`
}

type CreateLoginFormDate struct {
	LoginForm LoginFormData `json: "loginForm"`
	Name string `json: "name"`
}

type loginResult struct {
	Result bool `json: "result"`
	Token string `json: "token"`
}

func UserRegisteration(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	createLoginFormDate := CreateLoginFormDate{}
	jsEerr := json.Unmarshal(body, &createLoginFormDate)
	if jsEerr != nil {
		log.Fatal(jsEerr.Error())
	}
	result, err := user.Create(createLoginFormDate)

	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		js, _ := json.Marshal(result)
		w.Write(js)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {

	loginFormData := LoginFormData{Email:"", Pwd:""}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &loginFormData)

	loginUser := user.GetUser(loginFormData)

	isvalid := isVaildPassword(loginUser, loginFormData.Pwd)

	if isvalid {
		result, _ := json.Marshal(loginResult{
			Result:isvalid,
			Token: loginUser.Access_token,
		})
		w.Write(result)
	} else {
		result, _ := json.Marshal(loginResult{
			Result:isvalid,
			Token: "",
		})
		w.Write(result)
	}
}