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

type Player struct {
	Saccount           bson.ObjectId        `bson:"_id"`
	Caccount           string               `bson:"c_account"`
	Uid                int64                `bson:"uid"`
	WonderfulFriends   WonderfulFriendsData `bson:"WonderfulFriends"`
	RechargeFlow       RechargeFlowData     `bson:"RechargeFlow"`
	RegistTime         int64                `bson:"regist_time"`
	LastSignInTime     int64                `bson:"last_signin_time"`
	FreeSignInOperTime int64                `bson:"free_signin_oper_time"`
	Levels             []*cspb.CSStageNtf   `bson:"Levels"`
	Money              *cspb.CSMoneyReq     `bson:"Money"`
}
type WonderfulFriendsData struct {
	LastSignInTime     int64 `bson:"last_signin_time"`
	FreeSignInOperTime int64 `bson:"free_signin_oper_time"`
}
type Messagecenter struct {
	Id       int32
	Title    string `orm:"size(100)"`
	Content  string `orm:"size(256)"`
	Title2   string `orm:"size(256)"`
	Content2 string `orm:"size(256)"`
	IsActive int32
	Time     int64
}

// func LoadPlayer(account string) (Player, int32) {
// 	log.Debug("account:%s", account)
// 	var player Player
// 	if len(account) <= 0 {
// 		log.Error("lenth(%d) of account is invalid", int(len(account)))
// 		return player, -1
// 	}
// 	c := db_session.DB("zoo").C("player")

// 	err := c.Find(bson.M{"_id": bson.ObjectIdHex(account)}).One(&player)
// 	if err == mgo.ErrNotFound {
// 		log.Error("load player no found err:%v", err)
// 		return player, 1
// 	} else if err != nil {
// 		log.Error("load player fail err:%v", err)
// 		return player, -1
// 	}

// 	log.Debug("player_info:%s", fmt.Sprint(player))
// 	return player, 0
// }
func init() {
	orm.RegisterDataBase("default", "mysql", "root:@/Monsters?charset=utf8")
}
func makeMessage(message_id int32, message_title string,
	message_content string, message_isActive int32, message_time int64) *cspb.CSMessageNtf {

	message_ntf := new(cspb.CSMessageNtf)
	*message_ntf = cspb.CSMessageNtf{
		Id:      proto.Int32(message_id),
		Title:   proto.String(message_title),
		Content: proto.String(message_content),

		IsActive: proto.Int32(message_isActive),
		Time:     proto.Int64(message_time),
	}

	beego.Debug("message_ntf:%v", message_ntf)
	return message_ntf
}
func LoadPlayer(clientAccount string, serverAccount string, uid int64) (int32, Player) {
	beego.Info("-----------Testing mysql---------")
	var message []models.Messagecenter
	var cond *orm.Condition
	cond = orm.NewCondition()

	cond = cond.And("Id__gte", 1)
	cond = cond.And("IsActive__contains", 1)
	var qs orm.QuerySeter
	qs = orm.NewOrm().QueryTable("messagecenter").Limit(20).SetCond(cond)
	cnt, err := qs.All(&message)

	beego.Debug(message, cnt, err)

	var res_messages []*cspb.CSMessageNtf

	for i := range message {

		res_messages = append(res_messages, makeMessage(message[i].Id, message[i].Title, message[i].Content, message[i].IsActive, message[i].Time))

	}

	beego.Info("******RES-----messages", res_messages)
	beego.Info("-----------Testing mysql---------")

	beego.Debug("******LoadPlayer:%v", clientAccount)

	var player Player
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
	var result Player
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

func SetLevelValuesTest(account string, levelid int32, stage_level int32, stage_score int32) int32 {

	beego.Debug("********account is %v levelid is %v stage_level is %v stage_score is %v", account, levelid, stage_level, stage_score)

	c := db_session.DB("zoo").C("player")

	//var stageScore int32
	var result Player
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
