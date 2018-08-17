package main

import (
	_"github.com/go-sql-driver/mysql"
	"errors"
)

type User struct {
	id int `json: "id" gorm:"primary_key`
	Name string `json: "name"`
	Email string `json: "email"`
	Pwd string `json: "pwd"`
	Recommend_id string `json: "recommend_id"`
	Permission int `json: "permission"`
	Access_token string `json: "acess_token"`
}

func (user User) GetUser(data LoginFormData) User {
	db.Where("email = ?", data.Email).First(&user)
	return user
}

func (user User) TokenCheck(token string) bool {
	db.Where("access_token = ?", token).First(&user)
	return user.Access_token != token
}

func checkUserDuplicate(email string) bool {
	var num int
	db.Model(&User{}).Where("email = ?", email).Count(&num)
	return num < 1
}

func (user User) Create(data CreateLoginFormDate) (User, error) {
	isVaild := checkUserDuplicate(data.LoginForm.Email)
	data.LoginForm.Pwd = string(MakePassword(data.LoginForm.Pwd))
	 newUser := &User{
			Name:data.Name,
			Email:data.LoginForm.Email,
			Pwd:data.LoginForm.Pwd,
			Recommend_id:RandToken(),
			Permission:0,
			Access_token:RandToken(),
		}
	if isVaild {
		db.Create(newUser)
		return *newUser, nil
	} else {
		return *newUser, errors.New("login failed, email duplicate")
	}
}