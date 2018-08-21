package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qowns8/sample-web/utils"
	"log"
)

type Canvas_permission struct {
	Id int `json:"id"`
	Canvas_id int `json:"canvas_id"`
	User_id int `json:"user_id"`
	Permission int `json:"permission"` // 0 게스트, 1 읽기, 2 등록,수정,읽기, 3 마스터 권한 (캔버스를 만든 사람)
}

func (cp Canvas_permission) GetCanvasPermission(canvas_id int, user_id int) Canvas_permission {
	ps := Canvas_permission{}
	utils.Db.
		Table("canvas_permission").
		Where("canvas_id = ? AND user_id = ?", canvas_id, user_id).
		First(&ps)
	return ps
}

func (cp Canvas_permission) GetCanvasPermissionByCanvasId(canvas_id int) []map[string]string{
	ps := []map[string]string{}
	utils.Db.
		Table("canvas_permission as cp").
		Select("cp.id, cp.Canvas_id, cp.user_id, ca.name, ca.id as canvas_id").
		Joins("left join canvas ca on  ca.id = cp.canvas_id").
		Where("canvas_id = ?", canvas_id).
		Scan(&ps)

	return ps
}

func (cp Canvas_permission) GetCanvasPermissionByUserId(user_id int) []map[string]string{
	ps := []map[string]string{}
	utils.Db.
		Table("canvas_permission as cp").
		Select("cp.id, cp.Canvas_id, cp.user_id, ca.name, ca.id as canvas_id").
		Joins("left join canvas ca on  ca.id = cp.canvas_id").
		Where("user_id = ?", user_id).
		Scan(&ps)
	return ps
}


// permission 3 이상일 경우에만 등록 가능함
func CreateCanvasPermission(tx *gorm.DB,  token string, canvas_id int, permission int) bool {
	// 토큰 유효성 확인
	user := User{}
	user = user.GetUserByToken(token)
	if user.Access_token == "" {
		return false
	}
	// 캔버스 권한 생성
	canvas_permission := Canvas_permission{
		Canvas_id:canvas_id,
		User_id:user.Id,
		Permission:permission,
	}
	err := tx.Create(&canvas_permission).Error
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

/* todo permission 3 이상일 경우 유저들을 캔버스 그룹에서 수정, 삭제

func DeleteCanvasPermission() {

}

func UpdateCanvasPermission() {

}*/