package utils

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/jinzhu/gorm"
	"os"
	"encoding/base64"
	"crypto/rand"
     _ "github.com/go-sql-driver/mysql"

)

var Db  = NewRDB()

func NewRDB() *gorm.DB {
	root := os.Getenv("APP_MYSQL")
	if root == "" {
		root = "root:admin@tcp(127.0.0.1:3306)/testdb"
	}
	db, err := gorm.Open("mysql", root) //"root:admin@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}

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

func MakeErrorRequestJson(code int, msg string) ErrorRequest {
	return ErrorRequest{
		Code:code,
		Result:"failed",
		Message:msg,
	}
}