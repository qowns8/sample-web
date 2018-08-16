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
		root = "root:pwd@tcp(127.0.0.1:3306)/testdb"
	}
	db, _ := gorm.Open("mysql", root) //"root:pwd@tcp(127.0.0.1:3306)/testdb")
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
	err := bcrypt.CompareHashAndPassword([]byte(str), []byte(user.pwd))
	return err == nil
}