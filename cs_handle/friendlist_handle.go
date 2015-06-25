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
import resmgr "tuojie.com/piggo/quickstart.git/res_mgr"

func FriendlistHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********FriendmessageHandle Start**********")
	req_data := req.GetBody().GetFriendlistReq()
	beego.Info(req_data)
	ret := int32(1)

	var FriendListNtf []*cspb.CSFriendListNtf
	//测试代码
	for i := range resmgr.FriendlisttestData.GetItems() {
		FriendListNtf = append(FriendListNtf, makefriendlist(resmgr.FriendlisttestData.GetItems()[i].GetID(), resmgr.FriendlisttestData.GetItems()[i].GetUID(), resmgr.FriendlisttestData.GetItems()[i].GetName(), resmgr.FriendlisttestData.GetItems()[i].GetStarId(), resmgr.FriendlisttestData.GetItems()[i].GetStarName(),
			resmgr.FriendlisttestData.GetItems()[i].GetFighting(), resmgr.FriendlisttestData.GetItems()[i].GetDressID(), resmgr.FriendlisttestData.GetItems()[i].GetDressName(), resmgr.FriendlisttestData.GetItems()[i].GetLevel(),
			resmgr.FriendlisttestData.GetItems()[i].GetMedal(), resmgr.FriendlisttestData.GetItems()[i].GetMedalLevelId(), resmgr.FriendlisttestData.GetItems()[i].GetStagelevel()))
	}
	//正式代码
	c := db_session.DB("zoo").C("player")
	var player models.Player
	c.Find(bson.M{"c_account": res_list.GetCAccount()}).One(&player)
	for i := range player.FriendList {

		var players models.Player
		err_ := c.Find(bson.M{"uid": player.FriendList[i].Friendid}).One(&players)
		if err_ != nil {
			beego.Error(err_)
		}
		//查找当前明星的信息,先给默认
		StarName := resmgr.FriendlisttestData.GetItems()[1].GetStarName()
		Fighting := resmgr.FriendlisttestData.GetItems()[1].GetFighting()
		DressId := resmgr.FriendlisttestData.GetItems()[1].GetDressID()
		Dress := resmgr.FriendlisttestData.GetItems()[1].GetDressName()
		Level := resmgr.FriendlisttestData.GetItems()[1].GetLevel()
		Medal := resmgr.FriendlisttestData.GetItems()[1].GetMedal()
		MedalLevelID := resmgr.FriendlisttestData.GetItems()[1].GetMedalLevelId()
		Stagelevel := resmgr.FriendlisttestData.GetItems()[i].GetStagelevel() //最大关卡有待商定
		for j := range players.Star {
			if players.Star[j].StarId == players.StarId {
				StarName = players.Star[j].Starname
				Fighting = players.Star[j].Fighting
				DressId = players.Star[j].Dress
				Dress = players.Star[j].Dressname
				Level = players.Star[j].Level
				Medal = players.Star[j].Medal
				MedalLevelID = players.Star[j].MedalLevelId

				break
			}
		}

		FriendListNtf = append(FriendListNtf, makefriendlist(int32(i), player.FriendList[i].Friendid, players.Name, players.StarId, StarName, Fighting, DressId, Dress, Level, Medal, MedalLevelID, Stagelevel))
		beego.Info("FriendListNtf is:", FriendListNtf)
	}
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
