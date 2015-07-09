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

func FriendStageHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********ZooHandle Start**********")
	req_data := req.GetBody().GetFriendstageReq()
	beego.Info(req_data)
	ret := int32(1)
	var Friendstage []*cspb.CSFriendStageNtf
	//测试代码
	for i := range resmgr.FriendstagetestData.GetItems() {
		Friendstage = append(Friendstage, makeFriendStagentf(resmgr.FriendstagetestData.GetItems()[i].GetPlayerId(), resmgr.FriendstagetestData.GetItems()[i].GetName(), resmgr.FriendstagetestData.GetItems()[i].GetStage()))
	}
	//正式代码
	c := db_session.DB("zoo").C("player")
	var player models.Player
	c.Find(bson.M{"c_account": res_list.GetCAccount()}).One(&player)
	for i := range player.FriendList {
		if player.FriendList[i].IsActive == int32(1) {
			Friendstage = append(Friendstage, makeFriendStagentf(player.FriendList[i].Friendid, player.FriendList[i].FriendName, GetMaxStage(player.FriendList[i].Friendid)))
		}
	}
	res_data := new(cspb.CSFriendStageRes)
	*res_data = cspb.CSFriendStageRes{
		Friendstage: Friendstage,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		FriendstageRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kFriendStageRes),
		res_pkg_body, res_list)
	return ret

}
