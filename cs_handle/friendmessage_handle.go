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
	switch req_data.GetMessageType() {
	case int32(1):
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
				for j := range req_data.GetMessagesNtf()[i].GetElement() {

					_, err := c.Upsert(bson.M{"uid": uid},
						bson.M{"$set": bson.M{"cards." + fmt.Sprint(req_data.GetMessagesNtf()[i].GetElement()[j].GetCardId()-1) + ".card_id": req_data.GetMessagesNtf()[i].GetElement()[j].GetCardId()}})
					c.Upsert(bson.M{"uid": uid},
						bson.M{"$set": bson.M{"cards." + fmt.Sprint(req_data.GetMessagesNtf()[i].GetElement()[j].GetCardId()-1) + ".card_num": -req_data.GetMessagesNtf()[i].GetElement()[j].GetElementNum()}})

					if err != nil {
						beego.Error("减少卡片失败")
						isGive = int32(0)
					} else {
						beego.Info("减少卡片成功")
					}

				}
			}

		}

		//生成消息存在mysql

	case int32(2):
		//接受，处理掉消息
		beego.Info("接受，处理消息")
		switch req_data.GetElementType() {
		case int32(1):
			//红花
			beego.Info("接受小红花")
			c.Find(bson.M{"uid": uid}).One(&player)
			for i := range req_data.GetMessagesNtf() {
				for j := range req_data.GetMessagesNtf()[i].GetElement() {
					_, err := c.Upsert(bson.M{"uid": uid},
						bson.M{"$inc": bson.M{"flower": req_data.GetMessagesNtf()[i].GetElement()[j].GetElementNum()}})
					if err != nil {
						beego.Error("增加失败")
						isGive = int32(0)
					} else {
						beego.Info("增加小红花成功")
					}
				}
			}
		case int32(2):
			//卡片
			beego.Info("接受卡片")
			c.Find(bson.M{"uid": uid}).One(&player)

			for i := range req_data.GetMessagesNtf() {
				beego.Info(i)
				for j := range req_data.GetMessagesNtf()[i].GetElement() {
					beego.Info(j)

					for n := range player.Cards {

						if player.Cards[n].CardId == req_data.GetMessagesNtf()[i].GetElement()[j].GetCardId() {
							_, err := c.Upsert(bson.M{"uid": uid},
								bson.M{"$set": bson.M{"cards." + fmt.Sprint(n) + ".card_id": req_data.GetMessagesNtf()[i].GetElement()[j].GetCardId()}})
							c.Upsert(bson.M{"uid": uid},
								bson.M{"$set": bson.M{"cards." + fmt.Sprint(n) + ".card_num": req_data.GetMessagesNtf()[i].GetElement()[j].GetElementNum()}})

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
	default:
		//传的值有问题
		beego.Error("value is Error")
		isGive = int32(0)
	}
	if req_data.GetElementType() == int32(3) {
		//加好友的消息
		beego.Info("加好友消息")
		//同意就在表里加一条记录，不同意不操作
		for i := range req_data.GetMessagesNtf() {

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
				}
			}
		}
	}
	//消息表
	c.Find(bson.M{"c_account": res_list.GetCAccount()}).One(&player)
	if req_data.GetElementType() != int32(3) {
		o := orm.NewOrm()
		var messages models.Messages
		messages.Fromuid = int32(res_list.GetUid())
		messages.Fromname = player.Name
		messages.FromStarId = player.StarId
		messages.Time = time.Now().Unix()
		messages.IsFinish = int32(0)
		messages.Messagetype = req_data.GetMessageType()
		messages.ElementType = req_data.GetElementType()

		if req_data.GetMessageType() == int32(1) && req_data.GetElementType() == int32(1) {
			//接受小红花
			messages.Tag = int32(1)
		} else if req_data.GetMessageType() == int32(2) && req_data.GetElementType() == int32(1) {
			//赠送小红花
			messages.Tag = int32(2)
		} else if req_data.GetMessageType() == int32(2) && req_data.GetElementType() == int32(2) {
			//接受卡片
			messages.Tag = int32(3)
		}
		for i := range req_data.GetMessagesNtf() {
			messages.Touid = req_data.GetMessagesNtf()[i].GetPlayuid()
			messageid, err := o.Insert(&messages)
			if err != nil {
				beego.Error(err)
				isGive = int32(0)
			}

			beego.Info(messageid)
			if req_data.GetMessageType() == int32(1) && req_data.GetElementType() == int32(2) {
				m := len(player.Cardrecord)

				beego.Info(m)

				for j := range req_data.GetMessagesNtf()[i].GetElement() {
					if m == 0 {
						_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
							bson.M{"$push": bson.M{"cardrecord": bson.M{"message_id": messageid, "card_id": req_data.GetMessagesNtf()[i].GetElement()[j].GetCardId(), "card_num": req_data.GetMessagesNtf()[i].GetElement()[j].GetElementNum()}}})
						if err != nil {
							beego.Error("首次赠送卡片存储失败")
							isGive = int32(0)
						} else {
							beego.Info("首次赠送卡片存储成功")
						}
					} else {
						_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
							bson.M{"$set": bson.M{"cardrecord." + fmt.Sprint(m) + ".message_id": messageid, "cardrecord." + fmt.Sprint(m) + ".card_id": req_data.GetMessagesNtf()[i].GetElement()[j].GetCardId(), "cardrecord." + fmt.Sprint(m) + ".card_num": req_data.GetMessagesNtf()[i].GetElement()[j].GetElementNum()}})

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

	//已发送的消息列表删除

	// var message models.Messages
	// o := orm.NewOrm()
	// for i := range req_data.GetMysqlIdNtf() {
	// 	beego.Info(i)
	// 	message = models.Messages{Id: req_data.GetMysqlIdNtf()[i].GetMysqlId()}
	// 	o.Read(&message)
	// 	beego.Info(message)
	// 	message.IsFinish = int32(1)
	// 	id, err := o.Update(&message, "IsFinish")
	// 	beego.Info(err, id)
	// }
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
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{

		FriendmessageRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kFriendmessageRes),
		res_pkg_body, res_list)
	return ret

}
