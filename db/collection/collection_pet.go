package collection

import "labix.org/v2/mgo"
import "labix.org/v2/mgo/bson"
import log "code.google.com/p/log4go"
import db_session "tuojie.com/piggo/quickstart.git/db/session"
import "fmt"

type Pet struct {
	Account      string `bson:"account"`
	PetId        int32  `bson:"pet_id"`
	PetLevel     int32  `bson:"pet_level"`
	PetCurExp    int32  `bson:"pet_cur_exp"`
	PetTotalExp  int32  `bson:"pet_total_exp"`
	PetStarLevel int32  `bson:"pet_star_level"`
}

type Chip struct {
	Account  string `bson:"account"`
	ChipId   int32  `bson:"chip_id"`
	ChipType int32  `bson:"chip_type"`
	ChipNum  int32  `bson:"chip_num"`
}

func GetPetList(account string) ([]Pet, int32) {
	log.Debug("account:%s", account)
	c := db_session.DB("zoo").C("pet")
	var pet_list []Pet
	err := c.Find(bson.M{"account": account}).
		Sort("+pet_id").All(&pet_list)
	if err == mgo.ErrNotFound {
		log.Error("load pet_list no found player. account:%s, err:%v", account, err)
		return pet_list, 1
	} else if err != nil {
		log.Error("load pet_list fail err:%v", err)
		return pet_list, -1
	}

	log.Debug("pet_list_info:%s", fmt.Sprint(pet_list))
	return pet_list, 0
}

func GetPetById(account string, pet_id int32) (Pet, int32) {
	log.Debug("account:%s. pet_id:%d", account, pet_id)
	c := db_session.DB("zoo").C("pet")
	var pet_info Pet
	err := c.Find(bson.M{"account": account, "pet_id": pet_id}).One(&pet_info)
	if err == mgo.ErrNotFound {
		log.Error("load pet_info no found player. account:%s, err:%v", account, err)
		return pet_info, 1
	} else if err != nil {
		log.Error("load pet_info fail err:%v", err)
		return pet_info, -1
	}

	log.Debug("pet_info:%s", fmt.Sprint(pet_info))
	return pet_info, 0
}

func SetPetInfo(account string, pet_id int32,
	pet_level int32, pet_cur_exp int32,
	pet_total_exp int32, pet_star_level int32) int32 {

	log.Debug("account:%s, pet_id:%d, pet_level:%d, pet_cur_exp:%d, pet_total_exp:%d, pet_star_level:%d",
		account, pet_id, pet_level, pet_cur_exp, pet_total_exp, pet_star_level)

	c := db_session.DB("zoo").C("pet")
	_, err := c.Upsert(bson.M{"account": account, "pet_id": pet_id},

		bson.M{"$set": bson.M{
			"pet_level":      pet_level,
			"pet_cur_exp":    pet_cur_exp,
			"pet_total_exp":  pet_total_exp,
			"pet_star_level": pet_star_level}})
	if err != nil {
		log.Error("SetPetInfo fail err:%v", err)
		return -1
	}

	log.Debug("SetPetInfo succsess")
	return 0
}

func GetChipList(account string) ([]Chip, int32) {
	log.Debug("account:%s", account)
	c := db_session.DB("zoo").C("chip")
	var chip_list []Chip
	err := c.Find(bson.M{"account": account}).
		Sort("+chip_type").All(&chip_list)
	if err == mgo.ErrNotFound {
		log.Error("load chip_list no found. account:%s, err:%v", account, err)
		return chip_list, 1
	} else if err != nil {
		log.Error("load chip_list fail err:%v", err)
		return chip_list, -1
	}

	log.Debug("chip_list_info:%s", fmt.Sprint(chip_list))
	return chip_list, 0
}

func GetPetChipByType(account string, chip_type int32) (Chip, int32) {
	log.Debug("account:%s, chip_type:%d", account, chip_type)
	c := db_session.DB("zoo").C("chip")
	var chip_info Chip
	err := c.Find(bson.M{"account": account, "chip_type": chip_type}).One(&chip_info)
	if err == mgo.ErrNotFound {
		log.Error("load chip_info no found. account:%s, err:%v", account, err)
		return chip_info, 1
	} else if err != nil {
		log.Error("load chip_info fail err:%v", err)
		return chip_info, -1
	}

	log.Debug("chip_info:%s", fmt.Sprint(chip_info))
	return chip_info, 0
}

func GetExpChip(account string) ([]Chip, int32) {
	log.Debug("account:%s", account)
	c := db_session.DB("zoo").C("chip")
	var chip_list []Chip
	err := c.Find(bson.M{"account": account, "chip_type": 0}).All(&chip_list)
	if err == mgo.ErrNotFound {
		log.Error("load chip_list no found. account:%s, err:%v", account, err)
		return chip_list, 1
	} else if err != nil {
		log.Error("load chip_list fail err:%v", err)
		return chip_list, -1
	}

	log.Debug("chip_list_info:%s", fmt.Sprint(chip_list))
	return chip_list, 0
}

func ChangeChip(account string, chip_type int32, chip_num int32, chip_id int32) int32 {

	log.Debug("account:%s, chip_type:%d, chip_num:%d, chip_id:%d",
		account, chip_type, chip_num, chip_id)

	c := db_session.DB("zoo").C("chip")

	if chip_num > 0 { //增加碎片数量
		err := c.Update(bson.M{"account": account, "chip_id": chip_id},

			bson.M{"$inc": bson.M{"chip_num": chip_num}})
		if err == mgo.ErrNotFound {
			log.Debug("no found pet chip, chip_type:%d, err:%v", chip_type, err)
			_, err = c.Upsert(bson.M{"account": account, "chip_id": chip_id},
				bson.M{"$set": bson.M{"chip_num": chip_num, "chip_type": chip_type}})
			if err != nil {
				log.Error("add a new pet chip fail err:%v", err)
				return -1
			}
			log.Debug("add a new pet chip, err:%v", err)
		} else if err != nil {
			log.Error("ChangeChip fail err:%v", err)
			return -1
		}
	} else if chip_num < 0 { //减少碎片数量
		err := c.Update(
			bson.M{"account": account,
				"chip_id":  chip_id,
				"chip_num": bson.M{"$gte": -chip_num}},

			bson.M{"$inc": bson.M{"chip_num": chip_num}})
		if err == mgo.ErrNotFound {
			log.Error("no found pet chip, chip_type:%d, err:%v", chip_type, err)
		} else if err != nil {
			log.Error("ChangeChip fail err:%v", err)
			return -1
		}
	}

	log.Debug("ChangeChip succsess")
	return 0
}
