package cs_handle

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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
	ret := int32(1)
	uid := req_data.GetUid()
	friendId := req_data.GetFriendId()
	IsAdd := int32(1)
	Tips := "发送成功！"

	c := db_session.DB("zoo").C("player")
	var player models.Player
	//查找用户是否存在
	//是否是
	err := c.Find(bson.M{"uid": friendId}).One(&player)
	if err == nil {
		//是否存在
		if uid != friendId {
			//不为自己
			var myself models.Player
			c.Find(bson.M{"uid": uid}).One(&myself)
			for i := range myself.FriendList {
				if myself.FriendList[i].Friendid == friendId {
					//已经申请或者已经是好友
					IsAdd = int32(1)
					Tips = "你们已经申请或者已经是好友！"
				}
			}

		} else {
			IsAdd = int32(0)
			Tips = "请不要添加自己为好友！"
		}
	} else {
		IsAdd = int32(0)
		Tips = "用户不存在！"
	}

	if IsAdd == int32(1) {
		beego.Info("操作1")
		//自己的表
		//检测是否已经申请过了
		var myself models.Player
		c.Find(bson.M{"uid": uid}).One(&myself)

		isadded := int32(0)
		for i := range myself.FriendList {
			if myself.FriendList[i].Friendid == friendId {
				//已经申请或者已经是好友
				isadded = int32(1)
				Tips = "你们已经申请或者已经是好友！"
			}
		}
		if isadded == int32(0) {
			c.Find(bson.M{"uid": uid}).One(&player)

			i := len(player.FriendList)
			_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
				bson.M{"$set": bson.M{"FriendList." + fmt.Sprint(i) + ".friendid": friendId, "FriendList." + fmt.Sprint(i) + ".isActive": int32(0), "FriendList." + fmt.Sprint(i) + ".accepttime": int64(0)}})
			beego.Info(err)
			//朋友的申请列表

			var friend models.Player
			c.Find(bson.M{"uid": friendId}).One(&friend)
			j := len(friend.ApplyFriendList)

			c.Upsert(bson.M{"uid": friendId},
				bson.M{"$set": bson.M{"ApplyFriendList." + fmt.Sprint(j) + ".applyuid": uid, "ApplyFriendList." + fmt.Sprint(j) + ".isAccept": int32(0), "ApplyFriendList." + fmt.Sprint(j) + ".isrefuse": int32(0), "ApplyFriendList." + fmt.Sprint(j) + ".applytime": time.Now().Unix(), "ApplyFriendList." + fmt.Sprint(j) + ".oprationtime": int64(0)}})
			k := len(friend.FriendList)
			_, errs := c.Upsert(bson.M{"uid": friendId},
				bson.M{"$set": bson.M{"FriendList." + fmt.Sprint(k) + ".friendid": int32(req_data.GetUid()), "FriendList." + fmt.Sprint(k) + ".isActive": int32(0), "FriendList." + fmt.Sprint(k) + ".accepttime": int64(0)}})
			beego.Info(errs)
			//消息通知mysql表
			var players models.Player
			err_self := c.Find(bson.M{"uid": uid}).One(&players)
			beego.Info(err_self)
			o := orm.NewOrm()
			var messages models.Messages
			messages.Fromuid = int32(res_list.GetUid())
			messages.Fromname = players.Name
			messages.FromStarId = players.StarId
			messages.Time = time.Now().Unix()
			messages.IsFinish = int32(0)
			messages.Messagetype = int32(0)
			messages.ElementType = int32(3)

			messages.Touid = friendId
			messages.Tag = int32(5)
			id, err := o.Insert(&messages)
			beego.Info(id, err)

			if err == nil {
				beego.Info("插入成功！", id)
			}
		}
	}
	beego.Info(err)
	res_data := new(cspb.CSAddFriendRes)
	*res_data = cspb.CSAddFriendRes{
		Ret:      proto.Int32(ret),
		Uid:      &uid,
		FriendId: &friendId,
		IsAdd:    &IsAdd,
		Tips:     &Tips,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		AddfriendRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kAddFriendRes),
		res_pkg_body, res_list)
	return ret

}
