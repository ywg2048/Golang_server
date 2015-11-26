package cs_handle

import (
	"fmt"
	"github.com/astaxie/beego"
	// "time"
	models "tuojie.com/piggo/quickstart.git/models"
)
import cspb "protocol"

import proto "code.google.com/p/goprotobuf/proto"

// import db "tuojie.com/piggo/quickstart.git/db/collection"
import "labix.org/v2/mgo/bson"

import db_session "tuojie.com/piggo/quickstart.git/db/session"

import resmgr "tuojie.com/piggo/quickstart.git/res_mgr"

func ReceiveRewardHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********ReceiveRewardHandle Start**********")
	req_data := req.GetBody().GetReceiverewardReq()
	beego.Info(req_data)
	ret := int32(1)
	IsReceive := int32(0)
	c := db_session.DB("zoo").C("player")
	var player models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player)
	for i := range player.Achievement {
		if player.Achievement[i].AchievementId == req_data.GetAchievementid() {
			if player.Achievement[i].StarLevel <= req_data.GetStarLevel() {
				//只有没有领取的时候才能领取
				for j := range resmgr.AchievementData.GetItems() {
					if resmgr.AchievementData.GetItems()[j].GetId() == req_data.GetAchievementid() && resmgr.AchievementData.GetItems()[j].GetStarnum() == req_data.GetStarLevel() {
						_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
							bson.M{"$inc": bson.M{"achievement." + fmt.Sprint(i) + ".starlevel": 1, "Gold": resmgr.AchievementData.GetItems()[j].GetGold(),
								"diamond": resmgr.AchievementData.GetItems()[j].GetDiamond(), "flower": resmgr.AchievementData.GetItems()[j].GetFlower(),
								"experience_pool": resmgr.AchievementData.GetItems()[j].GetSolution(), "medal": resmgr.AchievementData.GetItems()[j].GetMedal()}})

						if err == nil {
							_, errs := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
								bson.M{"$set": bson.M{"achievement." + fmt.Sprint(i) + ".isreceive": int32(1)}})
							if errs == nil {
								beego.Info("领取奖励成功！")
								IsReceive = int32(1)
							} else {
								beego.Error("isreceive变更失败")
							}
						} else {
							beego.Error("领取奖励失败！", err)
						}
					}
				}

			} else {
				beego.Error("该奖励已经领取！")
			}
		}
	}

	//返回资源
	var player_return models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player_return)
	// AttrValue = append(AttrValue, makeAttrValue(player_return.Diamond, player_return.Gold, player_return.Flower, player_return.ExperiencePool))
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

	Achievementid := req_data.GetAchievementid()
	StarLevel := req_data.GetStarLevel()
	beego.Info(IsReceive)
	res_data := new(cspb.CSReceiveRewardRes)
	*res_data = cspb.CSReceiveRewardRes{
		IsReceive:     &IsReceive,
		Achievementid: &Achievementid,
		StarLevel:     &StarLevel,
		ResouceInfo:   Resource,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		ReceiverewardRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kReceiveRewardRes),
		res_pkg_body, res_list)
	return ret

}
