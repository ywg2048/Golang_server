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

// import resmgr "tuojie.com/piggo/quickstart.git/res_mgr"

func ZooAnimalHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********ZooAnimalHandle Start**********")
	req_data := req.GetBody().GetZooanimalReq()
	beego.Info(req_data)
	ret := int32(1)

	c := db_session.DB("zoo").C("player")
	var player models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player)

	if req_data.GetIsLocked() == int32(1) {
		//解锁动物
		if req_data.GetUptolevel() == int32(1) {
			//检测动物是否为1级
			for i := range player.Zoo {
				if player.Zoo[i].AnimalId == req_data.GetAnimalId() {
					beego.Error("此动物已经解锁了，不要再次解锁")
				} else {
					_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
						bson.M{"$push": bson.M{"zoo": bson.M{"animal_id": req_data.GetAnimalId(), "animal_level": int32(1), "is_locked": int32(0)}}})
					if err == nil {
						beego.Info("解锁成功！")
					} else {
						beego.Error("解锁失败！")
					}
				}
			}
		} else {
			beego.Error("新动物不为1级")
		}
	} else if req_data.GetIsLocked() == int32(0) {

		//升级动物
		for i := range player.Zoo {
			if req_data.GetAnimalId() == player.Zoo[i].AnimalId {
				if player.Zoo[i].AnimalLevel == req_data.GetUptolevel()-int32(1) {
					//检查等级是否合法
					if player.Gold >= req_data.GetRequiredGold() {
						//检查金币是否足够
						_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
							bson.M{"$inc": bson.M{"zoo." + fmt.Sprint(i) + ".animal_level": int32(1), "gold": -req_data.GetRequiredGold()}})
						if err == nil {
							beego.Info("动物升级成功！")
						} else {
							beego.Error("动物升级失败！")
						}
					} else {
						beego.Error("金币不够")
					}
				} else {
					beego.Error("等级错误！")
				}
			}
		}
	}
	res_data := new(cspb.CSZooAnimalRes)

	*res_data = cspb.CSZooAnimalRes{}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		ZooanimalRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kZooRes),
		res_pkg_body, res_list)
	return ret

}
