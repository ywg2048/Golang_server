package cs_handle

import (
	// "fmt"
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
	uid := int32(2885377)
	//正式代码
	c := db_session.DB("zoo").C("player")
	var player models.Player

	switch req_data.GetMessageType() {
	case int32(1):
		//赠送,产生消息
		switch req_data.GetElementType() {
		case int32(1):
			//红花
			// for i := range req_data.GetMessagesNtf() {
			// 	playerId := req_data.GetMessagesNtf()[i].GetPlayuid()
			// 	c.Find(bson.M{"uid": playerId}).One(&player)
			// 	player.Flower += req_data.GetElement()[1].GetElementNum()
			// 	c.Upsert(bson.M{"uid": playerId},
			// 		bson.M{"$set": bson.M{"flower": player.Flower}})
			// }

		case int32(2):
			// 卡片
			for i := range req_data.GetMessagesNtf() {
				playerId := req_data.GetMessagesNtf()[i].GetPlayuid()
				c.Find(bson.M{"uid": playerId}).One(&player)

				for m := range req_data.GetElement() {
					for k := range player.Cards {
						if player.Cards[k].CardId == req_data.GetElement()[m].GetCardId() {
							player.Cards[k].CardNum += req_data.GetElement()[m].GetElementNum()

							//自己的卡片减少
							var player_self models.Player
							c.Find(bson.M{"uid": uid}).One(&player_self)

							//当前角色的小伙伴扣除卡片
							for a := range player_self.Cards {
								if player_self.Cards[a].CardId == req_data.GetElement()[m].GetCardId() {
									c.Upsert(bson.M{"uid": uid},
										bson.M{"$set": bson.M{"cards.card_num": player.Cards[a].CardNum - req_data.GetElement()[m].GetElementNum()}})
								}
							}
						}
					}

				}
			}

		case int32(3):
			//加好友的消息

			//同意就在表里加一条记录，不同意不操作
			c.Find(bson.M{"uid": uid}).One(&player)

			c.Upsert(bson.M{"uid": uid},
				bson.M{"$push": bson.M{"FriendList.friendid": req_data.GetMessagesNtf()[1].GetPlayuid(), "FriendList.isActive": int32(1), "FriendList.accepttime": time.Now().Unix()}})

			c.Update(bson.M{"uid": uid},
				bson.M{"$pull": bson.M{"ApplyFriendList.Applyuid": req_data.GetMessagesNtf()[1].GetPlayuid()}})
			//对方的朋友表
			c.Upsert(bson.M{"uid": req_data.GetMessagesNtf()[1].GetPlayuid(), "FriendList.friendid": uid},
				bson.M{"$set": bson.M{"FriendList.$.isActive": int32(1), "FriendList.$.accepttime": time.Now().Unix()}})
		}

		//生成消息存在mysql

	case int32(2):
		//接受，处理掉消息
		switch req_data.GetElementType() {
		case int32(1):
			//红花
			c.Find(bson.M{"uid": uid}).One(&player)
			for i := range req_data.GetElement() {
				c.Upsert(bson.M{"uid": uid},
					bson.M{"$inc": bson.M{"flower": req_data.GetElement()[i].GetElementNum()}})
			}
		case int32(2):
			//卡片
			c.Find(bson.M{"uid": uid}).One(&player)
			for i := range req_data.GetElement() {
				for j := range player.Star {
					if player.StarId == player.Star[j].StarId {
						//将卡片添加到当前明星中
						c.Upsert(bson.M{"uid": uid, "star.cards.card_id": req_data.GetElement()[i].GetCardId()},
							bson.M{"$inc": bson.M{"star.cards.$.card_num": req_data.GetElement()[i].GetElementNum()}})
					}
				}
			}

		case int32(3):
			//加好友的消息
			c.Find(bson.M{"uid": uid}).One(&player)

			c.Upsert(bson.M{"uid": uid},
				bson.M{"$push": bson.M{"FriendList.friendid": req_data.GetMessagesNtf()[1].GetPlayuid(), "FriendList.isActive": int32(1), "FriendList.accepttime": time.Now().Unix()}})

			c.Update(bson.M{"uid": uid},
				bson.M{"$pull": bson.M{"ApplyFriendList.Applyuid": req_data.GetMessagesNtf()[1].GetPlayuid()}})
			//对方的朋友表
			c.Upsert(bson.M{"uid": req_data.GetMessagesNtf()[1].GetPlayuid(), "FriendList.friendid": uid},
				bson.M{"$set": bson.M{"FriendList.$.isActive": int32(1), "FriendList.$.accepttime": time.Now().Unix()}})
		}
	default:
		//传的值有问题
		beego.Error("value is Error")
		isGive = int32(0)
	}
	//消息表
	c.Find(bson.M{"c_account": res_list.GetCAccount()}).One(&player)

	o := orm.NewOrm()
	var messages models.Messages
	messages.Fromuid = int32(res_list.GetUid())
	messages.Fromname = player.Name
	messages.FromStarId = player.StarId
	messages.Time = time.Now().Unix()
	messages.IsFinish = int32(0)
	messages.Messagetype = req_data.GetMessageType()
	messages.ElementType = req_data.GetElementType()
	if req_data.GetElementType() == int32(2) {
		//卡片类型
	}
	if req_data.GetMessageType() == int32(1) && req_data.GetElementType() == int32(1) {
		//接受小红花
		messages.Tag = int32(1)
	} else if req_data.GetMessageType() == int32(2) && req_data.GetElementType() == int32(1) {
		//赠送小红花
		messages.Tag = int32(2)
	} else if req_data.GetMessageType() == int32(1) && req_data.GetElementType() == int32(2) {
		//接受卡片
		messages.Tag = int32(3)
	} else if req_data.GetMessageType() == int32(2) && req_data.GetElementType() == int32(1) {
		//赠送卡片
		messages.Tag = int32(4)
	} else if req_data.GetElementType() == int32(3) {
		//加好友的消息
		messages.Tag = int32(5)
	}
	for i := range req_data.GetMessagesNtf() {
		messages.Touid = req_data.GetMessagesNtf()[i].GetPlayuid()

	}
	for j := range req_data.GetElement() {
		messages.CardId = req_data.GetElement()[j].GetCardId()
		messages.Number = req_data.GetElement()[j].GetElementNum()

		id, err := o.Insert(&messages)
		if err != nil {
			beego.Error(err)
		}
		beego.Info(id)
	}
	res_data := new(cspb.CSFriendmessageRes)
	messageId := req_data.GetMessageId()
	friendListId := req_data.GetFriendListId()
	elementType := req_data.GetElementType()
	*res_data = cspb.CSFriendmessageRes{
		IsGive:       &isGive,
		Uid:          &uid,
		MessageId:    &messageId,
		FriendListId: &friendListId,
		ElementType:  &elementType,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{

		FriendmessageRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kFriendmessageRes),
		res_pkg_body, res_list)
	return ret

}
