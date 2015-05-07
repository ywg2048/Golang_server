package cs_handle

import (
	"github.com/astaxie/beego"
)
import cspb "protocol"
import proto "code.google.com/p/goprotobuf/proto"

import db "tuojie.com/piggo/quickstart.git/db/collection"
import "labix.org/v2/mgo/bson"
import db_session "tuojie.com/piggo/quickstart.git/db/session"
import (
	"fmt"
	"time"
)

func stageReportHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {

	// log.Debug("******StageReportHandle, req is %v, res is %v", req, res_list)
	// ret, player := db.LoadPlayer(res_list.GetCAccount(), res_list.GetSAccount(), res_list.GetUid())
	// log.Debug("Db LoadPlayer Player, ret is %d, player is %v", ret, player)
	req_data := req.GetBody().GetStageReportReq()

	c := db_session.DB("zoo").C("player")
	var player db.Player
	err := c.Find(bson.M{"c_account": res_list.GetCAccount()}).One(&player)
	beego.Debug("*********StageReportHandle result is %v err is %v********", player, err)
	//log.Debug("*********StageReportHandle result.level1 is %v********", player.Levels[1].GetStageScore())
	//时间戳
	ret := int32(0)
	if err == nil {
		for i := range req_data.StageNtf {
			var stage = req_data.StageNtf[i]
			if len(player.Levels) > 0 {
				if len(req_data.StageNtf) <= len(player.Levels) {
					if int32(*stage.StageScore) > int32(player.Levels[i].GetStageScore()) {
						c.Upsert(bson.M{"c_account": res_list.GetCAccount()},
							bson.M{"$set": bson.M{"Levels." + fmt.Sprint(i): stage}})
						c.Upsert(bson.M{"c_account": res_list.GetCAccount()},
							bson.M{"$set": bson.M{"Levels." + fmt.Sprint(i) + ".timestamp": time.Now().Unix()}})
						beego.Info("更新用户分数成功")
					}
				} else {
					c.Upsert(bson.M{"c_account": res_list.GetCAccount()},
						bson.M{"$set": bson.M{"Levels." + fmt.Sprint(i): stage}})
					c.Upsert(bson.M{"c_account": res_list.GetCAccount()},
						bson.M{"$set": bson.M{"Levels." + fmt.Sprint(i) + ".timestamp": time.Now().Unix()}})
					beego.Info("新关卡用户分数保存成功")
				}

			} else {
				//*stage.StageLevel = int32(1)
				c.Upsert(bson.M{"c_account": res_list.GetCAccount()},
					bson.M{"$set": bson.M{"Levels." + fmt.Sprint(i): stage}})
				c.Upsert(bson.M{"c_account": res_list.GetCAccount()},
					bson.M{"$set": bson.M{"Levels." + fmt.Sprint(i) + ".timestamp": time.Now().Unix()}})
				beego.Info("插入新用户分数成功")
			}

			// for j:= range player.Levels {
			// 	if *stage.StageId == player.Levels[j].GetStageId() {
			// 	}
			// }

			// if int32(*stage.StageScore) > int32(player.Levels[i].GetStageScore()) {

			// }
			ret = 1
		}

	} else {
		ret = 0
	}

	res_data := new(cspb.CSStageReportRes)
	*res_data = cspb.CSStageReportRes{
		Ret:      proto.Int32(ret),
		StageNtf: req_data.StageNtf,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		StageReportRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kStageReportRes),
		res_pkg_body, res_list)
	return ret

	// //查询的时间戳
	// log.Debug("*********DB TimeStamp is %v", result)

	// //时间戳
	// Timestamp := time.Now().Unix()
	// _, err := c.Upsert(bson.M{"c_account": res_list.GetCAccount()},
	// 	bson.M{"$set": bson.M{"Levels": req_data.StageNtf, "timestamp": Timestamp}})

	// if err != nil {
	// 	log.Error("插入关数失败:%v", err)
	// 	ret = int32(cspb.ErrorCode_PlayerInsertFail)
	// 	player.Uid = int64(0)
	// }

	//log.Debug("******StageReportHandle, levelid is %v, level is %v,score is %v,res is %v", levelid, stage_level, stage_score, res)
	// return makeStageReportResPkg(req, res_list, ret)
}
