package models

import (
	_"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"errors"
	"github.com/sample-web/ctrl"
)

type User struct {
	gorm.Model
	id int `json: "id"`
	name string `json: "name"`
	email string `json: "email"`
	pwd string `json: "pwd"`
	recommend_id string `json: "recommend_id"`
	permission int `json: "permission"`
	access_token string `json: "acess_token"`
}

func (user User) GetUser(data LoginFormData) User {
	db.Where("email = ?", data.email).First(&user)
	return user
}

func checkUserDuplicate(email string) bool {
	var num int
	db.Where("email = ?", email).Count(num)
	return num < 1
}

func (user User) Create(data CreateLoginFormDate) (User, error) {
	isVaild := checkUserDuplicate(data.loginFormData.email)
	data.loginFormData.pwd = string(MakePassword(data.loginFormData.pwd))
	 newUser := &User{
			name:data.name,
			email:data.loginFormData.email,
			pwd:data.loginFormData.pwd,
			recommend_id:RandToken(),
			permission:0,
			access_token:RandToken(),
		}

	if isVaild {
		db.Create(newUser)
		return *newUser, nil
	} else {
		return *newUser, errors.New("login failed, email duplicate")
	}
}