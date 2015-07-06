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

func FriendmessageListHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********FriendmessageHandle Start**********")
	req_data := req.GetBody().GetFriendmessagelistReq()
	beego.Info(req_data)
	ret := int32(1)
	//测试代码
	var Friendntf []*cspb.CSFriendNtf
	// for i := range resmgr.FriendntftestData.GetItems() {
	// 	Friendntf = append(Friendntf, makeFriendntf(resmgr.FriendntftestData.GetItems()[i].GetFriendId(), resmgr.FriendntftestData.GetItems()[i].GetStarId(), resmgr.FriendntftestData.GetItems()[i].GetName()))
	// }
	var FriendmessagelistNtf []*cspb.CSFriendmessagelistNtf
	// for j := range resmgr.FriendmessagelisttestData.GetItems() {
	// 	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(resmgr.FriendmessagelisttestData.GetItems()[j].GetMessageType(), resmgr.FriendmessagelisttestData.GetItems()[j].GetElementType(),
	// 		Friendntf, resmgr.FriendmessagelisttestData.GetItems()[j].GetCardId(), resmgr.FriendmessagelisttestData.GetItems()[j].GetCardColor(), resmgr.FriendmessagelisttestData.GetItems()[j].GetElementNum(),
	// 		resmgr.FriendmessagelisttestData.GetItems()[j].GetMessageId()))
	// }
	// uid := int32(2885377)
	uid := int32(res_list.GetUid())
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

	cond = cond.And("Id__gte", 1)
	cond = cond.And("Touid__contains", int32(res_list.GetUid()))
	cond = cond.And("IsFinish__contains", int32(0))
	beego.Info(cond)
	var qs orm.QuerySeter
	qs = orm.NewOrm().QueryTable("messages").SetCond(cond)

	cnt, err := qs.All(&messages)
	beego.Info(cnt, err, messages)
	//从sql中取出消息,5种消息类型

	for k := range messages {
		if messages[k].Tag == int32(1) {
			Friendntf_tag_1 = append(Friendntf, makeFriendntf(messages[k].Fromuid, messages[k].FromStarId, messages[k].Fromname))

		} else if messages[k].Tag == int32(2) {
			Friendntf_tag_2 = append(Friendntf, makeFriendntf(messages[k].Fromuid, messages[k].FromStarId, messages[k].Fromname))
		} else if messages[k].Tag == int32(3) {
			Friendntf_tag_3 = append(Friendntf, makeFriendntf(messages[k].Fromuid, messages[k].FromStarId, messages[k].Fromname))

		} else if messages[k].Tag == int32(4) {
			Friendntf_tag_4 = append(Friendntf, makeFriendntf(messages[k].Fromuid, messages[k].FromStarId, messages[k].Fromname))

		} else if messages[k].Tag == int32(5) {
			Friendntf_tag_5 = append(Friendntf, makeFriendntf(messages[k].Fromuid, messages[k].FromStarId, messages[k].Fromname))
		}

		// FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(messages[k].Messagetype, messages[k].ElementType, Friendntf, messages[k].CardId, "红色", messages[k].Number, messages[k].Id))
	}
	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(1), int32(1), Friendntf_tag_1, int32(0), "", messages[k].Number, messages[k].Id))
	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(1), Friendntf_tag_2, int32(0), "红色", messages[k].Number, messages[k].Id))
	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(1), int32(2), Friendntf_tag_3, messages[k].CardId, "红色", messages[k].Number, messages[k].Id))
	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), Friendntf_tag_4, messages[k].CardId, "红色", messages[k].Number, messages[k].Id))
	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(3), Friendntf_tag_5, messages[k].CardId, "红色", messages[k].Number, messages[k].Id))

	beego.Info("FriendmessagelistNtf is :", FriendmessagelistNtf)
	//发完之后就讲mysql中的Isfinish变为1
	var message models.Messages
	o := orm.NewOrm()
	for i := range messages {
		beego.Info(i)
		message = models.Messages{Touid: int32(res_list.GetUid()), IsFinish: int32(0)}
		o.Read(&message, "Touid", "IsFinish")
		beego.Info(message)
		message.IsFinish = int32(1)
		id, err := o.Update(&message, "IsFinish")
		beego.Info(err, id)
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
