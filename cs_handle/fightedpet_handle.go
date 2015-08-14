package cs_handle

import (
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	models "tuojie.com/piggo/quickstart.git/models"
)

import cspb "protocol"

import proto "code.google.com/p/goprotobuf/proto"
import "labix.org/v2/mgo/bson"

import db_session "tuojie.com/piggo/quickstart.git/db/session"

// import "time"

func FightedPetHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("******FruitLevelUpHandle Start******")
	req_data := req.GetBody().GetFightedpetReq()
	beego.Info(req_data)
	ret := int32(1)
	c := db_session.DB("zoo").C("player")
	var player models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player)

	if req_data.GetCurrentFightedPetId() != int32(-1) {
		_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
			bson.M{"$set": bson.M{"starid": req_data.GetCurrentFightedPetId()}})
		if err != nil {
			beego.Error("小伙伴更换失败！")
		} else {
			beego.Info("小伙伴更换成功！")
		}

		//mysql排名表

		o := orm.NewOrm()
		var ranking models.Ranking

		ranking = models.Ranking{Uid: int32(res_list.GetUid())}
		o.Read(&ranking, "uid")
		beego.Info(ranking)
		ranking.StarId = req_data.GetCurrentFightedPetId()
		id, errs := o.Update(&ranking, "starid")
		if errs == nil {
			beego.Info("mysql小伙伴变更成功", id)
		} else {
			beego.Error("mysql小伙伴变更失败", err)
		}
	}
	var player_return models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player_return)

	res_data := new(cspb.CSFightedPetRes)
	*res_data = cspb.CSFightedPetRes{
		CurrentFightedPetId: proto.Int32(player_return.StarId),
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		FightedpetRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kFightedPetRes),
		res_pkg_body, res_list)

	return ret
}
