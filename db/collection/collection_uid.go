package collection

import "labix.org/v2/mgo"
import "labix.org/v2/mgo/bson"
import log "code.google.com/p/log4go"
import db_session "tuojie.com/piggo/quickstart.git/db/session"

type UserId struct {
	Id  bson.ObjectId `bson:"_id"`
	Uid int64         `bson:"uid"`
}

func InitUserId() {
	log.Debug("******InitUserId")
	c := db_session.DB("zoo").C("userid")

	log.Debug("new userid")
	user_id := new(UserId)
	log.Debug("Find %v %v %v", user_id, c, bson.M{"uid": bson.M{"$gt": 0}})

	err := c.Find(bson.M{"uid": bson.M{"$gt": 0}}).One(user_id)
	log.Debug("err:%v", err)
	if err == mgo.ErrNotFound {
		log.Debug("not found")
		user_id.Id = bson.NewObjectId()
		user_id.Uid = 100000
		err = c.Insert(user_id)
		if err != nil {
			log.Error("Init insert fail err:%v", err)
			return
		}
		log.Info("Init first uid:%d", user_id.Uid)
	} else if err != nil {
		log.Error("Init fail err:%v", err)
	}
	log.Debug("InitUserId******")
}

func GetUid() int64 {
	log.Debug("******GetUid")
	c := db_session.DB("zoo").C("userid")
	result := bson.M{}
	_, err := c.Find(bson.M{"uid": bson.M{"$gt": 0}}).Apply(
		mgo.Change{Update: bson.M{"$inc": bson.M{"uid": 1}}, ReturnNew: true},
		result)
	if err == nil {
		uid := result["uid"].(int64)
		return uid
	} else {
		log.Error("GetUid fail err:%v", err)
		return -1
	}
}
