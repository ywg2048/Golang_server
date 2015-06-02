package cs_handle

import (
	"github.com/astaxie/beego"
	models "tuojie.com/piggo/quickstart.git/models"
)
import cspb "protocol"
import proto "code.google.com/p/goprotobuf/proto"

// import db "tuojie.com/piggo/quickstart.git/db/collection"
import "labix.org/v2/mgo/bson"

import db_session "tuojie.com/piggo/quickstart.git/db/session"

func ApplyListHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********ApplyListHandle Start**********")
	req_data := req.GetBody().GetApplylistReq()
	beego.Info(req_data)
	c := db_session.DB("zoo").C("player")
	var player models.Player
	err := c.Find(bson.M{"c_account": res_list.GetCAccount()}).One(&player)
	beego.Debug("*********ApplyListHandle result is*******", player, err)

	ret := int32(1)
	var res_applylist []*cspb.CSApplyListNtf
	for i := range player.ApplyFriendList {
		if player.ApplyFriendList[i].IsAccept == 0 && player.ApplyFriendList[i].Isrefuse == 0 {
			res_applylist = append(res_applylist, makeApplylist(player.ApplyFriendList[i].Applyuid, player.ApplyFriendList[i].IsAccept, player.ApplyFriendList[i].Isrefuse, player.ApplyFriendList[i].Applytime, player.ApplyFriendList[i].Oprationtime))
		}
	}
	//拉取申请列表
	res_data := new(cspb.CSApplyListRes)
	*res_data = cspb.CSApplyListRes{
		Ret:          proto.Int32(ret),
		ApplyListNtf: res_applylist,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		ApplylistRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kApplyListRes),
		res_pkg_body, res_list)
	return ret

}
