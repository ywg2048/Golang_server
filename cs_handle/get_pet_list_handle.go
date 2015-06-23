package cs_handle

import (
	"github.com/astaxie/beego"
)
import cspb "protocol"
import proto "code.google.com/p/goprotobuf/proto"
import pet_db "tuojie.com/piggo/quickstart.git/db/collection"

func getPetListHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {

	pet_db_list, ret := pet_db.GetPetList(res_list.GetSAccount())
	if ret == 1 {
		return makePetListResPkg(req, res_list, ret)
	} else if ret == 0 {
		var pet_list []*cspb.PetInfo

		var cardNtf []*cspb.CSPetCardNtf
		for _, db_info := range pet_db_list {
			for i := range db_info.Petcard {
				cardNtf = append(cardNtf, makecardNtf(db_info.Petcard[i].CardId, db_info.Petcard[i].CardNum))
			}
			pet_list = append(pet_list, makePet(db_info.PetId,
				db_info.PetLevel,
				db_info.PetCurExp,
				db_info.PetTotalExp,
				db_info.PetStarLevel,
				db_info.Petmedallevel,
				db_info.PetmedalNum,
				db_info.DressId,
				cardNtf,
			))
		}
		//添加pet_ntf到res_list中
		makePetNtf(pet_list, res_list)
	} else {
		beego.Error("get pet list fail ret:%d", ret)
	}

	return makePetListResPkg(req, res_list, ret)
}

func makePetListResPkg(req *cspb.CSPkg,
	res_list *cspb.CSPkgList, ret int32) int32 {

	//填充petlistres回包
	res_data := new(cspb.CSPetListRes)
	*res_data = cspb.CSPetListRes{
		Ret: proto.Int32(ret),
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		PetListRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kPetListRes),
		res_pkg_body, res_list)
	return ret
}
