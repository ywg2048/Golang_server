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

func DressHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********ZooHandle Start**********")
	req_data := req.GetBody().GetDressReq()
	beego.Info(req_data)
	ret := int32(1)

	res_data := new(cspb.CSDressRes)
	starId := int32(1)
	var StarInfo []*cspb.CSStarInfo
	StarInfo = append(StarInfo, makeStarInfo(int32(9), int32(1)))
	StarInfo = append(StarInfo, makeStarInfo(int32(8), int32(2)))
	StarInfo = append(StarInfo, makeStarInfo(int32(11), int32(3)))
	StarInfo = append(StarInfo, makeStarInfo(int32(12), int32(4)))

	playerId := int32(res_list.GetUid())
	*res_data = cspb.CSDressRes{
		StarId:   &starId,
		PlayerId: &playerId,
		StarInfo: StarInfo,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		DressRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kDressRes),
		res_pkg_body, res_list)
	return ret

}
