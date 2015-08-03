package cs_handle

import (
	"github.com/astaxie/beego"
	models "tuojie.com/piggo/quickstart.git/models"
)
import cspb "protocol"

// import proto "code.google.com/p/goprotobuf/proto"
// import db "tuojie.com/piggo/quickstart.git/db/collection"
import "labix.org/v2/mgo/bson"
import db_session "tuojie.com/piggo/quickstart.git/db/session"

// import resmgr "tuojie.com/piggo/quickstart.git/res_mgr"

func petReportHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	req_data := req.GetBody().GetReportPetReq()
	pet_id := req_data.GetPetId()
	ret := int32(1)
	Ret := int32(0)
	c := db_session.DB("zoo").C("player")
	var player models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player)
	_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
		bson.M{"$set": bson.M{"starid": pet_id}})
	if err == nil {
		Ret = int32(1)
		beego.Info("明星更换成功！")
	} else {
		beego.Error("明星更换失败！")
	}

	//更换小伙伴
	// if req_data.GetType == int32(1) {
	// 	_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
	// 		bson.M{"$set": bson.M{"starid": pet_id}})
	// 	if err == nil {
	// 		Ret = int32(1)
	// 		beego.Info("明星更换成功！")
	// 	} else {
	// 		beego.Error("明星更换失败！")
	// 	}
	// } else {
	// //解锁小伙伴
	// _, err := c.Upsert(bson.M{"uid": player.Uid},
	// 	bson.M{"$push": bson.M{"star": bson.M{"starid": pet_id,  "level": int32(1), "current_exp": int32(0), "dress": int32(1),
	// 		"dressname": "初级套装", "fighting": int32(16500), "satisfaction": int32(50), "fight_exp": int32(0), "is_active": int32(1)}}})

	// if err != nil {
	// 	beego.Error("解锁小伙伴失败", err)
	// }
	// }
	res_data := new(cspb.CSReportPetRes)
	*res_data = cspb.CSReportPetRes{
		Ret: &Ret,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		ReportPetRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kReportPetRes),
		res_pkg_body, res_list)
	return ret
}
