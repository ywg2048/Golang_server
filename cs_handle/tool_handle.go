package cs_handle

import (
	"fmt"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
	// "strconv"
	// "strings"
	// "time"
	models "tuojie.com/piggo/quickstart.git/models"
)
import cspb "protocol"

import proto "code.google.com/p/goprotobuf/proto"

// import db "tuojie.com/piggo/quickstart.git/db/collection"
import "labix.org/v2/mgo/bson"

import db_session "tuojie.com/piggo/quickstart.git/db/session"

import resmgr "tuojie.com/piggo/quickstart.git/res_mgr"

func ToolHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********ToolHandle Start**********")
	req_data := req.GetBody().GetToolReq()
	beego.Info(req_data)
	ret := int32(1)
	isSuccess := bool(false)
	c := db_session.DB("zoo").C("player")
	var player models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player)
	if len(player.Tool) == 0 {
		//Tool初始化
		for i := range resmgr.ToolConfigData.GetItems() {
			_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
				bson.M{"$push": bson.M{"tool": bson.M{"tool_id": resmgr.ToolConfigData.GetItems()[i].GetId(), "tool_count": int32(0), "tool_name": resmgr.ToolConfigData.GetItems()[i].GetSzName(), "tool_pricetype": resmgr.ToolConfigData.GetItems()[i].GetIToolPriceType()}}})
			if err != nil {
				beego.Error("道具初始化失败")
			}
		}

	}

	for i := range req_data.GetToolInfo() {
		for j := range player.Tool {
			if player.Tool[j].ToolId == req_data.GetToolInfo()[i].GetToolID() {
				beego.Info("找到ID", req_data.GetToolInfo()[i].GetToolID())
				if req_data.GetToolInfo()[i].GetToolPriceType() == 2 {
					//购买道具(砖石)

					_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
						bson.M{"$inc": bson.M{"tool." + fmt.Sprint(j) + ".tool_count": req_data.GetToolInfo()[i].GetToolCount(), "diamond": -req_data.GetToolInfo()[i].GetToolPrice()}})
					if err != nil {
						beego.Error("道具变更失败")
					} else {
						beego.Info("道具变更成功")
						isSuccess = true
					}
				} else if req_data.GetToolInfo()[i].GetToolPriceType() == 1 {
					_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
						bson.M{"$inc": bson.M{"tool." + fmt.Sprint(j) + ".tool_count": req_data.GetToolInfo()[i].GetToolCount(), "gold": -req_data.GetToolInfo()[i].GetToolPrice()}})
					if err != nil {
						beego.Error("道具变更失败")
					} else {
						beego.Info("道具变更成功")
						isSuccess = true
					}
				} else if req_data.GetToolInfo()[i].GetToolPriceType() == 3 {
					_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
						bson.M{"$inc": bson.M{"tool." + fmt.Sprint(j) + ".tool_count": req_data.GetToolInfo()[i].GetToolCount()}})
					if err != nil {
						beego.Error("道具变更失败")
					} else {
						beego.Info("道具变更成功")
						isSuccess = true
					}
				}
			} else {
				beego.Error("没有找到ID", req_data.GetToolInfo()[i].GetToolID())
			}

		}

	}
	//回包
	var player_return models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player_return)
	var Resource []*cspb.AttrInfo
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

	var ToolInfo []*cspb.CSToolInfo
	for i := range player_return.Tool {
		for j := range req_data.GetToolInfo() {
			if player_return.Tool[i].ToolId == req_data.GetToolInfo()[j].GetToolID() {
				ToolInfo = append(ToolInfo, makeToolInfo(player_return.Tool[i].ToolId, player_return.Tool[i].ToolName, player_return.Tool[i].ToolCount, req_data.GetToolInfo()[j].GetToolPriceType(), req_data.GetToolInfo()[j].GetToolPrice()))
			}
		}
	}

	res_data := new(cspb.CSToolRes)
	*res_data = cspb.CSToolRes{
		ToolInfo:     ToolInfo,
		ResourceInfo: Resource,
		IsSuccess:    &isSuccess,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		ToolRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kToolRes),
		res_pkg_body, res_list)
	return ret

}
