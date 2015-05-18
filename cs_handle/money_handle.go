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

func MoneyHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********MoneyHandle Start**********")
	req_data := req.GetBody().GetMoneyReq()
	beego.Info(req_data)
	c := db_session.DB("zoo").C("player")
	var player models.Player
	err := c.Find(bson.M{"c_account": res_list.GetCAccount()}).One(&player)
	beego.Debug("*********MoneyHandle result is %v err is %v********", player, err)

	ret := int32(0)

	var money = req_data
	if player.Money != nil {

		if player.Money.GetTime() <= *req_data.Time {
			//服务器时间比客户端旧就更新,服务器时间比客户端新就不做任何动作
			c.Upsert(bson.M{"c_account": res_list.GetCAccount()},
				bson.M{"$set": bson.M{"money": money}})
		}
		ret = int32(1)
	} else {
		//服务器没有记录，就把客户端的数据保存
		c.Upsert(bson.M{"c_account": res_list.GetCAccount()},
			bson.M{"$set": bson.M{"money": money}})
		ret = int32(1)
	}

	var moneyres models.Player
	errs := c.Find(bson.M{"c_account": res_list.GetCAccount()}).One(&moneyres)
	if errs != nil {
		beego.Error("查询money失败")
	}
	res_data := new(cspb.CSMoneyRes)
	*res_data = cspb.CSMoneyRes{
		Ret:      proto.Int32(ret),
		Diamond:  moneyres.Money.Diamond,
		Goldcoin: moneyres.Money.Goldcoin,
		Power:    moneyres.Money.Power,
		Time:     moneyres.Money.Time,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		MoneyRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kMoneyRes),
		res_pkg_body, res_list)
	return ret

}
