package cs_handle

import (
	"fmt"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
	// "strconv"
	// "strings"
	// "time"
	models "tuojie.com/piggo/quickstart.git/models"
)
import cspb "protocol"

import proto "code.google.com/p/goprotobuf/proto"

// import db "tuojie.com/piggo/quickstart.git/db/collection"
import "labix.org/v2/mgo/bson"

import db_session "tuojie.com/piggo/quickstart.git/db/session"

// import resmgr "tuojie.com/piggo/quickstart.git/res_mgr"

func DropStageHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********DropStageHandle Start**********")
	req_data := req.GetBody().GetDropStageReq()
	beego.Info(req_data)
	ret := int32(1)
	isSuccess := bool(false)
	c := db_session.DB("zoo").C("player")
	var player models.Player
	err := c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player)
	_, errs := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
		bson.M{"$inc": bson.M{"flower": req_data.GetFlower(), "experience_pool": req_data.GetSolution(), "gold": req_data.GetGold(), "diamond": req_data.GetDiamond()}})
	if errs == nil {
		for i := range req_data.GetCardNtf() {
			for j := range player.Cards {
				if player.Cards[j].CardId == req_data.GetCardNtf()[i].GetCardId() {
					_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
						bson.M{"$inc": bson.M{"card." + fmt.Sprint(j) + ".card_num": req_data.GetCardNtf()[i].GetCardNum()}})
					if err == nil {
						beego.Info("卡片掉落存储成功！")
						isSuccess = bool(true)
					} else {
						beego.Error("卡片掉落失败！")
					}
				}
			}

		}
	} else {
		beego.Error("道具掉落更新失败！", err)
	}

	var player_return models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player_return)
	var CardArr []*cspb.CSCardNtf
	for i := range player_return.Cards {
		CardArr = append(CardArr, makecardNtf1(player_return.Cards[i].CardId, player_return.Cards[i].CardNum))
	}
	var Resource []*cspb.AttrInfo
	AttrValue := new(cspb.AttrValue)
	*AttrValue = cspb.AttrValue{
		Diamond:  proto.Int32(player_return.Diamond),
		Gold:     proto.Int32(player_return.Gold),
		Flower:   proto.Int32(player_return.Flower),
		Solution: proto.Int32(player_return.ExperiencePool),
		Medal:    proto.Int32(player_return.Medal),
	}
	Resource = append(Resource, makeAttrInfo(int32(1), AttrValue, int32(3)))
	Resource = append(Resource, makeAttrInfo(int32(2), AttrValue, int32(3)))
	Resource = append(Resource, makeAttrInfo(int32(3), AttrValue, int32(3)))
	Resource = append(Resource, makeAttrInfo(int32(8), AttrValue, int32(3)))
	Resource = append(Resource, makeAttrInfo(int32(9), AttrValue, int32(3)))

	res_data := new(cspb.CSDropStageRes)
	*res_data = cspb.CSDropStageRes{
		CardNtf:      CardArr,
		ResourceInfo: Resource,
		IsSuccess:    &isSuccess,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		DropStageRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kDropStageRes),
		res_pkg_body, res_list)
	return ret

}
