package cs_handle

import (
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "time"
	models "tuojie.com/piggo/quickstart.git/models"
)
import cspb "protocol"

// import proto "code.google.com/p/goprotobuf/proto"

// import db "tuojie.com/piggo/quickstart.git/db/collection"
import "labix.org/v2/mgo/bson"

import db_session "tuojie.com/piggo/quickstart.git/db/session"

// import resmgr "tuojie.com/piggo/quickstart.git/res_mgr"

func RemoveMessageHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********RemoveMessageHandle Start**********")
	req_data := req.GetBody().GetRemovemessageReq()
	beego.Info(req_data)
	ret := int32(1)
	isRemove := int32(0)

	c := db_session.DB("zoo").C("player")
	var player models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player)

	var message models.Messages
	o := orm.NewOrm()
	for i := range req_data.GetMysqlIdNtf() {
		mysqlid := req_data.GetMysqlIdNtf()[i].GetMysqlId()
		message = models.Messages{Id: mysqlid}
		o.Read(&message)
		beego.Info(message)
		message.IsFinish = int32(1)
		id, err := o.Update(&message, "IsFinish")
		beego.Info(err, id)
	}

	res_data := new(cspb.CSRemoveMessageRes)
	*res_data = cspb.CSRemoveMessageRes{
		IsRemove: isRemove,
	}
	beego.Info(res_data)
	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		RemovemessageRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kRemoveMessageRes),
		res_pkg_body, res_list)
	return ret

}
