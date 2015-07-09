package cs_handle

import (
	// "fmt"
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

func FriendlistHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********FriendmessageHandle Start**********")
	// req_data := req.GetBody().GetFriendlistReq()
	// beego.Info(req_data)
	ret := int32(1)

	var FriendListNtf []*cspb.CSFriendListNtf
	// 测试代码
	// for i := range resmgr.FriendlisttestData.GetItems() {
	// 	FriendListNtf = append(FriendListNtf, makefriendlist(resmgr.FriendlisttestData.GetItems()[i].GetID(), resmgr.FriendlisttestData.GetItems()[i].GetUID(), resmgr.FriendlisttestData.GetItems()[i].GetName(), resmgr.FriendlisttestData.GetItems()[i].GetStarId(), resmgr.FriendlisttestData.GetItems()[i].GetStarName(),
	// 		resmgr.FriendlisttestData.GetItems()[i].GetFighting(), resmgr.FriendlisttestData.GetItems()[i].GetDressID(), resmgr.FriendlisttestData.GetItems()[i].GetDressName(), resmgr.FriendlisttestData.GetItems()[i].GetLevel(),
	// 		resmgr.FriendlisttestData.GetItems()[i].GetMedal(), resmgr.FriendlisttestData.GetItems()[i].GetMedalLevelId(), resmgr.FriendlisttestData.GetItems()[i].GetStagelevel()))
	// }
	//正式代码
	c := db_session.DB("zoo").C("player")
	var player models.Player
	c.Find(bson.M{"c_account": res_list.GetCAccount()}).One(&player)
	beego.Info(player.FriendList[0], player.FriendList[1])

	var StarName string
	var Fighting int32
	var DressId int32
	var Dress string
	var Level int32
	var Medal int32
	var MedalLevelID int32
	var Stagelevel int32
	for i := range player.FriendList {

		var players models.Player
		err_ := c.Find(bson.M{"uid": player.FriendList[i].Friendid}).One(&players)
		beego.Info(players.FriendList[0], players.FriendList[1])
		if err_ != nil {
			beego.Error(err_)
		}
		//定义

		Stagelevel = GetMaxStage(player.FriendList[i].Friendid)
		for j := range players.Star {
			if players.Star[j].StarId == players.StarId {
				beego.Info("明星选择")
				StarName = players.Star[j].Starname
				Fighting = players.Star[j].Fighting
				DressId = players.Star[j].Dress
				Dress = players.Star[j].Dressname
				Level = players.Star[j].Level
				Medal = players.Star[j].Medal
				MedalLevelID = players.Star[j].MedalLevelId

			}
		}

		FriendListNtf = append(FriendListNtf, makefriendlist(int32(i), player.FriendList[i].Friendid, players.Name, players.StarId, StarName, Fighting, DressId, Dress, Level, Medal, MedalLevelID, Stagelevel))
		beego.Info("FriendListNtf is:", FriendListNtf)
	}
	beego.Info("FriendListNtf is:", FriendListNtf)
	res_data := new(cspb.CSFriendlistRes)
	*res_data = cspb.CSFriendlistRes{
		FriendListNtf: FriendListNtf,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		FriendlistRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kFriendlistRes),
		res_pkg_body, res_list)
	return ret

}
func GetMaxStage(uid int32) int32 {
	//查找最大关卡数
	MaxStage := int32(0)
	c := db_session.DB("zoo").C("player")
	var player models.Player
	err := c.Find(bson.M{"uid": uid}).One(&player)
	if err != nil {
		beego.Error("没有这样的玩家！", err)
	}
	arr := [...]int32{0}
	s := arr[0:]
	for i := range player.Levels {
		if player.Levels[i].GetStageId() < 10000 {
			//普通关卡
			s = append(s, player.Levels[i].GetStageId())
			MaxStage = int32(len(s)) - int32(1)
		}
	}
	return MaxStage
}
