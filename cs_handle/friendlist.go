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

func FriendlistHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********FriendlistHandle Start**********")
	req_data := req.GetBody().GetFriendlistReq()
	beego.Info(req_data)
	ret := int32(1)

	var FriendListNtf []*cspb.CSFriendListNtf
	FriendListNtf = append(FriendListNtf, makeFriendlist(int32(100073), "大广", "春春", int32(5), int32(1), int64(14282635486), int64(14282686592)))
	FriendListNtf = append(FriendListNtf, makeFriendlist(int32(100074), "小明", "春春", int32(6), int32(1), int64(14282635486), int64(14282686592)))
	FriendListNtf = append(FriendListNtf, makeFriendlist(int32(100075), "小红", "春春", int32(7), int32(1), int64(14282635486), int64(14282686592)))

	res_data := new(cspb.CSFriendListRes)
	*res_data = cspb.CSFriendListRes{
		FrendListNtf: FriendListNtf,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		FriendlistRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kFriendListRes),
		res_pkg_body, res_list)
	return ret

}
