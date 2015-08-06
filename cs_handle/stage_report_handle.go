package cs_handle

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	models "tuojie.com/piggo/quickstart.git/models"
)
import cspb "protocol"
import proto "code.google.com/p/goprotobuf/proto"

// import db "tuojie.com/piggo/quickstart.git/db/collection"
import "labix.org/v2/mgo/bson"
import db_session "tuojie.com/piggo/quickstart.git/db/session"
import (
	"fmt"
	"time"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:@/Monsters?charset=utf8")
}
func stageReportHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {

	// log.Debug("******StageReportHandle, req is %v, res is %v", req, res_list)
	// ret, player := db.LoadPlayer(res_list.GetCAccount(), res_list.GetSAccount(), res_list.GetUid())
	// log.Debug("Db LoadPlayer Player, ret is %d, player is %v", ret, player)
	req_data := req.GetBody().GetStageReportReq()

	c := db_session.DB("zoo").C("player")
	var player models.Player
	err := c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player)
	beego.Debug("*********StageReportHandle result is %v err is %v********", player, err)

	//时间戳
	ret := int32(0)
	if err == nil {
		for i := range req_data.GetStageNtf() {

			if len(player.Levels) > 0 {
				if len(req_data.GetStageNtf()) <= len(player.Levels) {
					if int32(req_data.GetStageNtf()[i].GetStageScore()) > int32(player.Levels[i].StageScore) {
						_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
							bson.M{"$set": bson.M{"Levels." + fmt.Sprint(i) + ".stage_id": req_data.GetStageNtf()[i].GetStageId(),
								"Levels." + fmt.Sprint(i) + ".stage_level": req_data.GetStageNtf()[i].GetStageLevel(),
								"Levels." + fmt.Sprint(i) + ".stage_score": req_data.GetStageNtf()[i].GetStageScore(),
								"Levels." + fmt.Sprint(i) + ".medal_isadd": req_data.GetStageNtf()[i].GetMedalIsAdd(),
								"Levels." + fmt.Sprint(i) + ".time_stamp":  time.Now().Unix()}})
						if err == nil {
							beego.Info("更新用户分数成功")
						}
					}
				} else {
					_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
						bson.M{"$set": bson.M{"Levels." + fmt.Sprint(i) + ".stage_id": req_data.GetStageNtf()[i].GetStageId(),
							"Levels." + fmt.Sprint(i) + ".stage_level": req_data.GetStageNtf()[i].GetStageLevel(),
							"Levels." + fmt.Sprint(i) + ".stage_score": req_data.GetStageNtf()[i].GetStageScore(),
							"Levels." + fmt.Sprint(i) + ".medal_isadd": req_data.GetStageNtf()[i].GetMedalIsAdd(),
							"Levels." + fmt.Sprint(i) + ".time_stamp":  time.Now().Unix()}})
					if err == nil {
						beego.Info("新关卡用户分数保存成功")
					}
				}

			} else {
				//*stage.StageLevel = int32(1)
				_, err := c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
					bson.M{"$set": bson.M{"Levels." + fmt.Sprint(i) + ".stage_id": req_data.GetStageNtf()[i].GetStageId(),
						"Levels." + fmt.Sprint(i) + ".stage_level": req_data.GetStageNtf()[i].GetStageLevel(),
						"Levels." + fmt.Sprint(i) + ".stage_score": req_data.GetStageNtf()[i].GetStageScore(),
						"Levels." + fmt.Sprint(i) + ".medal_isadd": req_data.GetStageNtf()[i].GetMedalIsAdd(),
						"Levels." + fmt.Sprint(i) + ".time_stamp":  time.Now().Unix()}})
				if err == nil {
					beego.Info("插入新用户分数成功")
				}
			}

			ret = 1
		}
	}

	//存储卡片

	//取出分数放入mysql里面
	// var playerscore models.Player
	// errs := c.Find(bson.M{"c_account": res_list.GetCAccount()}).One(&playerscore)

	// o := orm.NewOrm()
	// var userscore models.Userscore

	// var userscore_read []models.Userscore
	// var cond *orm.Condition
	// cond = orm.NewCondition()

	// cond = cond.And("Uid__contains", int64(playerscore.Uid))
	// var qs orm.QuerySeter
	// qs = orm.NewOrm().QueryTable("userscore").SetCond(cond)
	// cnt, err := qs.All(&userscore_read)
	// beego.Info(cnt, err)
	// if errs == nil {
	// 	if len(userscore_read) == 0 {
	// 		//如何mysql中无数据
	// 		for j := range playerscore.Levels {
	// 			userscore.Uid = int64(playerscore.Uid)
	// 			userscore.Level = playerscore.Levels[j].GetStageId()
	// 			userscore.Score = playerscore.Levels[j].GetStageScore()
	// 			userscore.Startnum = playerscore.Levels[j].GetStageLevel()
	// 			userscore.Time = int64(playerscore.Levels[j].GetTimestamp())
	// 			id, err := o.Insert(&userscore)
	// 			if err == nil {

	// 				beego.Debug("插入成功！！", id)

	// 			} else {

	// 				beego.Error("插入失败！！！")

	// 			}
	// 		}
	// 	} else {
	// 		//如果mysql中有数据
	// 		if len(userscore_read) == len(playerscore.Levels) {
	// 			//没有添加新关数
	// 			for j := range playerscore.Levels {
	// 				userscore.Id = userscore_read[j].Id
	// 				userscore.Uid = int64(playerscore.Uid)
	// 				userscore.Level = playerscore.Levels[j].GetStageId()
	// 				userscore.Score = playerscore.Levels[j].GetStageScore()
	// 				userscore.Startnum = playerscore.Levels[j].GetStageLevel()
	// 				userscore.Time = int64(playerscore.Levels[j].GetTimestamp())

	// 				if num, err := o.Update(&userscore); err == nil {

	// 					beego.Debug("更新成功！！", num)

	// 				} else {

	// 					beego.Error("更新失败！！！")

	// 				}
	// 			}
	// 		} else if len(userscore_read) < len(playerscore.Levels) {
	// 			//前面已有的更新
	// 			for j := range userscore_read {
	// 				userscore.Id = userscore_read[j].Id
	// 				userscore.Uid = int64(playerscore.Uid)
	// 				userscore.Level = playerscore.Levels[j].GetStageId()
	// 				userscore.Score = playerscore.Levels[j].GetStageScore()
	// 				userscore.Startnum = playerscore.Levels[j].GetStageLevel()
	// 				userscore.Time = int64(playerscore.Levels[j].GetTimestamp())

	// 				if num, err := o.Update(&userscore); err == nil {

	// 					beego.Debug("更新成功！！", num)

	// 				} else {

	// 					beego.Error("更新失败！！！")

	// 				}
	// 			}
	// 			//后面的插入
	// 			for k := len(userscore_read); k < len(playerscore.Levels); k++ {
	// 				userscore.Uid = int64(playerscore.Uid)
	// 				userscore.Level = playerscore.Levels[k].GetStageId()
	// 				userscore.Score = playerscore.Levels[k].GetStageScore()
	// 				userscore.Startnum = playerscore.Levels[k].GetStageLevel()
	// 				userscore.Time = int64(playerscore.Levels[k].GetTimestamp())
	// 				id, err := o.Insert(&userscore)
	// 				if err == nil {

	// 					beego.Debug("插入成功！！", id)

	// 				} else {

	// 					beego.Error("插入失败！！！")

	// 				}
	// 			}
	// 		}
	// 	}
	// }

	var StageNtf []*cspb.CSStageNtf
	var stagereport models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&stagereport)
	for i := range stagereport.Levels {
		StageNtf = append(StageNtf, makeStageNtf(stagereport.Levels[i].StageId, stagereport.Levels[i].StageLevel, stagereport.Levels[i].StageScore, stagereport.Levels[i].GetMedal, stagereport.Levels[i].MedalIsAdd))
	}
	beego.Info("StageNtf:", StageNtf)
	res_data := new(cspb.CSStageReportRes)
	*res_data = cspb.CSStageReportRes{
		Ret:      proto.Int32(ret),
		StageNtf: StageNtf,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		StageReportRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kStageReportRes),
		res_pkg_body, res_list)
	return ret

}
