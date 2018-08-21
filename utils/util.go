package utils

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/jinzhu/gorm"
	"os"
	"encoding/base64"
	"crypto/rand"
	"github.com/qowns8/sample-web/models"
	"github.com/qowns8/sample-web/ctrls"
)

var Db  = NewRDB()

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

func IsVaildPassword(user models.User, str string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Pwd), []byte(str))
	return err == nil
}

func MakeErrorRequestJson(code int, msg string) ctrls.ErrorRequest {
	return ctrls.ErrorRequest{
		Code:code,
		Result:"failed",
		Message:msg,
	}
}