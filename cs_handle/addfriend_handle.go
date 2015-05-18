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

func AddFriendHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********AddFriendHandle Start**********")
	req_data := req.GetBody().GetAddfriendReq()
	beego.Info(req_data)

	ret := int32(0)
	IsAdd := int32(0) //是否申请加好友成功，成功1 失败0
	IsAddself := int32(0)
	IsAddop := int(0)
	c := db_session.DB("zoo").C("player")
	//自己的表
	var player models.Player
	err := c.Find(bson.M{"c_account": res_list.GetCAccount()}).One(&player)
	if err != nil {
		var friendlist models.FriendListData
		friendlist.Friendid = req_data.GetUid()
		friendlist.IsActive = int32(0)
		friendlist.Accepttime = int64(0)
		_, err := c.Upsert(bson.M{"c_account": res_list.GetCAccount()},
			bson.M{"$set": bson.M{"FriendList." + fmt.Sprint(req_data.GetUid()): friendlist}})
		if err != nil {
			IsAddself = 1
		}
	}
	//对方的表
	var friendinfo models.Player
	errs := c.Find(bson.M{"uid": req_data.GetUid()}).One(&friendinfo)
	if errs != nil {
		var applyfriendlist models.ApplyFriendListData
		applyfriendlist.Applyuid = req_data.GetUid()
		applyfriendlist.IsAccept = int32(0)
		applyfriendlist.Isrefuse = int32(0)
		applyfriendlist.Applytime = time.Now().Unix()
		applyfriendlist.Oprationtime = int64(0)
		_, err := c.Upsert(bson.M{"uid": req_data.GetUid()},
			bson.M{"$set": bson.M{"ApplyFriendList." + fmt.Sprint(res_list.GetCAccount()): applyfriendlist}})
		if err != nil {
			IsAddop = 1
		}
	}
	if IsAddself == 1 && IsAddop == 1 {
		IsAdd = 1
		ret = 1
	} else {
		IsAdd = 0
		ret = 1
	}

	uid := req_data.GetUid()
	res_data := new(cspb.CSAddFriendRes)
	*res_data = cspb.CSAddFriendRes{
		Ret:   proto.Int32(ret),
		Uid:   &uid,
		IsAdd: &IsAdd,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		AddfriendRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kAddFriendRes),
		res_pkg_body, res_list)
	return ret

}
