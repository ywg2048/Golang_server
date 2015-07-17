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

import resmgr "tuojie.com/piggo/quickstart.git/res_mgr"

func LevelUpHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********LevelUpHandle Start**********")
	req_data := req.GetBody().GetLevelUpReq()
	beego.Info(req_data)
	// beego.Info(resmgr.LevelupData.GetItems())
	ret := int32(1)

	c := db_session.DB("zoo").C("player")
	var player models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player)
	for k := range player.Star {
		//判定等级是否真实
		if player.Star[k].StarId == req_data.GetCurrentStarId() {
			if player.Star[k].Level == req_data.GetCurrentLevel()-int32(1) {
				//等级符合要求
				for a := range resmgr.ExpData.GetItems() {
					if resmgr.ExpData.GetItems()[a].GetLevelId() == player.Star[k].Level {
						if (player.Star[k].Currentexp - resmgr.ExpData.GetItems()[a].GetRequireExp()) == req_data.GetExp() {
							//经验足够的情况下
							for i := range player.Cards {
								for j := range req_data.GetCardNtf() {
									if player.Cards[i].CardId == req_data.GetCardNtf()[j].GetCardId() {
										if player.Cards[i].CardNum >= req_data.GetCardNtf()[j].GetCardNum() {
											// 卡片足够
											for m := range resmgr.LevelupData.GetItems() {
												if resmgr.LevelupData.GetItems()[m].GetLevel() == req_data.GetCurrentLevel() {

													_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
														bson.M{"$inc": bson.M{"cards." + fmt.Sprint(i) + ".card_num": -req_data.GetCardNtf()[j].GetCardNum(), "star." + fmt.Sprint(k) + ".level": int32(1)}})
													_, errs := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
														bson.M{"$set": bson.M{"star." + fmt.Sprint(k) + ".satisfaction": resmgr.LevelupData.GetItems()[m].GetSatisfaction(),
															"star." + fmt.Sprint(k) + ".fight_exp":   resmgr.LevelupData.GetItems()[m].GetFightExp(),
															"star." + fmt.Sprint(k) + ".current_exp": req_data.GetExp()}})
													if err == nil && errs == nil {
														beego.Info("小伙伴升级成功！")
													} else {
														beego.Error("小伙伴升级失败！", err, errs)
													}
												}
											}
										} else {
											beego.Error("卡片数量不足")
										}
									}
								}
							}
						} else {
							beego.Error("经验不足以升级")
						}
					}
				}
			} else {
				beego.Error("等级错误！")
			}
		}
	}
	// var player_after models.Player
	// c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player_after)
	// //返回值
	// currentStarId := req_data.GetCurrentStarId()
	// var currentLevel int32
	// currentSolution := player_after.ExperiencePool
	// // CurrentExp :=
	// for i := range player_after.Star {
	// 	if player_after.Star[i].StarId == req_data.GetCurrentStarId() {
	// 		currentLevel = player_after.Star[i].Level

	// 	}
	// }
	res_data := new(cspb.CSLevelUpRes)
	*res_data = cspb.CSLevelUpRes{
	// PetInfo:        PetInfo,
	// CurrentCardNtf: CurrentCardNtf,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		LevelUpRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kLevelUpRes),
		res_pkg_body, res_list)
	return ret

}
