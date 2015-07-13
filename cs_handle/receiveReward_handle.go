package cs_handle

import (
	"fmt"
	"github.com/astaxie/beego"
	// "time"
	models "tuojie.com/piggo/quickstart.git/models"
)
import cspb "protocol"

// import proto "code.google.com/p/goprotobuf/proto"

// import db "tuojie.com/piggo/quickstart.git/db/collection"
import "labix.org/v2/mgo/bson"

import db_session "tuojie.com/piggo/quickstart.git/db/session"

// import resmgr "tuojie.com/piggo/quickstart.git/res_mgr"

func ReceiveRewardHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********ReceiveRewardHandle Start**********")
	req_data := req.GetBody().GetReceiverewardReq()
	beego.Info(req_data)
	ret := int32(1)
	IsReceive := int32(0)
	c := db_session.DB("zoo").C("player")
	var player models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player)
	for i := range player.Achievement {
		if player.Achievement[i].AchievementId == req_data.GetAchievementid() {
			if player.Achievement[i].StarLevel <= req_data.GetStarLevel() {
				//只有没有领取的时候才能领取
				_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
					bson.M{"$set": bson.M{"achievement." + fmt.Sprint(i) + ".achievementid": req_data.GetAchievementid(), "achievement." + fmt.Sprint(i) + ".isreceive": int32(1)}})
				_, errs := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
					bson.M{"$inc": bson.M{"experience_pool": req_data.GetExp(), "diamond": req_data.GetDiamond(), "achievement." + fmt.Sprint(i) + ".starlevel": int32(1)}})
				if errs == nil && err == nil {
					beego.Info("领取奖励成功！")
					IsReceive = int32(1)
				} else {
					beego.Error("领取奖励失败！")
				}
			} else {
				beego.Error("该奖励已经领取！")
			}
		}
	}
	beego.Info(IsReceive)
	res_data := new(cspb.CSReceiveRewardRes)
	*res_data = cspb.CSReceiveRewardRes{
		IsReceive: &IsReceive,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		ReceiverewardRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kReceiveRewardRes),
		res_pkg_body, res_list)
	return ret

}
