package collection

import "labix.org/v2/mgo"
import "labix.org/v2/mgo/bson"
import log "code.google.com/p/log4go"
import db_session "tuojie.com/piggo/quickstart.git/db/session"

type Mail struct {
	MailId           bson.ObjectId `bson:"_id"`
	Uid              int64         `bson:"uid"`
	ExpirTime        int64         `bson:"expir_time"`
	MailType         int32         `bson:"mail_type"`
	Title            string        `bson:"title"`
	Contents         string        `bson:"contents"`
	AccessoryType    int32         `bson:"accessory_type"`
	AccessorySubType int32         `bson:"accessory_sub_type"`
	AccessoryNum     int32         `bson:"accessory_num"`
}

func AddMail(mail Mail) int32 {
	log.Debug("mail_info:%v", mail)
	c := db_session.DB("zoo").C("mail")

	mail.MailId = bson.NewObjectId()
	err := c.Insert(&mail)
	if err != nil {
		log.Error("insert mail fail err:%v", err)
		return -1
	}
	return 0
}

func GetMailAll(uid int64) ([]Mail, int32) {
	log.Debug("uid:%d", uid)
	c := db_session.DB("zoo").C("mail")
	var mail_list []Mail
	err := c.Find(bson.M{"uid": uid}).All(&mail_list)
	if err == mgo.ErrNotFound {
		log.Debug("mail is not found err:%v, uid:%d",
			err, uid)
		return mail_list, 1
	} else if err != nil {
		log.Error("query mail fail err:%v, uid:%d",
			err, uid)
		return mail_list, -1
	}
	return mail_list, 0
}

func RemoveMail(mail_id string) int32 {
	log.Debug("mail_id:%s", mail_id)
	c := db_session.DB("zoo").C("mail")
	err := c.Remove(bson.M{"_id": bson.ObjectIdHex(mail_id)})
	if err == mgo.ErrNotFound {
		log.Debug("mail is not found err:%v, mail_id:%s",
			err, mail_id)
		return 1
	} else if err != nil {
		log.Error("remove mail fail err:%v, mail_id:%s", err, mail_id)
		return -1
	}
	return 0
}
