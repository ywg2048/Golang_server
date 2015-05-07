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
	var message models.Messagecenter
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
	var messages []*cspb.CSMessageNtf
	for i := range message {
		messages[i].Id = &message[i].Id
		messages[i].Title = &message[i].Title
		messages[i].Content = &message[i].Content
		messages[i].Title2 = &message[i].Title2
		messages[i].Content2 = &message[i].Content2
		messages[i].IsActive = &message[i].IsActive
		messages[i].Time = &message[i].Time
	}
	beego.Info("******RES-----messages", messages)
	ret := int32(1)
	res_data := new(cspb.CSMessageCenterRes)
	beego.Debug("res_data", req_data)
	*res_data = cspb.CSMessageCenterRes{
		Ret:        proto.Int32(ret),
		MessageNtf: messages,
		// Id:       &message.Id,
		// Title:    &message.Title,
		// Content:  &message.Content,
		// Title2:   &message.Title2,
		// Content2: &message.Content2,
		// IsActive: &message.IsActive,
		// Time:     &message.Time,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		MessageCenterRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kMessageCenterRes),
		res_pkg_body, res_list)
	return ret

}
