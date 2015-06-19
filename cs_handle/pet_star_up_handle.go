package cs_handle

import (
	"github.com/astaxie/beego"
)
import cspb "protocol"
import proto "code.google.com/p/goprotobuf/proto"
import db "tuojie.com/piggo/quickstart.git/db/collection"

import resmgr "tuojie.com/piggo/quickstart.git/res_mgr"

func petStarUpHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {

	pet_id := req.GetBody().GetPetStarUpReq().GetPetId()
	pet_db, ret_pet := db.GetPetById(res_list.GetSAccount(), pet_id)
	if ret_pet < 0 {
		beego.Error("get pet info error ret:%d", ret_pet)
		return makePetStarUpResPkg(req, res_list, ret_pet)
	}

	if ret_pet == 1 {
		if pet_id >= int32(len(resmgr.PetData.GetItems())) {
			beego.Error("get pet info error pet_id:%d, ret:%d", pet_id, ret_pet)
			return makePetStarUpResPkg(req, res_list, ret_pet)
		}
		start_star_level := resmgr.PetData.GetItems()[pet_id-1].GetStartStarLevel()
		if start_star_level <= 0 {
			beego.Error("get pet info error pet_id:%d, ret:%d", pet_id, ret_pet)
			return makePetStarUpResPkg(req, res_list, ret_pet)
		}
		pet_db = db.Pet{
			Account:      res_list.GetSAccount(),
			PetId:        pet_id,
			PetLevel:     1,
			PetCurExp:    0,
			PetTotalExp:  0,
			PetStarLevel: start_star_level,
		}
	}

	chip_db, ret_chip := db.GetPetChipByType(res_list.GetSAccount(), pet_id)
	if ret_chip != 0 {
		beego.Error("get chip info error ret:%d", ret_chip)
		return makePetStarUpResPkg(req, res_list, ret_chip)
	}

	if pet_db.PetStarLevel < 0 {
		beego.Error("db pet star level is invalid star_leve:%d", pet_db.PetStarLevel)
		return makePetStarUpResPkg(req, res_list, int32(-1))
	}

	var max_star_level int32 = resmgr.PetData.GetItems()[pet_id-1].GetMaxStarLevel()
	if pet_db.PetStarLevel >= max_star_level {
		beego.Error("db pet star level is max star_level:%d, max_level:%d",
			pet_db.PetStarLevel, max_star_level)
		return makePetStarUpResPkg(req, res_list,
			int32(cspb.ErrorCode_PetStarLevelIsMax))
	}

	need_chip_num := getNeedChipNum(pet_id, int32(pet_db.PetStarLevel))

	if chip_db.ChipNum < need_chip_num || need_chip_num == 0 {
		beego.Error("db chip num(%d) is less need chip num(%d)",
			chip_db.ChipNum, need_chip_num)

		return makePetStarUpResPkg(req, res_list,
			int32(cspb.ErrorCode_ChipNumIsNotEnough))

	}

	db.ChangeChip(res_list.GetSAccount(), chip_db.ChipType, -need_chip_num,
		chip_db.ChipId)

	var chip_list []*cspb.ChipInfo
	chip_list = append(chip_list, makeChip(chip_db.ChipId,
		chip_db.ChipType,
		need_chip_num,
		int32(cspb.ChangeType_Deduct)))

	var pet_list []*cspb.PetInfo

	pet_list = append(pet_list, makePet(pet_db.PetId,
		pet_db.PetLevel,
		pet_db.PetCurExp,
		pet_db.PetTotalExp,
		pet_db.PetStarLevel+int32(1),
		pet_db.Petmedallevel,
		pet_db.PetmedalNum,
		pet_db.DressId,
	))
	db.SetPetInfo(res_list.GetSAccount(),
		pet_db.PetId, pet_db.PetLevel,
		pet_db.PetCurExp, pet_db.PetTotalExp, pet_db.PetStarLevel+int32(1))

	//添加pet_ntf到res_list中
	makeChipNtf(chip_list, res_list)

	//添加pet_ntf到res_list中
	makePetNtf(pet_list, res_list)

	return makePetStarUpResPkg(req, res_list, int32(0))
}

func makePetStarUpResPkg(req *cspb.CSPkg,
	res_list *cspb.CSPkgList, ret int32) int32 {

	//填充petlistres回包
	res_data := new(cspb.CSPetStarUpRes)
	*res_data = cspb.CSPetStarUpRes{
		Ret:   proto.Int32(ret),
		PetId: proto.Int32(req.GetBody().GetPetStarUpReq().GetPetId()),
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		PetStarUpRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kPetStarUpRes),
		res_pkg_body, res_list)
	return ret
}
