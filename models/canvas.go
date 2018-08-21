package models

import (
	"github.com/jinzhu/gorm"
	"errors"
	"github.com/qowns8/sample-web/utils"
	"log"
)

type Canvas struct {
	Id int `json:"id" gorm:"primary_key`
	Name string `json:"name"`
	Intro string `json:"intro"`
	Problem string `json: "problem"`
	Unique_value_propostion string `json: "unique_value_propostion"`
	Solution string `json: "solution"`
	Channel string `json: "channel"`
	Cost_structure string `json: "cost_structure"`
	Revenue_stream string `json: "revenue_stream"`
	Key_metric string `json: "key_metric"`
}


func (c *Canvas) CreateCanvasByToken(token string, canvas *Canvas) error {
	tx := utils.Db.Begin()
	defer tx.Commit()
	isCanvasCreateSuccess := createOnlyCanvas(tx, canvas)
	isPermissionCreateSuccess := CreateCanvasPermission(tx, token, canvas.Id, 3)
	if isCanvasCreateSuccess == false || isPermissionCreateSuccess == false {
		tx.Rollback()
		return errors.New("create Canvas Fail")
	}
	return nil
}

func createOnlyCanvas(tx *gorm.DB, canvas *Canvas) bool {
	err := tx.Create(canvas).Error
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

func (c *Canvas) UpdateCanvas(token string, canvas *Canvas) error {
	tx := utils.Db.Begin()
	defer tx.Commit()
	user := User{}
	user = user.GetUserByToken(token)
	permission := Canvas_permission{}
	permission = permission.GetCanvasPermission(canvas.Id, user.Id)
	if permission.Permission > 2 {
		tx.Update(&canvas)
	} else {
		return errors.New("")
	}

}