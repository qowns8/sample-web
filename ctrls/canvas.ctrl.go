package ctrls

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"log"
	"github.com/qowns8/sample-web/models"
)

type NewCanvasForm struct {
	Name string `json:"name"`
	Intro string `json:"Intro"`
}

func GetMyCanvas (w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("access_token")
	user := models.User{}
	permission := models.Canvas_permission{}
	userInfo := user.GetUserByToken(token)
	canvas_permissions := permission.GetCanvasPermissionByUserId(userInfo.Id)
	js, _ := json.Marshal(canvas_permissions)
	w.Write(js)
}

func CreateCanvas (w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("access_token")
	body, _ := ioutil.ReadAll(r.Body)
	var newCanvasForm NewCanvasForm
	json.Unmarshal(body, &newCanvasForm)
	canvas := models.Canvas{
		Name:newCanvasForm.Name,
		Intro:newCanvasForm.Intro,
	}
	err := canvas.CreateCanvasByToken(token, &canvas)
	if err != nil {
		log.Println(err.Error())
		w.Write([]byte("create fail"))
	}
	js, _ := json.Marshal(canvas)
	w.Write(js)
}

//todo 캔버스 수정 (권한 2 이상), 삭제 (권한 3)

func EditCanvas (w http.ResponseWriter, r *http.Request) {

}