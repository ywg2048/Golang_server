package collection

// import "labix.org/v2/mgo"
// import "labix.org/v2/mgo/bson"
// import (
// 	"github.com/astaxie/beego"
// 	// 	"github.com/astaxie/beego/orm"
// 	// 	"strconv"
// 	models "tuojie.com/piggo/quickstart.git/models"
// )

// //import log "code.google.com/p/log4go"
// import db_session "tuojie.com/piggo/quickstart.git/db/session"

// // import "fmt"
// // import "time"
// // import cspb "protocol"
// // import "regexp"
// // import "strings"

// func UpdateAchievements(uid int64) {
// 	var player models.Player
// 	c := db_session.DB("zoo").C("player")
// 	c.Find(bson.M{"uid": uid}).One(&player)
// 	beego.Info("---Collection Achievement---", player)

// 	//成就1
// 	c.Upsert(bson.M{"uid": uid},
// 		bson.M{"$set": bson.M{"achievement." + fmt.Sprint(1) + ".achievementid":int32(1),"achievement." + fmt.Sprint(1) + ".process":}})
// }
// func maxlevel(uid int64)int32 {
// 	//查找最大等级数
// 	MaxLevel := int32(0)
// 	c := db_session.DB("zoo").C("player")
// 	var player models.Player
// 	err := c.Find(bson.M{"uid": uid}).One(&player)
// 	if err != nil {
// 		beego.Error("没有这样的玩家！", err)
// 	}
// 	arr := [...]int32{0}
// 	s := arr[0:]

// 	return MaxStage
// }
