package cs_handle

import (
	// "fmt"
	"github.com/astaxie/beego"
	"time"
	models "tuojie.com/piggo/quickstart.git/models"
)
import cspb "protocol"

// import proto "code.google.com/p/goprotobuf/proto"

// import db "tuojie.com/piggo/quickstart.git/db/collection"
import "labix.org/v2/mgo/bson"

import db_session "tuojie.com/piggo/quickstart.git/db/session"

// import resmgr "tuojie.com/piggo/quickstart.git/res_mgr"

func ResourceHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********DressHandle Start**********")
	req_data := req.GetBody().GetResourceReq()
	beego.Info(req_data)
	ret := int32(1)
	c := db_session.DB("zoo").C("player")
	var player models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player)
	if req_data.GetDiamond() == player.Diamond {
		_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
			bson.M{"$inc": bson.M{"diamond": req_data.GetDiamondChange()}})
		if err == nil {
			beego.Info("钻石数量变更成功！")
		} else {
			beego.Error("钻石数量变更失败！")
		}
	} else {
		beego.Error("客户端钻石和服务器不一致")
	}
	if req_data.GetGoldcoin() == player.Gold {
		_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
			bson.M{"$inc": bson.M{"gold": req_data.GetGoldcoinChange()}})
		if err == nil {
			beego.Info("金币数量变更成功！")
		} else {
			beego.Error("金币数量变更失败！")
		}
	} else {
		beego.Error("客户端金币和服务器不一致")
	}
	if req_data.GetFlower() == player.Flower {
		_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
			bson.M{"$inc": bson.M{"flower": req_data.GetFlowerChange()}})
		if err == nil {
			beego.Info("小红花数量变更成功！")
		} else {
			beego.Error("小红花数量变更失败！")
		}
	} else {
		beego.Error("客户端小红花和服务器不一致")
	}
	var player_after models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player_after)
	Diamond := player_after.Diamond
	Gold := player_after.Gold
	Flower := player_after.Flower
	beego.Info(player_after)
	res_data := new(cspb.CSResourceRes)
	Servertime := time.Now().Unix()
	*res_data = cspb.CSResourceRes{
		Diamond:    &Diamond,
		Goldcoin:   &Gold,
		Flower:     &Flower,
		Servertime: &Servertime,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		ResourceRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kResourceRes),
		res_pkg_body, res_list)
	return ret

}