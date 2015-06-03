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

func AskPowerListtHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	//别人请求小红花
	beego.Info("*********AskPowerListtHandle Start**********")
	req_data := req.GetBody().GetAskpowerlistReq()
	beego.Info(req_data)
	ret := int32(0)
	var AskPowerlistNtf []*cspb.CSAskPowerNtf
	AskPowerlistNtf = append(AskPowerlistNtf, makeAskPowerlistmake(int32(100073), "大广", "春春"))
	AskPowerlistNtf = append(AskPowerlistNtf, makeAskPowerlistmake(int32(100074), "小明", "春春"))
	AskPowerlistNtf = append(AskPowerlistNtf, makeAskPowerlistmake(int32(100075), "小红", "春春"))
	res_data := new(cspb.CSAskPowerListRes)
	*res_data = cspb.CSAskPowerListRes{
		AskPowerlistNtf: AskPowerlistNtf,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		AskpowerlistRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kAskPowerListRes),
		res_pkg_body, res_list)
	return ret

}
