package cs_handle

import (
	"github.com/astaxie/beego"
	models "tuojie.com/piggo/quickstart.git/models"
)
import cspb "protocol"
import proto "code.google.com/p/goprotobuf/proto"

// import db "tuojie.com/piggo/quickstart.git/db/collection"
import "labix.org/v2/mgo/bson"

import db_session "tuojie.com/piggo/quickstart.git/db/session"

func FinduserHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********FinduserHandle Start**********")
	req_data := req.GetBody().GetFinduserReq()
	beego.Info(req_data)

	ret := int32(0)
	IsExist := int32(0)  //查找uid存在与否 0代表没有找到 1代表找到
	IsFriend := int32(0) //查看是否是好友 1代表已经是好友 2代表已经申请好友但没通过 0代表不是好友
	c := db_session.DB("zoo").C("player")
	var player models.Player
	var player_isfriend models.Player
	err := c.Find(bson.M{"uid": req_data.GetUid()}).One(&player)
	errs := c.Find(bson.M{"FriendList.friendid": req_data.GetUid()}).One(&player_isfriend)
	if err != nil {
		if errs != nil {
			if player_isfriend.FriendList[0].IsActive == 1 {
				IsFriend = 1
			} else if player_isfriend.FriendList[0].IsActive == 0 {
				IsFriend = 2
			} else {
				IsFriend = 0
			}
		} else {
			IsFriend = 0
		}

		IsExist = 1
		ret = 1
	} else {
		//所有的用户中没有找到uid
		IsFriend = 0
		IsExist = 0
		ret = 1
	}

	res_data := new(cspb.CSFindUserRes)
	*res_data = cspb.CSFindUserRes{
		Ret:      proto.Int32(ret),
		IsExist:  &IsExist,
		IsFriend: &IsFriend,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		FinduserRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kFindUserRes),
		res_pkg_body, res_list)
	return ret

}
