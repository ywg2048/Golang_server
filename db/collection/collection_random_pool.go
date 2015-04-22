package collection

import "labix.org/v2/mgo"
import "labix.org/v2/mgo/bson"
import log "code.google.com/p/log4go"
import db_session "tuojie.com/piggo/quickstart.git/db/session"

type RandomPool struct {
	Account    string `bson:"account"`
	Index      int32  `bson:"index"`
	PoolType   int32  `bson:"pool_type"`
	GoodsId    int32  `bson:"goods_id"`
	UpdateTime int64  `bson:"update_time"`
	HasRandom  int32  `bson:"has_random"`
}

func GetRandomPool(account string, pool_type int32) ([]RandomPool, int32) {
	log.Debug("account:%s, pool_type:%d", account, pool_type)
	c := db_session.DB("zoo").C("randompool")
	var random_pool_list []RandomPool
	err := c.Find(bson.M{"account": account, "pool_type": pool_type}).
		All(&random_pool_list)
	if err == mgo.ErrNotFound {
		log.Error("load random_pool_list no found player. account:%s, err:%v", account, err)
		return random_pool_list, 1
	} else if err != nil {
		log.Error("load random_pool_list fail err:%v", err)
		return random_pool_list, -1
	}

	log.Debug("random_pool_list:%v", random_pool_list)
	return random_pool_list, 0
}

func SetRandomPool(account string, index int32,
	pool_type int32, goods_id int32,
	update_time int64, has_random int32) int32 {

	log.Debug("account:%s, index:%d, pool_type:%d, goods_id:%d, update_time:%d, has_random:%d",
		account, index, pool_type, goods_id, update_time, has_random)

	c := db_session.DB("zoo").C("randompool")
	_, err := c.Upsert(bson.M{"account": account, "index": index, "pool_type": pool_type},

		bson.M{"$set": bson.M{
			"goods_id":    goods_id,
			"update_time": update_time,
			"has_random":  has_random}})
	if err != nil {
		log.Error("SetRandomPool fail err:%v", err)
		return -1
	}

	log.Debug("SetRandomPool succsess")
	return 0
}
