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

func AskCardListtHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	//别人请求小红花
	beego.Info("*********AskCardListtHandle Start**********")
	req_data := req.GetBody().GetAskcardlistReq()
	beego.Info(req_data)
	ret := int32(1)
	var AskCardlistNtf []*cspb.CSAskCardNtf
	AskCardlistNtf = append(AskCardlistNtf, makeAskcardlistmake(int32(100073), "大广", "春春", "红色"))
	AskCardlistNtf = append(AskCardlistNtf, makeAskcardlistmake(int32(100074), "小明", "春春", "红色"))
	AskCardlistNtf = append(AskCardlistNtf, makeAskcardlistmake(int32(100075), "小红", "春春", "红色"))
	res_data := new(cspb.CSAskCardListRes)
	*res_data = cspb.CSAskCardListRes{
		AskCardlistNtf: AskCardlistNtf,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		AskcardlistRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kAskCardListRes),
		res_pkg_body, res_list)
	return ret

}
