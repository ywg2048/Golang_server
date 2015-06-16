package cs_handle

import (
	// "fmt"
	"github.com/astaxie/beego"
	// "time"
	// models "tuojie.com/piggo/quickstart.git/models"
)
import cspb "protocol"

import proto "code.google.com/p/goprotobuf/proto"

// import db "tuojie.com/piggo/quickstart.git/db/collection"
// import "labix.org/v2/mgo/bson"

// import db_session "tuojie.com/piggo/quickstart.git/db/session"

func AddFriendHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********AddFriendHandle Start**********")
	req_data := req.GetBody().GetAddfriendReq()
	beego.Info(req_data)
	ret := int32(1)
	uid := req_data.GetUid()
	friendId := req_data.GetFriendId()
	IsAdd := int32(1)

	c := db_session.DB("zoo").C("player")
	var player models.Player
	err := c.Find(bson.M{"c_account": res_list.GetCAccount()}).One(&player)

	res_data := new(cspb.CSAddFriendRes)
	*res_data = cspb.CSAddFriendRes{
		Ret:      proto.Int32(ret),
		Uid:      &uid,
		FriendId: &friendId,
		IsAdd:    &IsAdd,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		AddfriendRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kAddFriendRes),
		res_pkg_body, res_list)
	return ret

}
