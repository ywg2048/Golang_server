package cs_handle

import (
	"fmt"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
	models "tuojie.com/piggo/quickstart.git/models"
)

import cspb "protocol"
import proto "code.google.com/p/goprotobuf/proto"
import "labix.org/v2/mgo/bson"

import db_session "tuojie.com/piggo/quickstart.git/db/session"

func getPetListHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("****getPetListHandle Start******")
	req_data := req.GetBody().GetPetListReq()
	beego.Info(req_data)
	ret := int32(1)
	c := db_session.DB("zoo").C("player")
	var player models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player)
	if req_data.GetData() != int32(-1) {
		for i := range player.Star {
			for j := range req_data.GetPetList() {
				if player.Star[i].StarId == req_data.GetPetList()[j].GetPetId() {
					c.Upsert(bson.M{"uid": int32(res_list.GetUid())},
						bson.M{"star." + fmt.Sprint(i) + ".level": req_data.GetPetList()[j].GetPetLevel(),
							"star." + fmt.Sprint(i) + ".current_exp":  req_data.GetPetList()[j].GetPetCurExp(),
							"star." + fmt.Sprint(i) + ".dress":        req_data.GetPetList()[j].GetDressId(),
							"star." + fmt.Sprint(i) + ".fighting":     req_data.GetPetList()[j].GetFighting(),
							"star." + fmt.Sprint(i) + ".satisfaction": req_data.GetPetList()[j].GetSatisfaction(),
							"star." + fmt.Sprint(i) + ".fight_exp":    req_data.GetPetList()[j].GetFightExp()})
				}
			}
		}
	}

	var player_return models.Player
	c.Find(bson.M{"uid": int32(res_list.GetUid())}).One(&player_return)
	var PetList []*cspb.PetInfo
	for i := range player_return.Star {
		PetList = append(PetList, makePet(player_return.Star[i].StarId, player_return.Star[i].Level, player_return.Star[i].Currentexp, player_return.Star[i].Satisfaction, player_return.Star[i].FightExp, player_return.Star[i].Fighting, player_return.Star[i].Dress))
	}
	//填充petlistres回包
	res_data := new(cspb.CSPetListRes)
	*res_data = cspb.CSPetListRes{
		Ret:     proto.Int32(ret),
		PetList: PetList,
	}
	beego.Info(res_data)
	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		PetListRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kPetListRes),
		res_pkg_body, res_list)
	return ret
}
