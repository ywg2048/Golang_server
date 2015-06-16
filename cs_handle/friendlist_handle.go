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
	beego.Info("*********FriendmessageHandle Start**********")
	req_data := req.GetBody().GetFriendlistReq()
	beego.Info(req_data)
	ret := int32(1)
	var FriendListNtf []*cspb.CSFriendListNtf
	FriendListNtf = append(FriendListNtf, makefriendlist(int32(1), int32(100073), "小明", int32(8), "春春", int32(25000), int32(1), "演唱会", int32(10), int32(150), int32(1), int32(20)))
	FriendListNtf = append(FriendListNtf, makefriendlist(int32(2), int32(100074), "小红", int32(8), "春春", int32(30000), int32(1), "演唱会", int32(12), int32(150), int32(1), int32(20)))
	FriendListNtf = append(FriendListNtf, makefriendlist(int32(3), int32(100075), "小王", int32(8), "春春", int32(35000), int32(1), "演唱会", int32(13), int32(150), int32(1), int32(20)))
	FriendListNtf = append(FriendListNtf, makefriendlist(int32(4), int32(100076), "小张", int32(8), "春春", int32(28000), int32(1), "演唱会", int32(14), int32(150), int32(1), int32(20)))
	FriendListNtf = append(FriendListNtf, makefriendlist(int32(5), int32(100077), "小静", int32(8), "春春", int32(27000), int32(1), "演唱会", int32(15), int32(150), int32(1), int32(20)))

	res_data := new(cspb.CSFriendlistRes)
	*res_data = cspb.CSFriendlistRes{
		FriendListNtf: FriendListNtf,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		FriendlistRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kFriendlistRes),
		res_pkg_body, res_list)
	return ret

}
