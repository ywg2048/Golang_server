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

func FriendmessageListHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********FriendmessageHandle Start**********")
	req_data := req.GetBody().GetFriendmessagelistReq()
	beego.Info(req_data)
	ret := int32(1)
	var Friendntf []*cspb.CSFriendNtf
	for i := range resmgr.FriendntftestData.GetItems() {
		Friendntf = append(Friendntf, makeFriendntf(resmgr.FriendntftestData.GetItems()[i].GetFriendId(), resmgr.FriendntftestData.GetItems()[i].GetStarId(), resmgr.FriendntftestData.GetItems()[i].GetName()))
	}
	var FriendmessagelistNtf []*cspb.CSFriendmessagelistNtf
	for j := range resmgr.FriendmessagelisttestData.GetItems() {
		FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(resmgr.FriendmessagelisttestData.GetItems()[j].GetMessageType(), resmgr.FriendmessagelisttestData.GetItems()[j].GetElementType(),
			Friendntf, resmgr.FriendmessagelisttestData.GetItems()[j].GetCardId(), resmgr.FriendmessagelisttestData.GetItems()[j].GetCardColor(), resmgr.FriendmessagelisttestData.GetItems()[j].GetElementNum(),
			resmgr.FriendmessagelisttestData.GetItems()[j].GetMessageId()))
	}
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
