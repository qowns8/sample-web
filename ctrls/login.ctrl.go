package ctrls

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"log"
	"github.com/qowns8/sample-web/models"
	"github.com/qowns8/sample-web/utils"
)


func UserRegisteration(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	createLoginFormDate := utils.CreateLoginFormDate{}
	jsEerr := json.Unmarshal(body, &createLoginFormDate)
	user := models.User{}
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

	loginFormData :=  utils.LoginFormData{Email:"", Pwd:""}
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &loginFormData)
	user := models.User{}
	loginUser := user.GetUser(loginFormData)

	isvalid := models.IsVaildPassword(loginUser, loginFormData.Pwd)

	if isvalid {
		result, _ := json.Marshal(utils.LoginResult{
			Result:isvalid,
			Token: loginUser.Access_token,
		})
		w.Write(result)
	} else {
		result, _ := json.Marshal(utils.LoginResult{
			Result:isvalid,
			Token: "",
		})
		w.Write(result)
	}
}