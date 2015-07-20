package cs_handle

import (
	"fmt"
	"github.com/astaxie/beego"
	// "time"
	models "tuojie.com/piggo/quickstart.git/models"
)

import cspb "protocol"

// import proto "code.google.com/p/goprotobuf/proto"

// import db "tuojie.com/piggo/quickstart.git/db/collection"
import "labix.org/v2/mgo/bson"

import db_session "tuojie.com/piggo/quickstart.git/db/session"

func EvolutionHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********EvolutionHandle Start**********")
	req_data := req.GetBody().GetEvolutionReq()
	beego.Info(req_data)

	ret := int32(1)
	c := db_session.DB("zoo").C("player")
	var player models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player)

	for i := range player.Star {
		if player.Star[i].StarId == req_data.GetStarId() {
			if player.Star[i].Currentexp == req_data.GetCurrentExp() {
				//检测客户端的经验是否和服务器一致
				if player.ExperiencePool >= req_data.GetSolution() {
					//检测药水的使用量是否合法
					_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
						bson.M{"$inc": bson.M{"experience_pool": -req_data.GetSolution(), "star." + fmt.Sprint(i) + ".current_exp": req_data.GetSolution()}})
					if err == nil {
						beego.Info("使用药水成功！")
					} else {
						beego.Error("使用药水失败！", err)
					}
				} else {
					beego.Error("药水使用量不合法！")
				}
			} else {
				beego.Error("客户端和服务器的经验值不一致！")
			}
		}
	}
	//返回值
	var player_return models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player_return)
	var starId int32
	var solution int32
	var currentExp int32

	starId = req_data.GetStarId()
	solution = player_return.ExperiencePool
	for j := range player_return.Star {
		if player_return.Star[j].StarId == req_data.GetStarId() {
			currentExp = player_return.Star[j].Currentexp
		}
	}
	res_data := new(cspb.CSEvolutionRes)
	*res_data = cspb.CSEvolutionRes{
		StarId:     &starId,
		Solution:   &solution,
		CurrentExp: &currentExp,
	}
	beego.Info(res_data)
	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		EvolutionRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kCSEvolutionReq),
		res_pkg_body, res_list)
	return ret

}
