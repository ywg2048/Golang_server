package cs_handle

import (
	// "fmt"
	"github.com/astaxie/beego"
	// "time"
	models "tuojie.com/piggo/quickstart.git/models"
)

import cspb "protocol"

// import proto "code.google.com/p/goprotobuf/proto"

// import db "tuojie.com/piggo/quickstart.git/db/collection"
import "labix.org/v2/mgo/bson"

import db_session "tuojie.com/piggo/quickstart.git/db/session"

func MainInfoHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********MainInfoHandle Start**********")
	req_data := req.GetBody().GetMaininfoReq()
	beego.Info(req_data)

	ret := int32(1)
	c := db_session.DB("zoo").C("player")
	var player models.Player
	err := c.Find(bson.M{"c_account": res_list.GetCAccount()}).One(&player)
	if err != nil {
		beego.Error("出错了！")
	}

	// Flower := playerinfo.Flower
	// Gold := playerinfo.Gold
	// Diamond := playerinfo.Diamond
	// Medal := playerinfo.Medal
	Flower := int32(10)
	Gold := int32(100000)
	Diamond := int32(1000)
	Medal := int32(250)
	Ranking := int32(4)
	Star := "春春"
	Starlevel := int32(8)
	Starsolution := int32(8000)

	res_data := new(cspb.CSMainInfoRes)
	*res_data = cspb.CSMainInfoRes{
		Flower:       &Flower,
		Gold:         &Gold,
		Diamond:      &Diamond,
		Medal:        &Medal,
		Ranking:      &Ranking,
		Star:         &Star,
		Starlevel:    &Starlevel,
		Starsolution: &Starsolution,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		MaininfoRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kMainInfoRes),
		res_pkg_body, res_list)
	return ret

}
