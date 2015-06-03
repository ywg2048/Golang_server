package cs_handle

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	models "tuojie.com/piggo/quickstart.git/models"
)
import cspb "protocol"

// import proto "code.google.com/p/goprotobuf/proto"

// import db "tuojie.com/piggo/quickstart.git/db/collection"
import "labix.org/v2/mgo/bson"

import db_session "tuojie.com/piggo/quickstart.git/db/session"

func GiveToSpecificFriendHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	//赠送特定好友卡片或者红花
	beego.Info("*********GiveToSpecificFriendHandle Start**********")
	req_data := req.GetBody().GetGivetospecificfriendReq()
	beego.Info(req_data)

	ret := int32(0)

	c := db_session.DB("zoo").C("player")
	var player models.Player
	err := c.Find(bson.M{"c_account": req_data.GetPlayuid()}).One(&player)
	if err != nil {
		beego.Error("出错了！")
	}
	//生成消息存在mysql中
	o := orm.NewOrm()
	var messages models.Messages
	messages.Fromuid = res_list.GetUid()
	messages.Touid = req_data.GetPlayuid()
	messages.Fromname = player.Name
	messages.Number = req_data.GetNumber()

	o.Insert(&messages)

	res_data := new(cspb.CSGiveToSpecificFriendRes)
	*res_data = cspb.CSGiveToSpecificFriendRes{}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		GivetospecificfriendRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kGiveToSpecificFriendRes),
		res_pkg_body, res_list)
	return ret

}
