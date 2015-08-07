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
	// var Friendntf []*cspb.CSFriendNtf
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
	err := c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player)
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
	var Friendntf_tag_1 []*cspb.CSFriendNtf
	var Friendntf_tag_2 []*cspb.CSFriendNtf
	var Friendntf_tag_3 []*cspb.CSFriendNtf

	var Friendntf_tag_5 []*cspb.CSFriendNtf
	m := int32(len(messages))
	for k := range messages {
		if messages[k].Tag == int32(1) {
			//赠送小红花

			Friendntf_tag_1 = append(Friendntf_tag_1, makeFriendntf(messages[k].Fromuid, messages[k].FromStarId, messages[k].Fromname, messages[k].Id))
			beego.Info("Friendntf_tag_1", Friendntf_tag_1)
		} else if messages[k].Tag == int32(2) {
			//接受小红花

			Friendntf_tag_2 = append(Friendntf_tag_2, makeFriendntf(messages[k].Fromuid, messages[k].FromStarId, messages[k].Fromname, messages[k].Id))
			beego.Info("Friendntf_tag_2", Friendntf_tag_2)
		} else if messages[k].Tag == int32(4) {
			//接受卡片

			Friendntf_tag_3 = append(Friendntf_tag_3, makeFriendntf(messages[k].Fromuid, messages[k].FromStarId, messages[k].Fromname, messages[k].Id))
			beego.Info("Friendntf_tag_3", Friendntf_tag_3)
		} else if messages[k].Tag == int32(3) {
			//赠送卡片
			var friend models.Player
			err := c.Find(bson.M{"uid": messages[k].Fromuid}).One(&friend)
			if err != nil {
				beego.Error(err)
			}

			for i := range friend.Cardrecord {

				if friend.Cardrecord[i].MessageId == messages[k].Id {
					var message models.Messages
					o := orm.NewOrm()
					var Friendntf_tag_4 []*cspb.CSFriendNtf
					message = models.Messages{Id: friend.Cardrecord[i].MessageId}
					o.Read(&message)
					beego.Info(message)
					Friendntf_tag_4 = append(Friendntf_tag_4, makeFriendntf(message.Fromuid, message.FromStarId, message.Fromname, messages[k].Id))
					switch friend.Cardrecord[i].CardId {
					case int32(1):
						FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), int32(1), friend.Cardrecord[i].CardNum, Friendntf_tag_4, m+3))
						// switch friend.Cardrecord[i].CardNum {
						// case int32(1):

						// 	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), int32(1), int32(1), Friendntf_tag_4, m+3))
						// case int32(5):

						// 	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), int32(1), int32(5), Friendntf_tag_4, m+3))
						// case int32(10):

						// 	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), int32(1), int32(10), Friendntf_tag_4, m+3))
						// }

					case int32(2):
						FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), int32(1), friend.Cardrecord[i].CardNum, Friendntf_tag_4, m+3))

						// var Friendntf_tag_4 []*cspb.CSFriendNtf
						// message = models.Messages{Id: friend.Cardrecord[i].MessageId}
						// o.Read(&message)
						// beego.Info(message)
						// Friendntf_tag_4 = append(Friendntf_tag_4, makeFriendntf(message.Fromuid, message.FromStarId, message.Fromname))

						// switch friend.Cardrecord[i].CardNum {
						// case int32(1):
						// 	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), int32(2), int32(1), Friendntf_tag_4, m+3))
						// case int32(5):
						// 	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), int32(2), int32(5), Friendntf_tag_4, m+3))
						// case int32(10):
						// 	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), int32(2), int32(10), Friendntf_tag_4, m+3))

						// }
					case int32(3):
						FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), int32(1), friend.Cardrecord[i].CardNum, Friendntf_tag_4, m+3))

						// var Friendntf_tag_4 []*cspb.CSFriendNtf
						// message = models.Messages{Id: friend.Cardrecord[i].MessageId}
						// o.Read(&message)
						// beego.Info(message)
						// Friendntf_tag_4 = append(Friendntf_tag_4, makeFriendntf(message.Fromuid, message.FromStarId, message.Fromname))

						// switch friend.Cardrecord[i].CardNum {
						// case int32(1):
						// 	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), int32(3), int32(1), Friendntf_tag_4, m+3))
						// case int32(5):
						// 	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), int32(3), int32(5), Friendntf_tag_4, m+3))
						// case int32(10):
						// 	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), int32(3), int32(10), Friendntf_tag_4, m+3))

						// }
					case 4:
						FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), int32(1), friend.Cardrecord[i].CardNum, Friendntf_tag_4, m+3))

						// var Friendntf_tag_4 []*cspb.CSFriendNtf
						// message = models.Messages{Id: friend.Cardrecord[i].MessageId}
						// o.Read(&message)
						// beego.Info(message)
						// Friendntf_tag_4 = append(Friendntf_tag_4, makeFriendntf(message.Fromuid, message.FromStarId, message.Fromname))

						// switch friend.Cardrecord[i].CardNum {
						// case int32(1):
						// 	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), int32(4), int32(1), Friendntf_tag_4, m+3))
						// case int32(5):
						// 	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), int32(4), int32(5), Friendntf_tag_4, m+3))
						// case int32(10):
						// 	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), int32(4), int32(10), Friendntf_tag_4, m+3))

						// }
					case 5:
						FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), int32(1), friend.Cardrecord[i].CardNum, Friendntf_tag_4, m+3))

						// var Friendntf_tag_4 []*cspb.CSFriendNtf
						// message = models.Messages{Id: friend.Cardrecord[i].MessageId}
						// o.Read(&message)
						// beego.Info(message)
						// Friendntf_tag_4 = append(Friendntf_tag_4, makeFriendntf(message.Fromuid, message.FromStarId, message.Fromname))

						// switch friend.Cardrecord[i].CardNum {
						// case int32(1):
						// 	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), int32(5), int32(1), Friendntf_tag_4, m+3))
						// case int32(5):
						// 	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), int32(5), int32(5), Friendntf_tag_4, m+3))
						// case int32(10):
						// 	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), int32(5), int32(10), Friendntf_tag_4, m+3))

						// }
					case 6:
						FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), int32(1), friend.Cardrecord[i].CardNum, Friendntf_tag_4, m+3))

						// var Friendntf_tag_4 []*cspb.CSFriendNtf
						// message = models.Messages{Id: friend.Cardrecord[i].MessageId}
						// o.Read(&message)
						// beego.Info(message)
						// Friendntf_tag_4 = append(Friendntf_tag_4, makeFriendntf(message.Fromuid, message.FromStarId, message.Fromname))

						// switch friend.Cardrecord[i].CardNum {
						// case int32(1):
						// 	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), int32(6), int32(1), Friendntf_tag_4, m+3))
						// case int32(5):
						// 	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), int32(6), int32(5), Friendntf_tag_4, m+3))
						// case int32(10):
						// 	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), int32(6), int32(10), Friendntf_tag_4, m+3))
						// }
					case 7:

						FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), int32(1), friend.Cardrecord[i].CardNum, Friendntf_tag_4, m+3))

						// var Friendntf_tag_4 []*cspb.CSFriendNtf
						// message = models.Messages{Id: friend.Cardrecord[i].MessageId}
						// o.Read(&message)
						// beego.Info(message)
						// Friendntf_tag_4 = append(Friendntf_tag_4, makeFriendntf(message.Fromuid, message.FromStarId, message.Fromname))

						// switch friend.Cardrecord[i].CardNum {
						// case int32(1):
						// 	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), int32(7), int32(1), Friendntf_tag_4, m+3))
						// case int32(5):
						// 	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), int32(7), int32(5), Friendntf_tag_4, m+3))
						// case int32(10):
						// 	FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(2), int32(7), int32(10), Friendntf_tag_4, m+3))

						// }

					}
				}
			}

		} else if messages[k].Tag == int32(5) {
			//添加好友

			Friendntf_tag_5 = append(Friendntf_tag_5, makeFriendntf(messages[k].Fromuid, messages[k].FromStarId, messages[k].Fromname, messages[k].Id))
			beego.Info("Friendntf_tag_5", Friendntf_tag_5)
		}

		// FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(messages[k].Messagetype, messages[k].ElementType, Friendntf, messages[k].CardId, "红色", messages[k].Number, messages[k].Id))
	}

	if len(Friendntf_tag_1) > 0 {
		FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(1), int32(1), int32(0), int32(1), Friendntf_tag_1, m))
	}
	if len(Friendntf_tag_2) > 0 {
		FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(2), int32(1), int32(0), int32(1), Friendntf_tag_2, m+1))
	}
	if len(Friendntf_tag_3) > 0 {
		for k := range messages {
			var friend models.Player
			err := c.Find(bson.M{"uid": messages[k].Fromuid}).One(&friend)
			if err != nil {
				beego.Error(err)
			}

			for i := range friend.Cardrecord {
				if friend.Cardrecord[i].MessageId == messages[k].Id {

					FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(1), int32(2), friend.Cardrecord[i].CardId, friend.Cardrecord[i].CardNum, Friendntf_tag_3, m+2))
				}
			}
		}
	}

	if len(Friendntf_tag_5) > 0 {
		FriendmessagelistNtf = append(FriendmessagelistNtf, makeFriendmessagelistNtf(int32(1), int32(3), int32(0), int32(0), Friendntf_tag_5, m+4))
	}
	beego.Info("FriendmessagelistNtf is :", FriendmessagelistNtf)
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
