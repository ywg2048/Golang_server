package db_session

//dbinterface

import "sync"
import "labix.org/v2/mgo"
import _ "labix.org/v2/mgo/bson"
import log "code.google.com/p/log4go"

var mongodbSession *mgo.Session
var dbInitLock sync.Mutex

func DB(data_base_name string) *mgo.Database {
	if mongodbSession == nil {
		log.Error("mongodbSession not init")
	}
	return mongodbSession.DB(data_base_name)
}

func Init(ip string) error {
	log.Info("******Init")

	if mongodbSession != nil {
		log.Error("dbSession has Init")
		return nil
	}

	dbInitLock.Lock()
	defer dbInitLock.Unlock()

	if mongodbSession != nil {
		log.Error("dbSession has Init")
		return nil
	}

	log.Info("dbSession connect %s", ip)
	session, err := mgo.Dial(ip)
	if err != nil {
		log.Error("connect db error %v", err)
		return err
	}
	session.SetMode(mgo.Monotonic, true)
	mongodbSession = session

	log.Info("dbSession succ******")
	return nil
}

func Finish() {
	log.Info("dbSession close")
	mongodbSession.Close()
}

func dbSession() *mgo.Session {
	return mongodbSession
}
