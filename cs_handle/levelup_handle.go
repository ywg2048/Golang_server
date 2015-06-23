package cs_handle

import (
	// "fmt"
	"github.com/astaxie/beego"
	// "time"
	models "tuojie.com/piggo/quickstart.git/models"
)

import cspb "protocol"

// import proto "code.google.com/p/goprotobuf/proto"

// import db "tuojie.com/piggo/quickstart.git/db/collection"
import "labix.org/v2/mgo/bson"

import db_session "tuojie.com/piggo/quickstart.git/db/session"

func LevelUpHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********LevelUpHandle Start**********")
	req_data := req.GetBody().GetLevelUpPageReq()
	beego.Info(req_data)

	ret := int32(1)
	c := db_session.DB("zoo").C("pet")
	var pet models.Pet
	err := c.Find(bson.M{"account": res_list.GetSAccount()}).One(&pet)
	beego.Debug("*********StageReportHandle result is %v err is %v********", pet, err)
	c.Upsert(bson.M{"account": res_list.GetSAccount(), "pet_id": req_data.GetPetinfo().GetPetId()},
		bson.M{"$set": bson.M{
			"pet_level":       req_data.GetPetinfo().GetPetLevel(),
			"pet_cur_exp":     req_data.GetPetinfo().GetPetCurExp(),
			"pet_total_exp":   req_data.GetPetinfo().GetPetTotalExp(),
			"pet_star_level":  req_data.GetPetinfo().GetPetStarLevel(),
			"pet_medal_level": req_data.GetPetinfo().GetPetMedalLevl(),
			"pet_medal_num":   req_data.GetPetinfo().GetPetMedalNum(),
			"dress_id":        req_data.GetPetinfo().GetDressId()}})
	for i := range req_data.GetPetinfo().GetPetCardNtf() {
		c.Upsert(bson.M{"account": res_list.GetSAccount(), "pet_id": req_data.GetPetinfo().GetPetId()},
			bson.M{"$set": bson.M{
				"pet_card.cardid":  req_data.GetPetinfo().GetPetCardNtf()[i].GetCardId(),
				"pet_card.cardnum": req_data.GetPetinfo().GetPetCardNtf()[i].GetCarNum(),
			}})
	}
	var petlist []models.Pet
	_err := c.Find(bson.M{"account": res_list.GetSAccount()}).All(&petlist)
	beego.Info(_err)
	var pet_list []*cspb.PetInfo

	var cardNtf []*cspb.CSPetCardNtf
	for j := range petlist {
		for i := range petlist[j].Petcard {
			cardNtf = append(cardNtf, makecardNtf(petlist[j].Petcard[i].CardId, petlist[j].Petcard[i].CardNum))
		}
		pet_list = append(pet_list, makePet(petlist[j].PetId,
			petlist[j].PetLevel,
			petlist[j].PetCurExp,
			petlist[j].PetTotalExp,
			petlist[j].PetStarLevel,
			petlist[j].Petmedallevel,
			petlist[j].PetmedalNum,
			petlist[j].DressId,
			cardNtf,
		))
	}
	res_data := new(cspb.CSLevelUpPageRes)
	*res_data = cspb.CSLevelUpPageRes{
		Petinfo: pet_list,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		LevelUpPageRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kLevelUpPageRes),
		res_pkg_body, res_list)
	return ret

}
