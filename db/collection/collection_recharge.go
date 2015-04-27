package collection

import (
	"github.com/astaxie/beego"
)
import "labix.org/v2/mgo"
import "labix.org/v2/mgo/bson"

import db_session "tuojie.com/piggo/quickstart.git/db/session"
import "fmt"

type RechargeFlowData struct {
	Account      string `bson:"account"`
	Uid          int64  `bson:"uid"`
	Rmb          int32  `bson:"rmb"`
	GoodsType    int32  `bson:"goods_type"`
	GoodsSubType int32  `bson:"goods_sub_type"`
	GoodsNum     int32  `bson:"goods_num"`
	RechargeTime int64  `bson:"recharge_time"`
	Version      string `bson:"version"`
	Code         string `bson:"code"`
	Channel      string `bson:"channel"`
}

func GetRechargeFlow(uid int64, start_time int64,
	end_time int64) ([]RechargeFlowData, int32) {

	beego.Debug("uid:%d, start_time:%d, end_time:%d", uid, start_time, end_time)
	c := db_session.DB("zoo").C("recharge")
	var recharge_list []RechargeFlowData
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

func AddRechargeFlow(recharge_flow *RechargeFlowData) int32 {
	beego.Debug("recharge_flow:%s", fmt.Sprint(recharge_flow))

	c := db_session.DB("zoo").C("recharge")
	err := c.Insert(&recharge_flow)
	if err != nil {
		beego.Error("insert recharge_flow fail err:%v", err)
		return -1
	}

	return 0
}
