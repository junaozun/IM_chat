package controller

import (
	"net/http"

	"github.com/junaozun/IM_chat/global"
	"github.com/junaozun/IM_chat/model"
	"github.com/junaozun/IM_chat/service"
	"github.com/junaozun/IM_chat/utils"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	_ = request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	password := request.PostForm.Get("password")
	user, err := service.UserService{}.Login(mobile, password)
	if err != nil {
		global.Response(writer, -1, err.Error(), nil)
	} else {
		global.Response(writer, 0, "登陆成功", user)
	}
	return
}

func Register(writer http.ResponseWriter, request *http.Request) {
	_ = request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	password := request.PostForm.Get("password")
	nickname := request.PostForm.Get("nickname")
	avatar := ""
	sex := model.SEX_UNKNOW

	user, err := service.UserService{}.Register(mobile, password, nickname, avatar, sex)

	if err != nil {
		global.Response(writer, -1, err.Error(), nil)
	} else {
		global.Response(writer, 0, "", user)
	}
	return
}

func GetUser(writer http.ResponseWriter, request *http.Request) {
	var user model.User
	if err := utils.Bind(request, &user); err != nil {
		global.ResponseFail(writer, err.Error())
	}
	user, err := service.UserService{}.GetUserById(user.Id)
	if err != nil {
		global.ResponseFail(writer, err.Error())
	} else {
		global.ResponseOk(writer, user, "获取用户信息成功")
	}
	return
}
