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

func FriendStageHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********ZooHandle Start**********")
	req_data := req.GetBody().GetFriendstageReq()
	beego.Info(req_data)
	ret := int32(1)
	var Friendstage []*cspb.CSFriendStageNtf
	Friendstage = append(Friendstage, makeFriendStagentf(int32(100073), "大广", int32(10)))
	Friendstage = append(Friendstage, makeFriendStagentf(int32(100074), "小敏", int32(12)))
	Friendstage = append(Friendstage, makeFriendStagentf(int32(100075), "小米", int32(15)))
	Friendstage = append(Friendstage, makeFriendStagentf(int32(100076), "阿牛", int32(23)))
	Friendstage = append(Friendstage, makeFriendStagentf(int32(2885377), "朱林", int32(30)))
	Friendstage = append(Friendstage, makeFriendStagentf(int32(100077), "小王", int32(35)))

	res_data := new(cspb.CSFriendStageRes)
	*res_data = cspb.CSFriendStageRes{
		Friendstage: Friendstage,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		FriendstageRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kFriendStageRes),
		res_pkg_body, res_list)
	return ret

}
