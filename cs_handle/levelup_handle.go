package cs_handle

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "time"
	models "tuojie.com/piggo/quickstart.git/models"
)

import cspb "protocol"

import proto "code.google.com/p/goprotobuf/proto"

// import db "tuojie.com/piggo/quickstart.git/db/collection"
import "labix.org/v2/mgo/bson"

import db_session "tuojie.com/piggo/quickstart.git/db/session"
import resmgr "tuojie.com/piggo/quickstart.git/res_mgr"

func LevelUpHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********LevelUpHandle Start**********")
	req_data := req.GetBody().GetLevelUpReq()
	beego.Info(req_data)

	ret := int32(1)
	isLevelUp := int32(0)
	c := db_session.DB("zoo").C("player")
	var player models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player)
	for k := range player.Star {
		if player.Star[k].StarId == req_data.GetCurrentStarId() {
			beego.Info("找到id")
			for i := range player.Cards {
				for j := range req_data.GetCardNtf() {
					if player.Cards[i].CardId == req_data.GetCardNtf()[j].GetCardId() {
						// if player.Cards[i].CardNum >= req_data.GetCardNtf()[j].GetCardNum() {
						// 卡片足够
						for m := range resmgr.LevelupData.GetItems() {
							if resmgr.LevelupData.GetItems()[m].GetLevel() == req_data.GetCurrentLevel()-1 {
								beego.Info(resmgr.LevelupData.GetItems()[m].GetNumber())
								_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
									bson.M{"$set": bson.M{"cards." + fmt.Sprint(i) + ".card_num": req_data.GetCardNtf()[j].GetCardNum()}})
								_, errs := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
									bson.M{"$set": bson.M{"star." + fmt.Sprint(k) + ".satisfaction": resmgr.LevelupData.GetItems()[m+1].GetAttack(),
										"star." + fmt.Sprint(k) + ".fight_exp":   resmgr.LevelupData.GetItems()[m+1].GetFightExp(),
										"star." + fmt.Sprint(k) + ".current_exp": req_data.GetExp(), "star." + fmt.Sprint(k) + ".level": req_data.GetCurrentLevel()}})
								if err == nil && errs == nil {
									beego.Info("小伙伴升级成功！")

									//mysql排名表的等级变化
									o := orm.NewOrm()
									var ranking models.Ranking

									ranking = models.Ranking{Uid: int32(res_list.GetUid())}
									o.Read(&ranking, "uid")
									beego.Info(ranking)
									ranking.Level = req_data.GetCurrentLevel()
									id, errs := o.Update(&ranking, "level")
									if errs == nil {
										beego.Info("mysql排名表等级变更成功", id)
									} else {
										beego.Error("mysql排名表等级变更失败", errs)
									}

									isLevelUp = int32(1)
								} else {
									beego.Error("小伙伴升级失败！", err, errs)
								}
								break
							}
						}
						// }
					}
				}
			}
		} else {
			beego.Error("没找到id")
		}
	}

	var player_return models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player_return)

	var PetId int32
	var PetLevel int32
	var DressId int32
	var PetCurExp int32
	var Satisfaction int32
	var Fighting int32
	var FightExp int32

	var CurrentCardNtf []*cspb.CSCardNtf
	for j := range player_return.Star {
		if player_return.Star[j].StarId == req_data.GetCurrentStarId() {
			PetId = player_return.Star[j].StarId
			PetLevel = player_return.Star[j].Level
			DressId = player_return.Star[j].Dress
			PetCurExp = player_return.Star[j].Currentexp
			Satisfaction = player_return.Star[j].Satisfaction
			Fighting = player_return.Star[j].Fighting
			FightExp = player_return.Star[j].FightExp
			for i := range player_return.Cards {
				CurrentCardNtf = append(CurrentCardNtf, makecardNtf1(player_return.Cards[i].CardId, player_return.Cards[i].CardNum))
			}
		}
	}
	PetInfo := new(cspb.PetInfo)
	*PetInfo = cspb.PetInfo{
		PetId:        proto.Int32(PetId),
		PetLevel:     proto.Int32(PetLevel),
		DressId:      proto.Int32(DressId),
		PetCurExp:    proto.Int32(PetCurExp),
		Satisfaction: proto.Int32(Satisfaction),
		FightExp:     proto.Int32(FightExp),
		Fighting:     proto.Int32(Fighting),
	}
	res_data := new(cspb.CSLevelUpRes)
	*res_data = cspb.CSLevelUpRes{
		IsLevelUp:      &isLevelUp,
		PetInfo:        PetInfo,
		CurrentCardNtf: CurrentCardNtf,
	}
	beego.Info(res_data)
	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		LevelUpRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kLevelUpRes),
		res_pkg_body, res_list)
	return ret

}
