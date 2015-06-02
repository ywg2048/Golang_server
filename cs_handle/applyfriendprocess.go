package cs_handle

import (
	"fmt"
	"github.com/astaxie/beego"
	"time"
	models "tuojie.com/piggo/quickstart.git/models"
)
import cspb "protocol"
import proto "code.google.com/p/goprotobuf/proto"

// import db "tuojie.com/piggo/quickstart.git/db/collection"
import "labix.org/v2/mgo/bson"

import db_session "tuojie.com/piggo/quickstart.git/db/session"

func ApplyFriendProcessHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********ApplyFriendProcessHandle Start**********")
	//处理客户端的申请列表的添加和拒绝逻辑
	req_data := req.GetBody().GetApplyfriendlistReq()
	beego.Info(req_data)
	c := db_session.DB("zoo").C("player")
	var player models.Player
	err := c.Find(bson.M{"c_account": res_list.GetCAccount()}).One(&player)
	beego.Debug("*********MoneyHandle result is %v err is %v********", player, err)

	ret := int32(0)
	IsAccept := int32(0)
	IsRefuse := int32(0)
	if req_data.GetAcceptOrrefuse() == "Accept" {
		//同意加好友

		//自己的好友表
		c.Upsert(bson.M{"c_account": res_list.GetCAccount()},
			bson.M{"$set": bson.M{"FriendList." + fmt.Sprint(req_data.GetApplyid()) + ".friendid": req_data.GetApplyid(), "FriendList." + fmt.Sprint(req_data.GetApplyid()) + ".isActive": 1, "FriendList." + fmt.Sprint(req_data.GetApplyid()) + ".accepttime": time.Now().Unix()}})

		c.Upsert(bson.M{"c_account": res_list.GetCAccount()},
			bson.M{"$set": bson.M{"ApplyFriendList." + fmt.Sprint(req_data.GetApplyid()) + ".applyuid": req_data.GetApplyid(), "ApplyFriendList." + fmt.Sprint(req_data.GetApplyid()) + ".isAccept": 1, "ApplyFriendList." + fmt.Sprint(req_data.GetApplyid()) + ".oprationtime": time.Now().Unix()}})

		//对方的表
		c.Upsert(bson.M{"uid": req_data.GetApplyid()},
			bson.M{"$set": bson.M{"FriendList." + fmt.Sprint(res_list.GetUid()) + ".friendid": res_list.GetUid(), "FriendList." + fmt.Sprint(res_list.GetUid()) + ".isActive": 1, "FriendList." + fmt.Sprint(res_list.GetUid()) + ".accepttime": time.Now().Unix()}})
		IsAccept = int32(1)
	} else {
		//拒绝加好友

		//自己的表
		c.Upsert(bson.M{"c_account": res_list.GetCAccount()},
			bson.M{"$set": bson.M{"ApplyFriendList." + fmt.Sprint(req_data.GetApplyid()) + ".applyuid": req_data.GetApplyid(), "ApplyFriendList." + fmt.Sprint(req_data.GetApplyid()) + ".isrefuse": 0, "ApplyFriendList." + fmt.Sprint(req_data.GetApplyid()) + ".oprationtime": time.Now().Unix()}})
		//对方的表
		c.Upsert(bson.M{"uid": req_data.GetApplyid()},
			bson.M{"$set": bson.M{"FriendList." + fmt.Sprint(res_list.GetUid()) + ".friendid": res_list.GetUid(), "FriendList." + fmt.Sprint(res_list.GetUid()) + ".isActive": 0, "FriendList." + fmt.Sprint(res_list.GetUid()) + ".accepttime": time.Now().Unix()}})
		IsRefuse = int32(1)
	}
	res_data := new(cspb.CSApplyFriendListRes)
	*res_data = cspb.CSApplyFriendListRes{
		Ret:      proto.Int32(ret),
		IsAccept: &IsAccept,
		IsRefuse: &IsRefuse,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		ApplyfriendlistRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kMoneyRes),
		res_pkg_body, res_list)
	return ret

}
