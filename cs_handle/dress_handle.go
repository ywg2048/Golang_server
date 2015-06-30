package cs_handle

import (
	// "fmt"
	"github.com/astaxie/beego"
	// "time"
	// models "tuojie.com/piggo/quickstart.git/models"
)
import cspb "protocol"

// import proto "code.google.com/p/goprotobuf/proto"

// import db "tuojie.com/piggo/quickstart.git/db/collection"
// import "labix.org/v2/mgo/bson"

// import db_session "tuojie.com/piggo/quickstart.git/db/session"

import resmgr "tuojie.com/piggo/quickstart.git/res_mgr"

func DressHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********DressHandle Start**********")
	req_data := req.GetBody().GetDressReq()
	beego.Info(req_data)
	ret := int32(1)

	res_data := new(cspb.CSDressRes)
	starId := int32(1)
	var StarInfo []*cspb.CSStarInfo
	for i := range resmgr.DresstestData.GetItems() {
		StarInfo = append(StarInfo, makeStarInfo(resmgr.DresstestData.GetItems()[i].GetStarId(), resmgr.DresstestData.GetItems()[i].GetDressId()))
	}
	//正式代码
	// c := db_session.DB("zoo").C("player")
	// var player models.Player
	// c.Find(bson.M{"c_account":res_list.GetCAccount()}).One(&player)

	playerId := int32(res_list.GetUid())
	*res_data = cspb.CSDressRes{
		StarId:   &starId,
		PlayerId: &playerId,
		StarInfo: StarInfo,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		DressRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kDressRes),
		res_pkg_body, res_list)
	return ret

}
