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

func AchievementupdateHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********AchievementupdateHandle Start**********")
	req_data := req.GetBody().GetAchievementupdateReq()
	beego.Info(req_data)
	ret := int32(1)
	IsUpdate := int32(0)
	c := db_session.DB("zoo").C("player")
	var player models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player)

	for i := range player.Achievement {
		if player.Achievement[i].AchievementId == req_data.GetAchievementid() {
			if player.Achievement[i].StarLevel == req_data.GetStarLevel() {
				//先检查星星等级是否合法

				_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
					bson.M{"$set": bson.M{"achievement." + fmt.Sprint(i) + ".process": req_data.GetProcess()}})
				if err == nil {
					beego.Info("成就进度更新成功！")
					IsUpdate = int32(1)
				} else {
					beego.Error("成就进度更新失败！")
				}
			} else {
				beego.Error("星星的等级不合法！")
			}
		}
	}

	res_data := new(cspb.CSAchievementUpdateRes)
	*res_data = cspb.CSAchievementUpdateRes{
		IsUpdate: &IsUpdate,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		AchievementupdateRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kAchievementRes),
		res_pkg_body, res_list)
	return ret

}
