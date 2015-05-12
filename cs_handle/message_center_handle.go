package cs_handle

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	models "tuojie.com/piggo/quickstart.git/models"
)
import cspb "protocol"
import proto "code.google.com/p/goprotobuf/proto"

func init() {

	orm.RegisterDataBase("default", "mysql", "root:@/Monsters?charset=utf8")

}

func MessageCenterHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("******messageCenterHandle Start*****")
	req_data := req.GetBody().GetMessageCenterReq()
	beego.Debug("********req_data****************", req_data)
	var message []models.Messagecenter
	var cond *orm.Condition
	cond = orm.NewCondition()

	cond = cond.And("Id__gte", 1)
	cond = cond.And("IsActive__contains", 1)
	var qs orm.QuerySeter
	qs = orm.NewOrm().QueryTable("messagecenter").SetCond(cond)
	cnt, err := qs.All(&message)
	if err != nil {
		beego.Debug("查询数据库失败")
	}
	beego.Debug(message, cnt, err)
	var res_messages []*cspb.CSMessageNtf
	for i := range message {

		res_messages = append(res_messages, makeMessage(message[i].Id, message[i].Title, message[i].Content, message[i].IsActive, message[i].Time))

	}
	ret := int32(1)
	res_data := new(cspb.CSMessageCenterRes)
	beego.Debug("res_data", req_data)
	*res_data = cspb.CSMessageCenterRes{
		Ret:        proto.Int32(ret),
		MessageNtf: res_messages,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		MessageCenterRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kMessageCenterRes),
		res_pkg_body, res_list)
	return ret

}
