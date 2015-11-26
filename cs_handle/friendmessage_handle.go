package cs_handle

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"

	models "tuojie.com/piggo/quickstart.git/models"
)
import cspb "protocol"

// import proto "code.google.com/p/goprotobuf/proto"

// import db "tuojie.com/piggo/quickstart.git/db/collection"
import "labix.org/v2/mgo/bson"

import db_session "tuojie.com/piggo/quickstart.git/db/session"

func FriendmessageHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********FriendmessageHandle Start**********")
	req_data := req.GetBody().GetFriendmessageReq()
	beego.Info(req_data)
	ret := int32(1)
	//测试代码
	isGive := int32(1)
	// uid := int32(2885377)
	//正式代码
	uid := int32(res_list.GetUid())
	c := db_session.DB("zoo").C("player")
	var player models.Player
	c.Find(bson.M{"uid": uid}).One(&player)
	//mongodb中卡片给出初始化的值
	beego.Info("player:", player)
	beego.Info("card_account", beego.AppConfig.String("card_account"))
	beego.Info("cards len", len(player.Cards))
	if len(player.Cards) == 0 {
		for i := 0; i <= stringToint(beego.AppConfig.String("card_account")); i++ {
			_, err := c.Upsert(bson.M{"uid": uid},
				bson.M{"$set": bson.M{"cards." + fmt.Sprint(i) + ".card_id": (i + 1), "cards." + fmt.Sprint(i) + ".card_num": int32(0)}})
			if err != nil {
				beego.Error("插入失败")
			} else {
				beego.Info("插入成功")
			}
		}
	}
	if req_data.GetMessageType() == int32(2) {

		//同意赠送
		if req_data.GetOperationType() == int32(1) {

			//赠送,产生消息
			beego.Info("赠送,产生消息")
			switch req_data.GetElementType() {
			case int32(1):
				//红花
				beego.Info("赠送红花")
				// for i := range req_data.GetMessagesNtf() {
				// 	playerId := req_data.GetMessagesNtf()[i].GetPlayuid()
				// 	c.Find(bson.M{"uid": playerId}).One(&player)
				// 	player.Flower += req_data.GetElement()[1].GetElementNum()
				// 	c.Upsert(bson.M{"uid": playerId},
				// 		bson.M{"$set": bson.M{"flower": player.Flower}})
				// }

			case int32(2):
				// 卡片
				beego.Info("赠送卡片")
				for i := range req_data.GetMessagesNtf() {
					if req_data.GetMessagesNtf()[i].GetStatus() != int32(2) {
						for j := range req_data.GetMessagesNtf()[i].GetElement() {

							_, err := c.Upsert(bson.M{"uid": uid},
								bson.M{"$inc": bson.M{"cards." + fmt.Sprint(req_data.GetMessagesNtf()[i].GetElement()[j].GetCardId()-1) + ".card_num": -req_data.GetMessagesNtf()[i].GetElement()[j].GetElementNum()}})

							if err != nil {
								beego.Error("减少卡片失败")
								isGive = int32(0)
							} else {
								beego.Info("减少卡片成功")
							}
						}
					}
				}

			}
		}

		//生成消息存在mysql

	} else if req_data.GetMessageType() == int32(1) && req_data.GetOperationType() != int32(3) {
		//请求赠送

		switch req_data.GetElementType() {
		case int32(1):
			beego.Info("请求赠送红花")
		case int32(2):
			if req_data.GetOperationType() == int32(1) {
				beego.Info("赠送卡片")
				for i := range req_data.GetMessagesNtf() {
					if req_data.GetMessagesNtf()[i].GetStatus() != int32(2) {
						for j := range req_data.GetMessagesNtf()[i].GetElement() {
							_, err := c.Upsert(bson.M{"uid": uid},
								bson.M{"$inc": bson.M{"cards." + fmt.Sprint(req_data.GetMessagesNtf()[i].GetElement()[j].GetCardId()-1) + ".card_num": -req_data.GetMessagesNtf()[i].GetElement()[j].GetElementNum()}})
							if err != nil {
								beego.Error("减少卡片失败")

								isGive = int32(0)
							} else {
								beego.Info("减少卡片成功")

							}
							// _, errs := c.Upsert(bson.M{"uid": uid},
							// 	bson.M{"$push": bson.M{"cardrecord": bson.M{"message_id": messageid, "card_id": req_data.GetMessagesNtf()[i].GetElement()[j].GetCardId(), "card_num": req_data.GetMessagesNtf()[i].GetElement()[j].GetElementNum()}}})

						}
					}
				}
				//card记录表里面要生成记录

			}
			beego.Info("请求赠送卡片")
		}
	} else {

		//传的值有问题
		beego.Error("value is Error")

	}
	if req_data.GetOperationType() == int32(2) {
		//接受，处理掉消息
		beego.Info("接受，处理消息")
		switch req_data.GetElementType() {
		case int32(1):
			//红花
			beego.Info("接受小红花")
			for i := range req_data.GetMessagesNtf() {
				_, err := c.Upsert(bson.M{"uid": uid},
					bson.M{"$inc": bson.M{"flower": 1}})
				if err == nil {
					beego.Info("接受小红花成功！", i)
				} else {
					beego.Error("接受小红花失败！", i)
				}
			}
			isGive = int32(1)
		case int32(2):
			//卡片
			beego.Info("接受卡片")

			c.Find(bson.M{"uid": uid}).One(&player)

			for i := range req_data.GetMessagesNtf() {
				if req_data.GetMessagesNtf()[i].GetStatus() != int32(2) {
					beego.Info(i)
					for j := range req_data.GetMessagesNtf()[i].GetElement() {
						beego.Info(j)

						for n := range player.Cards {

							if player.Cards[n].CardId == req_data.GetMessagesNtf()[i].GetElement()[j].GetCardId() {

								_, err := c.Upsert(bson.M{"uid": uid},
									bson.M{"$inc": bson.M{"cards." + fmt.Sprint(n) + ".card_num": req_data.GetMessagesNtf()[i].GetElement()[j].GetElementNum()}})

								if err != nil {
									beego.Error("插入失败")
									isGive = int32(0)
								} else {
									beego.Info("插入成功")
								}
							}

						}

					}
				}
			}

		}
	}
	if req_data.GetElementType() == int32(3) && req_data.GetOperationType() != int32(3) {
		//加好友的消息
		beego.Info("加好友消息")
		//同意就在表里加一条记录，不同意不操作
		for i := range req_data.GetMessagesNtf() {
			if req_data.GetMessagesNtf()[i].GetStatus() != int32(2) {
				for j := range player.FriendList {
					if req_data.GetMessagesNtf()[i].GetPlayuid() == player.FriendList[j].Friendid {
						_, err := c.Upsert(bson.M{"uid": uid},
							bson.M{"$set": bson.M{"FriendList." + fmt.Sprint(j) + ".friendid": req_data.GetMessagesNtf()[i].GetPlayuid(), "FriendList." + fmt.Sprint(j) + ".isActive": int32(1), "FriendList." + fmt.Sprint(j) + ".accepttime": time.Now().Unix()}})

						if err != nil {
							beego.Error("同意好友申请失败")
							isGive = int32(0)
						} else {
							beego.Info("同意好友申请成功")
						}

						//对方的朋友表
						// c.Upsert(bson.M{"uid": req_data.GetMessagesNtf()[i].GetPlayuid(), "FriendList.friendid": uid},
						// 	bson.M{"$set": bson.M{"FriendList.$.isActive": int32(1), "FriendList.$.accepttime": time.Now().Unix()}})
					}

				}
				//自己的朋友申请列表
				for k := range player.ApplyFriendList {
					if req_data.GetMessagesNtf()[i].GetPlayuid() == player.ApplyFriendList[k].Applyuid {
						_, err := c.Upsert(bson.M{"uid": uid},
							bson.M{"$set": bson.M{"ApplyFriendList." + fmt.Sprint(k) + ".isAccept": int32(1), "ApplyFriendList." + fmt.Sprint(k) + ".isrefuse": int32(0), "ApplyFriendList." + fmt.Sprint(k) + ".oprationtime": time.Now().Unix()}})
						if err != nil {
							beego.Error("同意好友,申请列表变更失败")
							isGive = int32(0)
						} else {
							beego.Info("同意好友,申请列表变更成功")
						}
					}
				}

				//朋友的朋友列表变更
				var friend models.Player
				c.Find(bson.M{"uid": req_data.GetMessagesNtf()[i].GetPlayuid()}).One(&friend)
				for m := range friend.FriendList {
					if friend.FriendList[m].Friendid == uid {
						_, err := c.Upsert(bson.M{"uid": req_data.GetMessagesNtf()[i].GetPlayuid()},
							bson.M{"$set": bson.M{"FriendList." + fmt.Sprint(m) + ".isActive": int32(1), "FriendList." + fmt.Sprint(m) + ".accepttime": time.Now().Unix()}})

						if err != nil {
							beego.Error("同意好友,好友的好友列表变更失败")
							isGive = int32(0)
						} else {
							beego.Info("同意好友,好友的好友列表变更成功")
						}
					}
				}

				//mysql朋友列表添加
				o := orm.NewOrm()
				var Friend models.Friend

				for i := range player.FriendList {
					Friend = models.Friend{FriendId: player.FriendList[i].Friendid}
					err := o.Read(&Friend, "Friendid")
					beego.Info(Friend)
					if err != nil {
						Friend.Uid = uid
						Friend.FriendId = player.FriendList[i].Friendid
						id, err := o.Insert(&Friend)
						beego.Info(id, err)

						Friend.Uid = player.FriendList[i].Friendid
						Friend.FriendId = uid
						id1, err1 := o.Insert(&Friend)
						beego.Info(id1, err1)
					}

				}
			} else if req_data.GetMessagesNtf()[i].GetStatus() == int32(2) {
				for j := range player.FriendList {
					if req_data.GetMessagesNtf()[i].GetPlayuid() == player.FriendList[j].Friendid {
						_, err := c.Upsert(bson.M{"uid": uid},
							bson.M{"$pull": bson.M{"FriendList": bson.M{"friendid": req_data.GetMessagesNtf()[i].GetPlayuid()}}})

						if err != nil {
							beego.Error("删除好友申请失败")
							isGive = int32(0)
						} else {
							beego.Info("删除好友申请成功")
						}

						//对方的朋友表
						// c.Upsert(bson.M{"uid": req_data.GetMessagesNtf()[i].GetPlayuid(), "FriendList.friendid": uid},
						// 	bson.M{"$set": bson.M{"FriendList.$.isActive": int32(1), "FriendList.$.accepttime": time.Now().Unix()}})
					}

				}
				//自己的朋友申请列表
				for k := range player.ApplyFriendList {
					if req_data.GetMessagesNtf()[i].GetPlayuid() == player.ApplyFriendList[k].Applyuid {
						_, err := c.Upsert(bson.M{"uid": uid},
							bson.M{"$pull": bson.M{"ApplyFriendList": bson.M{"applyuid": player.ApplyFriendList[k].Applyuid}}})
						if err != nil {
							beego.Error("同意好友,申请列表变更失败")
							isGive = int32(0)
						} else {
							beego.Info("同意好友,申请列表变更成功")
						}
					}
				}

				//朋友的朋友列表变更
				var friend models.Player
				c.Find(bson.M{"uid": req_data.GetMessagesNtf()[i].GetPlayuid()}).One(&friend)
				for m := range friend.FriendList {
					if friend.FriendList[m].Friendid == uid {
						_, err := c.Upsert(bson.M{"uid": req_data.GetMessagesNtf()[i].GetPlayuid()},
							bson.M{"$pull": bson.M{"FriendList": bson.M{"friendid": uid}}})

						if err != nil {
							beego.Error("删除好友的好友列表变更失败")
							isGive = int32(0)
						} else {
							beego.Info("删除好友的好友列表变更成功")
						}
					}
				}
			}
		}
	}
	//消息表
	c.Find(bson.M{"c_account": res_list.GetCAccount()}).One(&player)
	if req_data.GetElementType() != int32(3) && req_data.GetOperationType() != int32(3) && req_data.GetOperationType() != int32(2) {
		//生成消息mysql
		o := orm.NewOrm()
		var messages models.Messages
		messages.Fromuid = int32(res_list.GetUid())
		messages.Fromname = player.Name
		messages.FromStarId = player.StarId
		messages.Time = time.Now().Unix()
		messages.IsFinish = int32(0)

		messages.ElementType = req_data.GetElementType()

		if req_data.GetMessageType() == int32(2) && req_data.GetElementType() == int32(1) && req_data.GetOperationType() == int32(1) {
			//接受小红花
			messages.Messagetype = int32(2)
			messages.Tag = int32(2)
		} else if req_data.GetMessageType() == int32(1) && req_data.GetElementType() == int32(1) && req_data.GetOperationType() != int32(1) && req_data.GetOperationType() != int32(3) {
			//赠送小红花
			messages.Messagetype = int32(1)
			messages.Tag = int32(1)
		} else if req_data.GetMessageType() == int32(2) && req_data.GetElementType() == int32(2) && req_data.GetOperationType() != int32(3) {
			//主动赠送：接受卡片
			messages.Messagetype = int32(2)
			messages.Tag = int32(4)
		} else if req_data.GetMessageType() == int32(1) && req_data.GetElementType() == int32(1) && req_data.GetOperationType() == int32(1) {
			//接受小红花
			messages.Messagetype = int32(2)
			messages.Tag = int32(2)
		} else if req_data.GetMessageType() == int32(1) && req_data.GetElementType() == int32(2) && req_data.GetOperationType() == int32(1) {
			//申请赠送：接受卡片
			messages.Messagetype = int32(2)
			messages.Tag = int32(4)
		}
		for i := range req_data.GetMessagesNtf() {
			var meesages_arr []models.Messages
			var cond *orm.Condition
			cond = orm.NewCondition()

			cond = cond.And("Id__gte", 1)
			cond = cond.And("Touid__contains", req_data.GetMessagesNtf()[i].GetPlayuid())
			cond = cond.And("Fromuid__contains", int32(res_list.GetUid()))
			cond = cond.And("IsFinish__contains", int32(0))
			beego.Info(cond)
			var qs orm.QuerySeter
			qs = orm.NewOrm().QueryTable("messages").SetCond(cond)

			cnt, err := qs.All(&meesages_arr)
			beego.Info(cnt, err, meesages_arr)

			if req_data.GetMessageType() == int32(2) && req_data.GetElementType() == int32(2) && req_data.GetOperationType() == int32(0) {
				//主动赠送卡片情况，可以多次赠送
				if req_data.GetMessagesNtf()[i].GetStatus() != int32(2) {

					messages.Touid = req_data.GetMessagesNtf()[i].GetPlayuid()
					messageid, err := o.Insert(&messages)
					if err != nil {
						beego.Error(err)
						isGive = int32(0)
					}

					beego.Info(messageid)
					if req_data.GetOperationType() == int32(1) && req_data.GetElementType() == int32(2) {

						for j := range req_data.GetMessagesNtf()[i].GetElement() {

							_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
								bson.M{"$push": bson.M{"cardrecord": bson.M{"message_id": messageid, "card_id": req_data.GetMessagesNtf()[i].GetElement()[j].GetCardId(), "card_num": req_data.GetMessagesNtf()[i].GetElement()[j].GetElementNum()}}})
							if err != nil {
								beego.Error("主动赠送卡片存储失败")
								isGive = int32(0)
							} else {
								beego.Info("主动赠送卡片存储成功")
							}

						}

					}
				}
			} else {

				if req_data.GetMessagesNtf()[i].GetStatus() != int32(2) {

					messages.Touid = req_data.GetMessagesNtf()[i].GetPlayuid()
					messageid, err := o.Insert(&messages)
					if err != nil {
						beego.Error(err)
						isGive = int32(0)
					}

					beego.Info(messageid)
					if req_data.GetOperationType() == int32(1) && req_data.GetElementType() == int32(2) {
						//点击消息里的赠送卡片按钮
						for j := range req_data.GetMessagesNtf()[i].GetElement() {

							_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
								bson.M{"$push": bson.M{"cardrecord": bson.M{"message_id": messageid, "card_id": req_data.GetMessagesNtf()[i].GetElement()[j].GetCardId(), "card_num": req_data.GetMessagesNtf()[i].GetElement()[j].GetElementNum()}}})
							if err != nil {
								beego.Error("赠送卡片存储失败")
								isGive = int32(0)
							} else {
								beego.Info("赠送卡片存储成功")
							}

						}

					}
				}

			}
		}
	}
	if req_data.GetOperationType() == int32(3) {
		//删除消息
		var message models.Messages
		o := orm.NewOrm()
		if req_data.GetElementType() == 2 {
			//删除卡片消息
			mysqlid := req_data.GetMysqlIdNtf()[0].GetMysqlId()
			message = models.Messages{Id: mysqlid}
			o.Read(&message)
			beego.Info(message)
			var rmplayer models.Player
			c.Find(bson.M{"uid": message.Fromuid}).One(&rmplayer)
			rm_tag := int32(0)
			for i := range rmplayer.Cardrecord {
				for j := range req_data.GetMessagesNtf() {
					for k := range req_data.GetMessagesNtf()[j].GetElement() {
						if rmplayer.Cardrecord[i].MessageId == mysqlid && rmplayer.Cardrecord[i].CardId == req_data.GetMessagesNtf()[j].GetElement()[k].GetCardId() {
							rm_tag++
							c.Upsert(bson.M{"uid": message.Fromuid},
								bson.M{"$pull": bson.M{"cardrecord": bson.M{"message_id": mysqlid, "card_id": req_data.GetMessagesNtf()[j].GetElement()[k].GetCardId()}}})
						}
					}
				}
			}
			if rm_tag <= int32(1) {
				//当卡片消息删完的时候，就删掉这条消息
				for i := range req_data.GetMysqlIdNtf() {
					beego.Info(i)
					message = models.Messages{Id: req_data.GetMysqlIdNtf()[i].GetMysqlId()}
					o.Read(&message)
					beego.Info(message)
					message.IsFinish = int32(1)
					id, err := o.Update(&message, "IsFinish")
					beego.Info(err, id)
				}
			}
			isGive = int32(1)
		} else {

			for i := range req_data.GetMysqlIdNtf() {
				beego.Info(i)
				message = models.Messages{Id: req_data.GetMysqlIdNtf()[i].GetMysqlId()}
				o.Read(&message)
				beego.Info(message)
				message.IsFinish = int32(1)
				id, err := o.Update(&message, "IsFinish")
				beego.Info(err, id)
			}

			if req_data.GetElementType() == int32(3) {
				beego.Info("加好友消息")
				//同意就在表里加一条记录，不同意不操作
				for i := range req_data.GetMessagesNtf() {

					for j := range player.FriendList {
						if req_data.GetMessagesNtf()[i].GetPlayuid() == player.FriendList[j].Friendid {
							_, err := c.Upsert(bson.M{"uid": uid},
								bson.M{"$pull": bson.M{"FriendList": bson.M{"friendid": req_data.GetMessagesNtf()[i].GetPlayuid()}}})

							if err != nil {
								beego.Error("删除好友申请失败")
								isGive = int32(0)
							} else {
								beego.Info("删除好友申请成功")
							}

							//对方的朋友表
							// c.Upsert(bson.M{"uid": req_data.GetMessagesNtf()[i].GetPlayuid(), "FriendList.friendid": uid},
							// 	bson.M{"$set": bson.M{"FriendList.$.isActive": int32(1), "FriendList.$.accepttime": time.Now().Unix()}})
						}

					}
					//自己的朋友申请列表
					for k := range player.ApplyFriendList {
						if req_data.GetMessagesNtf()[i].GetPlayuid() == player.ApplyFriendList[k].Applyuid {
							_, err := c.Upsert(bson.M{"uid": uid},
								bson.M{"$pull": bson.M{"ApplyFriendList": bson.M{"applyuid": player.ApplyFriendList[k].Applyuid}}})
							if err != nil {
								beego.Error("同意好友,申请列表变更失败")
								isGive = int32(0)
							} else {
								beego.Info("同意好友,申请列表变更成功")
							}
						}
					}

					//朋友的朋友列表变更
					var friend models.Player
					c.Find(bson.M{"uid": req_data.GetMessagesNtf()[i].GetPlayuid()}).One(&friend)
					for m := range friend.FriendList {
						if friend.FriendList[m].Friendid == uid {
							_, err := c.Upsert(bson.M{"uid": req_data.GetMessagesNtf()[i].GetPlayuid()},
								bson.M{"$pull": bson.M{"FriendList": bson.M{"friendid": uid}}})

							if err != nil {
								beego.Error("删除好友的好友列表变更失败")
								isGive = int32(0)
							} else {
								beego.Info("删除好友的好友列表变更成功")
							}
						}
					}

				}
			}
		}

	}
	//已发送的消息列表删除

	var message models.Messages
	o := orm.NewOrm()
	for i := range req_data.GetMysqlIdNtf() {
		beego.Info(i)
		message = models.Messages{Id: req_data.GetMysqlIdNtf()[i].GetMysqlId()}
		o.Read(&message)
		beego.Info(message)
		message.IsFinish = int32(1)
		id, err := o.Update(&message, "IsFinish")
		beego.Info(err, id)
	}

	var player_return models.Player
	c.Find(bson.M{"uid": uid}).One(&player_return)
	var CardNtf []*cspb.CSCardNtf
	for i := range player_return.Cards {
		for j := range req_data.GetMessagesNtf() {

			for k := range req_data.GetMessagesNtf()[j].GetElement() {
				beego.Info(k)
				if player_return.Cards[i].CardId == req_data.GetMessagesNtf()[j].GetElement()[k].GetCardId() {
					beego.Info("find cardid")
					CardNtf = append(CardNtf, makecardNtf1(player_return.Cards[i].CardId, player_return.Cards[i].CardNum))
				}
			}
		}

	}

	res_data := new(cspb.CSFriendmessageRes)
	messageId := req_data.GetMessageId()
	friendListId := req_data.GetFriendListId()
	elementType := req_data.GetElementType()

	beego.Info("isGive =", isGive)
	beego.Info(CardNtf)
	*res_data = cspb.CSFriendmessageRes{
		IsGive:       &isGive,
		Uid:          &uid,
		MessageId:    &messageId,
		FriendListId: &friendListId,
		ElementType:  &elementType,
		CardNtf:      CardNtf,
		Flower:       &player_return.Flower,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{

		FriendmessageRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kFriendmessageRes),
		res_pkg_body, res_list)
	return ret

}
