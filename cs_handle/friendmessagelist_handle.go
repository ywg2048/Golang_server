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

func FriendmessageListHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********FriendmessageHandle Start**********")
	req_data := req.GetBody().GetFriendmessagelistReq()
	beego.Info(req_data)
	ret := int32(1)
	var Friendntf []*cspb.CSFriendNtf
	Friendntf = append(Friendntf, makeFriendntf(int32(10073), int32(8), "小明"))
	Friendntf = append(Friendntf, makeFriendntf(int32(10074), int32(8), "小红"))
	Friendntf = append(Friendntf, makeFriendntf(int32(10075), int32(8), "小张"))
	Friendntf = append(Friendntf, makeFriendntf(int32(10076), int32(8), "小王"))
	Friendntf = append(Friendntf, makeFriendntf(int32(10077), int32(8), "小静"))
	var FriendmessagelistNtf []*cspb.CSFriendmessagelistNtf
	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(1), int32(1), Friendntf, int32(0), "", int32(12), int32(1)))
	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(1), int32(2), Friendntf, int32(1), "红色", int32(12), int32(2)))
	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(1), Friendntf, int32(0), "", int32(12), int32(3)))
	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), Friendntf, int32(1), "红色", int32(12), int32(4)))
	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(3), Friendntf, int32(0), "", int32(0), int32(5)))
	uid := int32(2885377)
	res_data := new(cspb.CSFriendmessageListRes)
	*res_data = cspb.CSFriendmessageListRes{
		Uid:                  &uid,
		FriendmessagelistNtf: FriendmessagelistNtf,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		FriendmessagelistRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kFriendmessageListRes),
		res_pkg_body, res_list)
	return ret

}
