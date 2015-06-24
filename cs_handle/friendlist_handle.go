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
	c := db_session.DB("zoo").C("pet")
	var player models.Player
	c.Find(bson.M{"c_account": res_list.GetCAccount()}).One(&player)
	for i := range player.FriendList {
		var pet models.Pet
		err := c.Find(bson.M{"account": res_list.GetSAccount()}).One(&pet)
		if err != nil {
			beego.Error(err)
		}

		var players models.Player
		err_ := c.Find(bson.M{"uid": player.FriendList[i].Friendid}).One(&players)
		FriendListNtf = append(FriendListNtf, makefriendlist(int32(i), player.FriendList[i].Friendid, players.Name, players.Starid, "春春", Fighting, DressId, Dress, Level, Medal, MedalLevelID, Stagelevel))

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
