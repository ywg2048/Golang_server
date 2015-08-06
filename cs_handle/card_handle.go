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

func CardHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********CardHandle Start**********")
	req_data := req.GetBody().GetCardReq()
	beego.Info(req_data)
	ret := int32(1)
	c := db_session.DB("zoo").C("player")
	var player models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player)

	for i := range req_data.GetCardNtf() {
		for j := range player.Cards {
			if req_data.GetCardNtf()[i].GetCardId() == player.Cards[j].CardId {
				_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
					bson.M{"$set": bson.M{"cards." + fmt.Sprint(j) + ".card_num": req_data.GetCardNtf()[i].GetCardNum()}})
				if err == nil {
					beego.Info("卡片掉落存储成功!")
				} else {
					beego.Error("卡片掉落存储失败!")
				}
			}
		}
	}

	var player_return models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player_return)
	var CardNtf []*cspb.CSCardNtf
	for i := range player_return.Cards {
		CardNtf = append(CardNtf, makecardNtf1(player_return.Cards[i].CardId, player_return.Cards[i].CardNum))
	}
	res_data := new(cspb.CSCardRes)
	*res_data = cspb.CSCardRes{
		CardNtf: CardNtf,
	}
	beego.Info(res_data)
	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		CardRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kCardRes),
		res_pkg_body, res_list)
	return ret

}
