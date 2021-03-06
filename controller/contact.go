package controller

import (
	"log"
	"net/http"

	"github.com/junaozun/IM_chat/global"
	"github.com/junaozun/IM_chat/model"
	"github.com/junaozun/IM_chat/model/request"
	"github.com/junaozun/IM_chat/service"
	"github.com/junaozun/IM_chat/utils"
)

var contactService service.ContactService

func LoadFriend(w http.ResponseWriter, req *http.Request) {
	var arg request.ContactArg
	// 如果这个用的上,那么可以直接
	if err := utils.Bind(req, &arg); err != nil {
		global.ResponseFail(w, err.Error())
		return
	}
	users := contactService.SearchFriend(arg.Userid)
	global.ResponseOkList(w, users, len(users))
}

func LoadCommunity(w http.ResponseWriter, req *http.Request) {
	var arg request.ContactArg
	// 如果这个用的上,那么可以直接
	if err := utils.Bind(req, &arg); err != nil {
		global.ResponseFail(w, err.Error())
		return
	}
	comunitys := contactService.SearchComunity(arg.Userid)
	GroupPeopleMap, _ := service.ContactService{}.GetCommunityPeopleNum(comunitys)
	global.ResponseOk(w, map[string]interface{}{
		"comunitys": comunitys,
		"group_map": GroupPeopleMap,
	}, "获取社群信息成功")
	// global.ResponseOkList(w,comunitys,len(comunitys))
}

func CreateCommunity(w http.ResponseWriter, req *http.Request) {
	var arg model.Community
	// 如果这个用的上,那么可以直接
	if err := utils.Bind(req, &arg); err != nil {
		log.Println(err)
		global.ResponseFail(w, err.Error())
		return
	}
	com, err := contactService.CreateCommunity(arg)
	if err != nil {
		global.ResponseFail(w, err.Error())
	} else {
		global.ResponseOk(w, com, "")
	}
	return
}

func JoinCommunity(w http.ResponseWriter, req *http.Request) {
	var arg request.ContactArg
	if err := utils.Bind(req, &arg); err != nil {
		global.ResponseFail(w, err.Error())
		return
	}
	err := contactService.JoinCommunity(arg.Userid, arg.Dstid)
	AddGroupId(arg.Userid, arg.Dstid)
	if err != nil {
		global.ResponseFail(w, err.Error())
	} else {
		global.ResponseOk(w, nil, "")
	}
}

func AddFriend(w http.ResponseWriter, req *http.Request) {
	// 定义一个参数结构体
	/*request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	passwd := request.PostForm.Get("passwd")
	*/
	var arg request.ContactArg
	if err := utils.Bind(req, &arg); err != nil {
		global.ResponseFail(w, err.Error())
		return
	}
	// 调用service
	err := contactService.AddFriend(arg.Userid, arg.Dstid)
	if err != nil {
		global.ResponseFail(w, err.Error())
	} else {
		global.ResponseOk(w, nil, "好友添加成功")
	}
}
