package cs_handle

import (
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
	models "tuojie.com/piggo/quickstart.git/models"
)
import cspb "protocol"

// import proto "code.google.com/p/goprotobuf/proto"

// import db "tuojie.com/piggo/quickstart.git/db/collection"
import "labix.org/v2/mgo/bson"

import db_session "tuojie.com/piggo/quickstart.git/db/session"

// import resmgr "tuojie.com/piggo/quickstart.git/res_mgr"

func ApplyFlowerHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********ApplyFlowerHandle Start**********")
	req_data := req.GetBody().GetApplyflowerReq()
	beego.Info(req_data)
	ret := int32(1)

	c := db_session.DB("zoo").C("player")
	var player models.Player
	for i := range req_data.GetApplyFlowerList() {

		//查找用户是否存在
		err := c.Find(bson.M{"uid": req_data.GetApplyFlowerList()[i].GetFriendId()}).One(&player)
		if err != nil {
			beego.Error("找不到用户！")
		} else {
			beego.Info("找到用户！")
			//消息提示，存在Mysql中
			var players models.Player
			c.Find(bson.M{"c_account": res_list.GetCAccount()}).One(&players)

			o := orm.NewOrm()
			var messages models.Messages
			messages.Fromuid = int32(res_list.GetUid())
			messages.Fromname = players.Name
			messages.FromStarId = players.StarId
			messages.Time = time.Now().Unix()
			messages.IsFinish = int32(0)
			messages.Messagetype = int32(1)
			messages.ElementType = int32(1)
			messages.Number = int32(1)

			messages.Touid = req_data.GetApplyFlowerList()[i].GetFriendId()

			o.Insert(&messages)
		}
	}
	res_data := new(cspb.CSApplyFlowerRes)
	*res_data = cspb.CSApplyFlowerRes{}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		ApplyflowerRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kApplyFlowerRes),
		res_pkg_body, res_list)
	return ret

}
