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

func LevelUpHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********LevelUpHandle Start**********")
	req_data := req.GetBody().GetLevelUpPageReq()
	beego.Info(req_data)

	ret := int32(1)
	var CardNtf []*cspb.CSCardNtf
	CardNtf = append(CardNtf, makeCardNtf(int32(1), int32(45)))
	CardNtf = append(CardNtf, makeCardNtf(int32(2), int32(45)))
	CardNtf = append(CardNtf, makeCardNtf(int32(3), int32(45)))
	CardNtf = append(CardNtf, makeCardNtf(int32(4), int32(45)))
	CardNtf = append(CardNtf, makeCardNtf(int32(5), int32(45)))
	CardNtf = append(CardNtf, makeCardNtf(int32(6), int32(45)))

	level := int32(10)
	affinity := int32(2500)
	playExperience := int32(50)
	fighting := int32(40000)
	dressId := int32(3)
	dress := "高级套装"
	exp := int32(20)

	res_data := new(cspb.CSLevelUpPageRes)
	*res_data = cspb.CSLevelUpPageRes{
		CardNtf:        CardNtf,
		Level:          &level,
		Affinity:       &affinity,
		PlayExperience: &playExperience,
		Fighting:       &fighting,
		DressId:        &dressId,
		Dress:          &dress,
		Exp:            &exp,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		LevelUpPageRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kLevelUpPageRes),
		res_pkg_body, res_list)
	return ret

}
