package collection

import (
	"github.com/astaxie/beego"
)
import "labix.org/v2/mgo"
import "labix.org/v2/mgo/bson"

import db_session "tuojie.com/piggo/quickstart.git/db/session"

type UserId struct {
	Id  bson.ObjectId `bson:"_id"`
	Uid int64         `bson:"uid"`
}

func InitUserId() {
	beego.Debug("******InitUserId")
	c := db_session.DB("zoo").C("userid")

	beego.Debug("new userid")
	user_id := new(UserId)
	beego.Debug("Find %v %v %v", user_id, c, bson.M{"uid": bson.M{"$gt": 0}})

	err := c.Find(bson.M{"uid": bson.M{"$gt": 0}}).One(user_id)
	beego.Debug("err:%v", err)
	if err == mgo.ErrNotFound {
		beego.Debug("not found")
		user_id.Id = bson.NewObjectId()
		user_id.Uid = 100000
		err = c.Insert(user_id)
		if err != nil {
			beego.Error("Init insert fail err:%v", err)
			return
		}
		beego.Info("Init first uid:%d", user_id.Uid)
	} else if err != nil {
		beego.Error("Init fail err:%v", err)
	}
	beego.Debug("InitUserId******")
}

func GetUid() int64 {
	beego.Debug("******GetUid")
	c := db_session.DB("zoo").C("userid")
	result := bson.M{}
	_, err := c.Find(bson.M{"uid": bson.M{"$gt": 0}}).Apply(
		mgo.Change{Update: bson.M{"$inc": bson.M{"uid": 1}}, ReturnNew: true},
		result)
	if err == nil {
		uid := result["uid"].(int64)
		return uid
	} else {
		beego.Error("GetUid fail err:%v", err)
		return -1
	}
}
