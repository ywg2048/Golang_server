package cs_handle

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	models "tuojie.com/piggo/quickstart.git/models"
)

import cspb "protocol"
import proto "code.google.com/p/goprotobuf/proto"

import "time"

func init() {
	orm.RegisterDataBase("Monsters", "mysql", "root:@/Monsters?charset=utf8")

}
func ServerTimeHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {

	req_data := req.GetBody().GetServerTimeReq()
	beego.Info(req_data)
	beego.Info(res_list)
	beego.Info(res_list.GetUid())
	beego.Debug("******serverTimeHandle")
	ret := int32(0)
	now_time := time.Now().Unix()
	beego.Debug("server time :%d", now_time)
	//填充ServerTimeRes回包
	var messages []models.Messages

	var cond *orm.Condition
	cond = orm.NewCondition()

	cond = cond.And("Id__gte", 1)
	cond = cond.And("Touid__contains", int32(res_list.GetUid()))
	cond = cond.And("IsFinish__contains", int32(0))
	beego.Info(cond)
	var qs orm.QuerySeter
	qs = orm.NewOrm().QueryTable("messages").SetCond(cond)

	cnt, err := qs.All(&messages)
	beego.Info(cnt, err)
	var messagenTipsntf []*cspb.CSmessageTipsntf
	messagenTipsntf = append(messagenTipsntf, makemessageTipsntf(int32(1), int32(len(messages))))
	messagenTipsntf = append(messagenTipsntf, makemessageTipsntf(int32(2), int32(3)))
	//发完之后就讲mysql中的Isfinish变为1
	// var message models.Messages
	// o := orm.NewOrm()
	// for i := range messages {
	// 	beego.Info(i)
	// 	message = models.Messages{Touid: int32(res_list.GetUid()), IsFinish: int32(0)}
	// 	o.Read(&message, "Touid", "IsFinish")
	// 	beego.Info(message)
	// 	message.IsFinish = int32(1)
	// 	id, err := o.Update(&message, "IsFinish")
	// 	beego.Info(err, id)
	// }
	res_data := new(cspb.CSServerTimeRes)
	*res_data = cspb.CSServerTimeRes{
		Ret:             proto.Int32(ret),
		ServerTime:      proto.Int64(now_time),
		MessagenTipsntf: messagenTipsntf,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		ServerTimeRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kServerTimeRes),
		res_pkg_body, res_list)
	beego.Debug("serverTimeHandle******")
	return ret
}
