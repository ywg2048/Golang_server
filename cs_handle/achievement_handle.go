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

func AchievementHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********AchievementHandle Start**********")
	req_data := req.GetBody().GetAchievementReq()
	beego.Info(req_data)
	ret := int32(1)
	var AchievementNtf []*cspb.CSAchievementNtf
	for i := range resmgr.AchievementtestData.GetItems() {
		AchievementNtf = append(AchievementNtf, makeAchievementNtf(resmgr.AchievementtestData.GetItems()[i].GetAchievementid(), resmgr.AchievementtestData.GetItems()[i].GetStatus()))
	}

	res_data := new(cspb.CSAchievementRes)
	*res_data = cspb.CSAchievementRes{
		AchievementNtf: AchievementNtf,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		AchievementRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kAchievementRes),
		res_pkg_body, res_list)
	return ret

}