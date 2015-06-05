package cs_handle

import (
	"github.com/astaxie/beego"
	// models "tuojie.com/piggo/quickstart.git/models"
)
import cspb "protocol"

// import proto "code.google.com/p/goprotobuf/proto"

// import db "tuojie.com/piggo/quickstart.git/db/collection"
// import "labix.org/v2/mgo/bson"

// import db_session "tuojie.com/piggo/quickstart.git/db/session"

func HomePageHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	//好友主页
	beego.Info("*********HomePageHandle Start**********")
	req_data := req.GetBody().GetHomepageReq()
	beego.Info(req_data)

	ret := int32(1)
	// starid := 0
	// c := db_session.DB("zoo").C("player")
	// var player models.Player
	// err := c.Find(bson.M{"c_account": req_data.GetPlayuid()}).One(&player)
	// if err != nil {
	// 	beego.Error("出错了！")
	// }
	// for i := range player.Star {
	// 	if player.Star[i].IsActive == 1 {
	// 		starid = i
	// 		break
	// 	}
	// }
	// beego.Info(starid)
	// required int32 starid = 2;
	// required string dress = 3;
	// required int32 level = 4;
	// required int32 medal = 5;
	// required int32 exp =6;
	// required int32 satisfaction = 7;
	starid := int32(1)
	dress := "演唱会"
	level := int32(12)
	medal := int32(931)
	exp := int32(150)
	satisfaction := int32(250)

	res_data := new(cspb.CSHomePageRes)
	*res_data = cspb.CSHomePageRes{
		Starid:       &starid,
		Dress:        &dress,
		Level:        &level,
		Medal:        &medal,
		Exp:          &exp,
		Satisfaction: &satisfaction,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		HomepageRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kHomePageRes),
		res_pkg_body, res_list)
	return ret

}
