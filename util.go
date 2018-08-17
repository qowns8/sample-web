package main

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/jinzhu/gorm"
	"os"
	"encoding/base64"
	"crypto/rand"
)

var db  = NewRDB()

func NewRDB() *gorm.DB {
	root := os.Getenv("APP_MYSQL")
	if root == "" {
		root = "root:admin@tcp(127.0.0.1:3306)/testdb"
	}
	db, _ := gorm.Open("mysql", root) //"root:admin@tcp(127.0.0.1:3306)/testdb")
	return db
}

func RandToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func MakePassword(str string) []byte {
	token, _ := bcrypt.GenerateFromPassword([]byte(str) ,10)
	return token
}

func isVaildPassword(user User, str string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Pwd), []byte(str))
	return err == nil
}

func makeErrorRequestJson(code int, msg string) ErrorRequest {
	return ErrorRequest{
		Code:code,
		Result:"failed",
		Message:msg,
	}
}