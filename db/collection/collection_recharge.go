package collection

import (
	"github.com/astaxie/beego"
	models "tuojie.com/piggo/quickstart.git/models"
)
import "labix.org/v2/mgo"
import "labix.org/v2/mgo/bson"

import db_session "tuojie.com/piggo/quickstart.git/db/session"
import "fmt"

func GetRechargeFlow(uid int64, start_time int64,
	end_time int64) ([]models.RechargeFlowData, int32) {

	beego.Debug("uid:%d, start_time:%d, end_time:%d", uid, start_time, end_time)
	c := db_session.DB("zoo").C("recharge")
	var recharge_list []models.RechargeFlowData
	err := c.Find(bson.M{"uid": uid, "recharge_time": bson.M{"$gt": start_time, "$lt": end_time}}).
		Sort("+recharge_time").All(&recharge_list)
	if err == mgo.ErrNotFound {
		beego.Error("load recharge_list no found. uid:%d, err:%v", uid, err)
		return recharge_list, 1
	} else if err != nil {
		beego.Error("load recharge_list fail err:%v", err)
		return recharge_list, -1
	}

	beego.Debug("recharge_list:%v", recharge_list)
	return recharge_list, 0
}

func AddRechargeFlow(recharge_flow *models.RechargeFlowData) int32 {
	beego.Debug("recharge_flow:%s", fmt.Sprint(recharge_flow))

	c := db_session.DB("zoo").C("recharge")
	err := c.Insert(&recharge_flow)
	if err != nil {
		beego.Error("insert recharge_flow fail err:%v", err)
		return -1
	}

	return 0
}
