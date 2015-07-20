package cs_handle

import (
	"fmt"
	"github.com/astaxie/beego"
	// "time"
	models "tuojie.com/piggo/quickstart.git/models"
)
import cspb "protocol"

// import proto "code.google.com/p/goprotobuf/proto"

// import db "tuojie.com/piggo/quickstart.git/db/collection"
import "labix.org/v2/mgo/bson"

import db_session "tuojie.com/piggo/quickstart.git/db/session"

// import resmgr "tuojie.com/piggo/quickstart.git/res_mgr"

func BuyCardHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********BuyCardHandle Start**********")
	req_data := req.GetBody().GetBuycardReq()
	beego.Info(req_data)
	ret := int32(1)

	c := db_session.DB("zoo").C("player")
	var player models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player)

	if req_data.GetDiamond() <= player.Diamond {
		//钻石足够
		for i := range player.Cards {
			if player.Cards[i].CardId == req_data.GetCardId() {
				_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
					bson.M{"$inc": bson.M{"diamond": -req_data.GetDiamond(), "cards." + fmt.Sprint(i) + ".card_num": req_data.GetCardNum()}})
				if err == nil {
					beego.Info("卡片购买成功！")
				} else {
					beego.Error("卡片购买失败！")
				}
			}
		}
	} else {
		beego.Error("钻石不足！")
	}

	var player_return models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player_return)
	currentDiamond := player_return.Diamond
	var CurrentCardNtf []*cspb.CSCardNtf
	for j := range player_return.Cards {
		CurrentCardNtf = append(CurrentCardNtf, makecardNtf1(player_return.Cards[j].CardId, player_return.Cards[j].CardNum))
	}
	res_data := new(cspb.CSBuyCardRes)
	*res_data = cspb.CSBuyCardRes{
		CurrentDiamond: &currentDiamond,
		CurrentCardNtf: CurrentCardNtf,
	}
	beego.Info(res_data)
	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		BuycardRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kBuyCardRes),
		res_pkg_body, res_list)
	return ret

}