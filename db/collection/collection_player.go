package collection

import "labix.org/v2/mgo"
import "labix.org/v2/mgo/bson"
import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	models "tuojie.com/piggo/quickstart.git/models"
)

//import log "code.google.com/p/log4go"
import db_session "tuojie.com/piggo/quickstart.git/db/session"
import "fmt"
import "time"
import cspb "protocol"
import "regexp"
import "strings"

// import proto "code.google.com/p/goprotobuf/proto"

// import resmgr "tuojie.com/piggo/quickstart.git/res_mgr"

func init() {
	orm.RegisterDataBase("default", "mysql", "root:@/Monsters?charset=utf8")
}

func LoadPlayer(clientAccount string, serverAccount string, uid int64) (int32, models.Player) {

	beego.Debug("Before client is %s, serverAccount is %s, uid is %s", clientAccount, serverAccount, uid)
	//测试关数据
	//res := SetLevelValuesTest(clientAccount, 1, 2, 60000)
	//log.Debug("*****^^^^^^^^res is %v", res)
	var player models.Player
	ret := int32(0)
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

				player.Caccount = clientAccount
				//player.WonderfulFriends.RegistTime = time.Now().Unix()
				player.RegistTime = time.Now().Unix()
				player.Saccount = bson.NewObjectId()
				player.Uid = GetUid()
				player.Diamond = int32(50)
				player.Flower = int32(30)
				player.Gold = int32(2000)
				player.ExperiencePool = int32(3000)
				player.StarId = int32(8)
				beego.Info("Write player:%v", player)

				err = c.Insert(&player)

				// 如果插入失败，回退Saccount, Uid
				if err != nil {
					beego.Error("创建新用户失败:%v", err)
					ret = int32(cspb.ErrorCode_PlayerInsertFail)
					player.Uid = int64(0)
				} else {
					var player models.Player
					c.Find(bson.M{"uid": int32(GetUid())}).One(&player)
					beego.Info("cards len", len(player.Cards))
					if len(player.Cards) == 0 {
						for i := 0; i <= stringToint(beego.AppConfig.String("card_account")); i++ {
							_, err := c.Upsert(bson.M{"uid": int32(GetUid())},
								bson.M{"$set": bson.M{"cards." + fmt.Sprint(i) + ".card_id": (i + 1), "cards." + fmt.Sprint(i) + ".card_num": int32(0)}})
							if err != nil {
								beego.Error("插入失败")
							} else {
								beego.Info("插入成功")
							}
						}
					}
					if len(player.Star) == 0 {
						_, err := c.Upsert(bson.M{"uid": int32(GetUid())},
							bson.M{"$push": bson.M{"star": bson.M{"starid": int32(8), "starname": "皓皓", "level": int32(1), "current_exp": int32(0), "dress": int32(1),
								"dressname": "初级套装", "fighting": int32(16500), "satisfaction": int32(50), "fight_exp": int32(0), "is_active": int32(1)}}})
						if err != nil {
							beego.Error("小伙伴初始化失败", err)
						}
					}
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
func stringToint(value string) int {
	b, error := strconv.Atoi(value)
	if error != nil {
		fmt.Println("字符串转换成整数失败")
	}
	return b
}
