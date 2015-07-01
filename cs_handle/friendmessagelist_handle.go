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
import resmgr "tuojie.com/piggo/quickstart.git/res_mgr"

func FriendmessageListHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********FriendmessageHandle Start**********")
	req_data := req.GetBody().GetFriendmessagelistReq()
	beego.Info(req_data)
	ret := int32(1)
	//测试代码
	var Friendntf []*cspb.CSFriendNtf
	for i := range resmgr.FriendntftestData.GetItems() {
		Friendntf = append(Friendntf, makeFriendntf(resmgr.FriendntftestData.GetItems()[i].GetFriendId(), resmgr.FriendntftestData.GetItems()[i].GetStarId(), resmgr.FriendntftestData.GetItems()[i].GetName()))
	}
	var FriendmessagelistNtf []*cspb.CSFriendmessagelistNtf
	for j := range resmgr.FriendmessagelisttestData.GetItems() {
		FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(resmgr.FriendmessagelisttestData.GetItems()[j].GetMessageType(), resmgr.FriendmessagelisttestData.GetItems()[j].GetElementType(),
			Friendntf, resmgr.FriendmessagelisttestData.GetItems()[j].GetCardId(), resmgr.FriendmessagelisttestData.GetItems()[j].GetCardColor(), resmgr.FriendmessagelisttestData.GetItems()[j].GetElementNum(),
			resmgr.FriendmessagelisttestData.GetItems()[j].GetMessageId()))
	}
	uid := int32(2885377)
	//正式代码
	c := db_session.DB("zoo").C("player")
	var player models.Player
	err := c.Find(bson.M{"c_account": res_list.GetCAccount()}).One(&player)
	if err != nil {
		beego.Error(err)
	}

	var messages []models.Messages
	var cond *orm.Condition
	cond = orm.NewCondition()

	// cond = cond.And("Uid__contains", int32(res_list.GetUid()))
	cond = cond.And("IsFinish__contains", 0)
	cond = cond.And("Touid__contains", int32(res_list.GetUid()))
	var qs orm.QuerySeter
	qs = orm.NewOrm().QueryTable("messages").Limit(20).SetCond(cond)
	cnt, err := qs.All(&messages)
	beego.Info(cnt, err)
	//从sql中取出消息
	for k := range messages {

		Friendntf = append(Friendntf, makeFriendntf(messages[k].Fromuid, messages[k].FromStarId, messages[k].Fromname))
		FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(messages[k].Messagetype, messages[k].ElementType, Friendntf, messages[k].CardId, "红色", messages[k].Number, messages[k].Id))
	}

	//发完之后就讲mysql中的Isfinish变为1
	var message models.Messages
	o := orm.NewOrm()
	for i := range messages {
		beego.Info(i)
		message = models.Messages{Touid: int32(res_list.GetUid()), IsFinish: int32(0)}
		o.Read(&message, "Touid")

		message.IsFinish = int32(1)
		errs, id := o.Update(&message)
		beego.Info(errs, id)
	}

	res_data := new(cspb.CSFriendmessageListRes)
	*res_data = cspb.CSFriendmessageListRes{
		Uid:                  &uid,
		FriendmessagelistNtf: FriendmessagelistNtf,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		FriendmessagelistRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kFriendmessageListRes),
		res_pkg_body, res_list)
	return ret

}
