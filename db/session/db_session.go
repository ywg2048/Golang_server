package db_session

import (
	"github.com/astaxie/beego"
)

//dbinterface

import "sync"
import "labix.org/v2/mgo"
import _ "labix.org/v2/mgo/bson"

var mongodbSession *mgo.Session
var dbInitLock sync.Mutex

func DB(data_base_name string) *mgo.Database {
	if mongodbSession == nil {
		beego.Error("mongodbSession not init")
	}
	return mongodbSession.DB(data_base_name)
}

func Init(ip string) error {
	beego.Info("******Init")

	if mongodbSession != nil {
		beego.Error("dbSession has Init")
		return nil
	}

	dbInitLock.Lock()
	defer dbInitLock.Unlock()

	if mongodbSession != nil {
		beego.Error("dbSession has Init")
		return nil
	}

	beego.Info("dbSession connect %s", ip)
	session, err := mgo.Dial(ip)
	if err != nil {
		beego.Error("connect db error %v", err)
		return err
	}
	session.SetMode(mgo.Monotonic, true)
	mongodbSession = session

	beego.Info("dbSession succ******")
	return nil
}

func Finish() {
	beego.Info("dbSession close")
	mongodbSession.Close()
}

func dbSession() *mgo.Session {
	return mongodbSession
}
