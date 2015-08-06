package cs_handle

import (
	// "fmt"
	"github.com/astaxie/beego"
	// "time"
	models "tuojie.com/piggo/quickstart.git/models"
)
import cspb "protocol"

import proto "code.google.com/p/goprotobuf/proto"

// import db "tuojie.com/piggo/quickstart.git/db/collection"
import "labix.org/v2/mgo/bson"

import db_session "tuojie.com/piggo/quickstart.git/db/session"

// import resmgr "tuojie.com/piggo/quickstart.git/res_mgr"

func ResourceHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********ResourceHandle Start**********")
	req_data := req.GetBody().GetResourceReq()
	beego.Info(req_data)
	ret := int32(1)
	c := db_session.DB("zoo").C("player")
	var player models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player)

	var Resource []*cspb.AttrInfo

	if req_data.GetType() == int32(1) {
		_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
			bson.M{"$inc": bson.M{"diamond": req_data.GetDiamond(), "rmb": req_data.GetRMB()}})
		if err == nil {
			beego.Info("钻石数量变更成功！")
		} else {
			beego.Error("钻石数量变更失败！")
		}

	}

	if req_data.GetType() == int32(2) {
		_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
			bson.M{"$set": bson.M{"gold": req_data.GetGold(), "diamond": req_data.GetDiamond()}})
		if err == nil {
			beego.Info("金币数量变更成功！")
		} else {
			beego.Error("金币数量变更失败！")
		}
	}
	if req_data.GetType() == int32(3) {
		_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
			bson.M{"$set": bson.M{"flower": req_data.GetFlower(), "gold": req_data.GetGold()}})
		if err == nil {
			beego.Info("小红花数量变更成功！")
		} else {
			beego.Error("小红花数量变更失败！")
		}
	}
	if req_data.GetType() == int32(8) {
		_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
			bson.M{"$set": bson.M{"experience_pool": req_data.GetSolution(), "gold": req_data.GetGold()}})
		if err == nil {
			beego.Info("药水数量变更成功！")
		} else {
			beego.Error("药水花数量变更失败！")
		}
	}

	if req_data.GetType() == int32(9) {
		_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
			bson.M{"$set": bson.M{"medal": req_data.GetMedal()}})
		if err == nil {
			beego.Info("勋章变更成功！")
		} else {
			beego.Error("勋章变更失败！")
		}
	}
	if req_data.GetType() == int32(11) {
		_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
			bson.M{"$set": bson.M{"medal": req_data.GetMedal(), "diamond": req_data.GetDiamond(), "rmb": req_data.GetRMB(), "flower": req_data.GetFlower(), "experience_pool": req_data.GetSolution(), "gold": req_data.GetGold()}})
		if err == nil {
			beego.Info("信息变更成功！")
		} else {
			beego.Error("信息变更失败！")
		}
	}
	var player_return models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player_return)
	// AttrValue = append(AttrValue, makeAttrValue(player_return.Diamond, player_return.Gold, player_return.Flower, player_return.ExperiencePool))
	if req_data.GetType() == int32(10) || req_data.GetType() == int32(11) {
		AttrValue := new(cspb.AttrValue)
		*AttrValue = cspb.AttrValue{
			Diamond:  proto.Int32(player_return.Diamond),
			Gold:     proto.Int32(player_return.Gold),
			Flower:   proto.Int32(player_return.Flower),
			Solution: proto.Int32(player_return.ExperiencePool),
			Medal:    proto.Int32(player_return.Medal),
		}
		Resource = append(Resource, makeAttrInfo(int32(1), AttrValue, int32(3)))
		Resource = append(Resource, makeAttrInfo(int32(2), AttrValue, int32(3)))
		Resource = append(Resource, makeAttrInfo(int32(3), AttrValue, int32(3)))
		Resource = append(Resource, makeAttrInfo(int32(8), AttrValue, int32(3)))
		Resource = append(Resource, makeAttrInfo(int32(9), AttrValue, int32(3)))

	} else {
		AttrValue := new(cspb.AttrValue)
		*AttrValue = cspb.AttrValue{
			Diamond:  proto.Int32(player_return.Diamond),
			Gold:     proto.Int32(player_return.Gold),
			Flower:   proto.Int32(player_return.Flower),
			Solution: proto.Int32(player_return.ExperiencePool),
			Medal:    proto.Int32(player_return.Medal),
		}
		Resource = append(Resource, makeAttrInfo(req_data.GetType(), AttrValue, int32(3)))
	}
	res_data := new(cspb.CSResourceRes)
	*res_data = cspb.CSResourceRes{
		Resource: Resource,
	}
	beego.Info(res_data)
	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		ResourceRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kResourceRes),
		res_pkg_body, res_list)
	return ret

}
