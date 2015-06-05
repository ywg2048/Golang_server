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
	whiteCard := int32(15)
	redCard := int32(20)
	yellowCard := int32(5)
	level := int32(10)
	satisfaction := int32(2500)
	fighting := int32(40000)
	dress := "演唱会"
	exp := int32(20)
	solutionPool := int32(150)
	res_data := new(cspb.CSLevelUpPageRes)
	*res_data = cspb.CSLevelUpPageRes{
		WhiteCard:    &whiteCard,
		RedCard:      &redCard,
		YellowCard:   &yellowCard,
		Level:        &level,
		Satisfaction: &satisfaction,
		Fighting:     &fighting,
		Dress:        &dress,
		Exp:          &exp,
		SolutionPool: &solutionPool,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		LevelUpPageRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kLevelUpPageRes),
		res_pkg_body, res_list)
	return ret

}
