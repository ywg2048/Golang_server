package cs_handle

import (
	"github.com/astaxie/beego"
	models "tuojie.com/piggo/quickstart.git/models"
)
import cspb "protocol"

// import proto "code.google.com/p/goprotobuf/proto"

// import db "tuojie.com/piggo/quickstart.git/db/collection"
import "labix.org/v2/mgo/bson"

import db_session "tuojie.com/piggo/quickstart.git/db/session"

func FriendHomeHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	//好友主页
	beego.Info("*********FriendHomeHandle Start**********")
	req_data := req.GetBody().GetFriendHomeReq()
	beego.Info(req_data)

	ret := int32(0)
	starid := 0
	c := db_session.DB("zoo").C("player")
	var player models.Player
	err := c.Find(bson.M{"c_account": req_data.GetPlayuid()}).One(&player)
	if err != nil {
		beego.Error("出错了！")
	}
	for i := range player.Star {
		if player.Star[i].IsActive == 1 {
			starid = i
			break
		}
	}
	beego.Info(starid)

	res_data := new(cspb.CSFriendHomeRes)
	*res_data = cspb.CSFriendHomeRes{
		Star:     &player.Star[starid].Starname,
		Playname: &player.Name,
		Fighting: &player.Star[starid].Fighting,
		Dress:    &player.Star[starid].Dressname,
		Level:    &player.Star[starid].Level,
		Medal:    &player.Medal,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		FriendHomeRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kFriendHomeRes),
		res_pkg_body, res_list)
	return ret

}
