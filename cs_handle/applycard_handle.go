package cs_handle

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "strconv"
	// "strings"
	"time"
	models "tuojie.com/piggo/quickstart.git/models"
)
import cspb "protocol"

// import proto "code.google.com/p/goprotobuf/proto"

// import db "tuojie.com/piggo/quickstart.git/db/collection"
import "labix.org/v2/mgo/bson"

import db_session "tuojie.com/piggo/quickstart.git/db/session"

// import resmgr "tuojie.com/piggo/quickstart.git/res_mgr"

func ApplyCardHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********ApplyFlowerHandle Start**********")
	req_data := req.GetBody().GetApplycardReq()
	beego.Info(req_data)
	ret := int32(1)
	isApply := int32(0)
	c := db_session.DB("zoo").C("player")
	var player models.Player
	for i := range req_data.GetApplyFriendList() {
		var meesages_arr []models.Messages
		var cond *orm.Condition
		cond = orm.NewCondition()

		cond = cond.And("Id__gte", 1)
		cond = cond.And("Touid__contains", req_data.GetApplyFriendList()[i].GetFriendId())
		cond = cond.And("Fromuid__contains", int32(res_list.GetUid())) //req_data.GetApplyFriendList()[i].GetFriendId()
		cond = cond.And("IsFinish__contains", int32(0))
		beego.Info(cond)
		var qs orm.QuerySeter
		qs = orm.NewOrm().QueryTable("messages").SetCond(cond)

		cnt, err := qs.All(&meesages_arr)
		beego.Info(cnt, err, meesages_arr)
		if len(meesages_arr) <= 0 {
			//查找用户是否存在
			err := c.Find(bson.M{"uid": req_data.GetApplyFriendList()[i].GetFriendId()}).One(&player)
			if err != nil {
				beego.Error("找不到用户！")
			} else {
				beego.Info("找到用户")
				var players models.Player
				c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&players)
				o := orm.NewOrm()
				var messages models.Messages
				messages.Fromuid = int32(res_list.GetUid())
				messages.Fromname = players.Name
				messages.FromStarId = players.StarId
				messages.Time = time.Now().Unix()
				messages.IsFinish = int32(0)
				messages.Messagetype = int32(1)
				messages.ElementType = int32(2)
				messages.Tag = int32(3)
				messages.Touid = req_data.GetApplyFriendList()[i].GetFriendId()

				messageid, err := o.Insert(&messages)
				if err == nil {
					fmt.Println(messageid)

					for n := range req_data.GetApplyCardList() {

						_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
							bson.M{"$push": bson.M{"cardrecord": bson.M{"message_id": messageid, "card_id": req_data.GetApplyCardList()[n].GetCardId(), "card_num": req_data.GetApplyCardList()[n].GetCardNum()}}})

						if err != nil {
							beego.Error("申请赠送卡片存储失败")

						} else {
							beego.Info("申请赠送卡片存储成功")
							isApply = int32(1)
						}

					}
				}
			}
		} else {
			isApply = int32(1)
		}
	}
	res_data := new(cspb.CSApplyCardRes)
	*res_data = cspb.CSApplyCardRes{
		IsApply: &isApply,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		ApplycardRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kApplyCardRes),
		res_pkg_body, res_list)
	return ret

}
