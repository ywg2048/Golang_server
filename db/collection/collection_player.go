package collection

import "labix.org/v2/mgo"
import "labix.org/v2/mgo/bson"
import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	models "tuojie.com/piggo/quickstart.git/models"
)

//import log "code.google.com/p/log4go"
import db_session "tuojie.com/piggo/quickstart.git/db/session"
import "fmt"
import "time"
import cspb "protocol"
import "regexp"
import "strings"

import proto "code.google.com/p/goprotobuf/proto"

func init() {
	orm.RegisterDataBase("default", "mysql", "root:@/Monsters?charset=utf8")
}
func makeRank(playuid int32, playname string,
	star string, level int32, medalnum int32) *cspb.CSRankNtf {

	rank_ntf := new(cspb.CSRankNtf)
	*rank_ntf = cspb.CSRankNtf{
		Playuid:  proto.Int32(playuid),
		Playname: proto.String(playname),
		Star:     proto.String(star),
		Level:    proto.Int32(level),
		Medalnum: proto.Int32(medalnum),
	}

	beego.Debug("rank_ntf:%v", rank_ntf)
	return rank_ntf
}
func LoadPlayer(clientAccount string, serverAccount string, uid int64) (int32, models.Player) {
	beego.Info("-----------Testing mysql---------")
	var ranking []models.Ranking
	var cond *orm.Condition
	cond = orm.NewCondition()

	cond = cond.And("Id__gte", 1)

	var qs orm.QuerySeter
	var res_rank []*cspb.CSRankNtf
	qs = orm.NewOrm().QueryTable("ranking").SetCond(cond).OrderBy("-Medal")
	cnt, err := qs.All(&ranking)
	if err != nil {
		beego.Debug("查询数据库失败")
	}
	beego.Debug(ranking, cnt, err)
	for i := range ranking {
		res_rank = append(res_rank, makeRank(ranking[i].Uid, ranking[i].Name, ranking[i].Star, ranking[i].Level, ranking[i].Medal))
	}
	beego.Info("******RES-----messages", res_rank)
	beego.Info("-----------Testing mysql---------")

	beego.Debug("******LoadPlayer:%v", clientAccount)

	var player models.Player
	ret := int32(0)

	beego.Debug("Before client is %s, serverAccount is %s, uid is %s", clientAccount, serverAccount, uid)
	//测试关数据
	//res := SetLevelValuesTest(clientAccount, 1, 2, 60000)
	//log.Debug("*****^^^^^^^^res is %v", res)

	if len(serverAccount) != 24 {
		re := regexp.MustCompile("\"\\w+\"")
		serverAccount = re.FindString(serverAccount)
		serverAccount = strings.Replace(serverAccount, "\"", "", -1)
	}

	beego.Debug("After client is %s, serverAccount is %s, uid is %s", clientAccount, serverAccount, uid)

	if len(serverAccount) > 0 && bson.IsObjectIdHex(serverAccount) {
		player.Saccount = bson.ObjectIdHex(serverAccount)
	}

	//player.Uid = uid

	if len(clientAccount) <= 0 {
		beego.Error("ErrorCode_SysError: lenth(%d) of account is invalid", int(len(clientAccount)))
		ret = int32(cspb.ErrorCode_SysError)
		player.Uid = 0
	} else {
		c := db_session.DB("zoo").C("player")
		beego.Debug("*****clientAccount is %v", clientAccount)
		err := c.Find(bson.M{"c_account": clientAccount}).One(&player)

		beego.Debug("Find account is %v and player is %v, ", err, player)

		switch err {
		case mgo.ErrNotFound:
			beego.Debug("c_account mgo.ErrNotFound...")

			if uid < 0 {
				beego.Error("get uid fail uid:%d", uid)
				ret = int32(cspb.ErrorCode_PlayerInsertFail)
				player.Uid = int64(0)
			} else {

				// var wonderfulFriends WonderfulFriends
				// wonderfulFriends.RegistTime = time.Now().Unix()
				player.Caccount = clientAccount
				//player.WonderfulFriends.RegistTime = time.Now().Unix()
				player.RegistTime = time.Now().Unix()
				player.Saccount = bson.NewObjectId()
				player.Uid = GetUid()

				beego.Info("Write player:%v", player)

				err = c.Insert(&player)

				// 如果插入失败，回退Saccount, Uid
				if err != nil {
					beego.Error("创建新用户失败:%v", err)
					ret = int32(cspb.ErrorCode_PlayerInsertFail)
					player.Uid = int64(0)
				}
			}
		case nil:
			if player.Uid != uid {
				// 在设备上就有ID存在，以设备ID为准
				beego.Error("在设备上就有ID存在，以设备ID为准.服务器上查的MAYAID is %s, 设备存储 is %d ", player.Uid, uid)
				if uid != 0 {
					beego.Error("把服务器存储的id%d更新用户的设备ID%d", player.Uid, uid)
					player.Uid = uid
				}

				err = c.Update(bson.M{"_id": player.Saccount}, player)

				if err != nil {
					beego.Error("更新用户失败:%v", err)
				}
			} else {
				beego.Debug("获取用户%v ", player)
			}
		default:
			// 因为不是nill，就是出错了
			// 而且错误不是ErrNotFound
			beego.Error("ErrorCode_SysError in LoadingPlayer error is: %s", err)
			ret = int32(cspb.ErrorCode_SysError)
		}
	}

	beego.Debug("player_info:%s", fmt.Sprint(player))
	return ret, player
}

func ChangeFieldValue(account string, field_name string, add_value int32) int32 {
	beego.Debug("account:%s, field_name:%s, add_value:%d",
		account, field_name, add_value)
	if len(account) <= 0 {
		beego.Error("lenth(%d) of account is invalid", int(len(account)))
		return -1
	}

	c := db_session.DB("zoo").C("player")

	beego.Debug("Before account is %s", account)

	if len(account) != 24 {
		re := regexp.MustCompile("\"\\w+\"")
		account = re.FindString(account)
		account = strings.Replace(account, "\"", "", -1)
	}

	beego.Debug("Before account is %s", account)

	err := c.Update(bson.M{
		"_id":      bson.ObjectIdHex(account),
		field_name: bson.M{"$gt": -add_value}},

		bson.M{"$inc": bson.M{field_name: add_value}})
	if err == mgo.ErrNotFound {
		if add_value < 0 {
			beego.Error("no found this attr(%s) , attrvalue's count is 0", field_name)
			return -1
		}
		err = c.Update(bson.M{"_id": account},
			bson.M{field_name: add_value})
		beego.Debug("addElementValue new addFieldValue:%s", account)
		return 1
	} else if err != nil {
		beego.Error("addElementValue fail err:%v", err)
		return -1
	}
	return 0
}

func SetLastSignInTime(account string, last_signin_time int64,
	free_signin_oper_time int64) int32 {
	if len(account) <= 0 {
		beego.Error("lenth(%d) of account is invalid", int(len(account)))
		return -1
	}

	beego.Debug("account:%s, last_signin_time:%d, free_signin_oper_time:%d",
		account, last_signin_time, free_signin_oper_time)

	c := db_session.DB("zoo").C("player")

	// if len(account) != 24 {
	// 	re := regexp.MustCompile("\"\\w+\"")
	// 	account = re.FindString(account)
	// 	account = strings.Replace(account, "\"", "", -1)
	// }

	// log.Debug("Before account is %s", account)

	_, err := c.Upsert(bson.M{"c_account": account},
		bson.M{"$set": bson.M{"WonderfulFriends.last_signin_time": last_signin_time, "WonderfulFriends.free_signin_oper_time": free_signin_oper_time}})
	if err != nil {
		beego.Error("SetLastSignInTime fail err:%v", err)
		return -1
	}

	return 0
}

func SetFieldValue(account string, field_name string, set_value int32) int32 {
	beego.Debug("account:%s, field_name:%s, set_value:%d",
		account, field_name, set_value)

	if len(account) <= 0 {
		beego.Error("lenth(%d) of account is invalid", int(len(account)))
		return -1
	}

	c := db_session.DB("zoo").C("player")

	if len(account) != 24 {
		re := regexp.MustCompile("\"\\w+\"")
		account = re.FindString(account)
		account = strings.Replace(account, "\"", "", -1)
	}

	_, err := c.Upsert(bson.M{"_id": bson.ObjectIdHex(account)},
		bson.M{"$set": bson.M{field_name: set_value}})
	if err != nil {
		beego.Error("setElementValue fail err:%v", err)
		return -1
	}
	return 0
}

func SetFielsValues(account string, field_names []string, set_values []int32) int32 {
	beego.Debug("account:%s, field_names:%v, set_values:%v",
		account, field_names, set_values)

	if len(account) <= 0 {
		beego.Error("lenth(%d) of account is invalid", int(len(account)))
		return -1
	}

	map_value := bson.M{}
	for i := 0; i < len(field_names) && i < len(set_values); i++ {
		map_value[field_names[i]] = set_values[i]
	}
	c := db_session.DB("zoo").C("player")

	if len(account) != 24 {
		re := regexp.MustCompile("\"\\w+\"")
		account = re.FindString(account)
		account = strings.Replace(account, "\"", "", -1)
	}

	_, err := c.Upsert(bson.M{"_id": bson.ObjectIdHex(account)},
		bson.M{"$set": map_value})
	if err != nil {
		beego.Error("setElementValue fail err:%v", err)
		return -1
	}
	return 0
}
func SetLevelValues(account string, levelid int32, stage_level int32, stage_score int32) int32 {

	beego.Debug("********account is %v levelid is %v stage_level is %v stage_score is %v", account, levelid, stage_level, stage_score)

	c := db_session.DB("zoo").C("player")

	//var stageScore int32
	var result models.Player
	errs := c.Find(bson.M{"c_account": account}).One(&result)
	if errs != nil {
		beego.Error("Get StageReportDB fail err:%v", errs)
		return -1
	}
	beego.Debug("********result is %v", result)
	//stageScore = result.Levels.level.Score
	//log.Debug("DB StageScore is %v", stageScore)
	beego.Debug("StageScore is %v", stage_score)

	_, err := c.Upsert(bson.M{"c_account": account},
		bson.M{"$set": bson.M{"Levels.level'" + fmt.Sprintf("%d", levelid) + "'.Levelid": levelid, "Levels.level'" + fmt.Sprintf("%d", levelid) + "'.StageLevel": stage_level, "Levels.level'" + fmt.Sprintf("%d", levelid) + "'.Score": stage_score}})
	if err != nil {
		beego.Error("StageReportHandle fail err:%v", err)
		return -1
	}
	return 0
}
