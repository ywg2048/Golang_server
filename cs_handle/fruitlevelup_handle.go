package cs_handle

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	models "tuojie.com/piggo/quickstart.git/models"
)

import cspb "protocol"
import proto "code.google.com/p/goprotobuf/proto"
import "labix.org/v2/mgo/bson"

import db_session "tuojie.com/piggo/quickstart.git/db/session"

// import "time"

func init() {
	orm.RegisterDataBase("Monsters", "mysql", "root:@/Monsters?charset=utf8")

}
func FruitLevelUpHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("******FruitLevelUpHandle Start******")
	req_data := req.GetBody().GetFruitlevelupReq()
	beego.Info(req_data)
	ret := int32(1)
	issuccess := int32(0)
	c := db_session.DB("zoo").C("player")
	var player models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player)
	if req_data.GetCurrentFruitlevel() != int32(-1) {
		_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
			bson.M{"$inc": bson.M{"fruit_level": 1, "gold": -req_data.GetRequiredGold()}})
		if err == nil {
			beego.Info("水果升级成功！")
			issuccess = int32(1)
		} else {
			beego.Error("水果升级失败！", err)
		}
	}
	var player_return models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player_return)
	res_data := new(cspb.CSFruitLevelUpRes)
	*res_data = cspb.CSFruitLevelUpRes{
		CurrentFruitlevel: proto.Int32(player_return.FruitLevel),

		Issuccess:   &issuccess,
		CurrentGold: proto.Int32(player_return.Gold),
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		FruitlevelupRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kFruitLevelUpRes),
		res_pkg_body, res_list)

	return ret
}
