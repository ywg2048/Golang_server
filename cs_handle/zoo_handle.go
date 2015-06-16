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

func ZooHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********ZooHandle Start**********")
	req_data := req.GetBody().GetZooReq()
	beego.Info(req_data)
	ret := int32(1)

	var MyFriendNtf []*cspb.CSMyFriendNtf
	MyFriendNtf = append(MyFriendNtf, makeMyFriendNtf(int32(10073), "大广", int32(8)))
	MyFriendNtf = append(MyFriendNtf, makeMyFriendNtf(int32(10074), "小红", int32(8)))
	MyFriendNtf = append(MyFriendNtf, makeMyFriendNtf(int32(10075), "小张", int32(8)))

	var AnimalNtf []*cspb.CSAnimalNtf
	AnimalNtf = append(AnimalNtf, makeAnimalNtf(int32(1), int32(100), int32(10), MyFriendNtf))
	AnimalNtf = append(AnimalNtf, makeAnimalNtf(int32(2), int32(100), int32(10), MyFriendNtf))
	AnimalNtf = append(AnimalNtf, makeAnimalNtf(int32(3), int32(100), int32(10), MyFriendNtf))

	res_data := new(cspb.CSZooRes)

	*res_data = cspb.CSZooRes{
		AnimalNtf: AnimalNtf,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		ZooRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kZooRes),
		res_pkg_body, res_list)
	return ret

}
