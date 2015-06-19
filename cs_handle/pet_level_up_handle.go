package cs_handle

import (
	"github.com/astaxie/beego"
)
import cspb "protocol"
import proto "code.google.com/p/goprotobuf/proto"
import db "tuojie.com/piggo/quickstart.git/db/collection"

import resmgr "tuojie.com/piggo/quickstart.git/res_mgr"

func petLevelUpHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {

	pet_id := req.GetBody().GetPetLevelUpReq().GetPetId()
	pet_db, ret_pet := db.GetPetById(res_list.GetSAccount(), pet_id)
	if ret_pet < 0 {
		beego.Error("get pet info error ret:%d", ret_pet)
		return makePetLevelUpResPkg(req, res_list,
			int32(cspb.ErrorCode_PetIdIsInvalid))
	}

	if ret_pet == 1 {
		if pet_id >= int32(len(resmgr.PetData.GetItems())) {
			beego.Error("get pet info error pet_id:%d, ret:%d", pet_id, ret_pet)
			return makePetLevelUpResPkg(req, res_list, ret_pet)
		}
		start_star_level := resmgr.PetData.GetItems()[pet_id-1].GetStartStarLevel()
		if start_star_level <= 0 {
			beego.Error("get pet info error pet_id:%d, ret:%d", pet_id, ret_pet)
			return makePetLevelUpResPkg(req, res_list, ret_pet)
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

	if pet_db.PetLevel >= getPetMaxLvByStarLv(pet_id, pet_db.PetStarLevel) {
		beego.Error("account:%s, pet_level:%d is max level",
			res_list.GetSAccount(), pet_db.PetLevel)
		return makePetLevelUpResPkg(req, res_list,
			int32(cspb.ErrorCode_PetLevelIsMax))
	}

	chip_db_list, ret_chip := db.GetExpChip(res_list.GetSAccount())
	if ret_chip != 0 {
		beego.Error("get chip list error ret:%d", ret_chip)
		return makePetLevelUpResPkg(req, res_list,
			int32(cspb.ErrorCode_ChipIdIsInvalid))
	}

	map_chip_db := make(map[int32]db.Chip)
	for _, db_chip := range chip_db_list {
		map_chip_db[db_chip.ChipId] = db_chip
	}
	beego.Debug("map_chip_db:%v", map_chip_db)

	map_chip_req := make(map[int32]int32) //map[chip_id]chip_num
	for _, chip_info := range req.GetBody().GetPetLevelUpReq().GetChipList() {
		map_chip_req[chip_info.GetChipId()] += chip_info.GetChipNum()
	}
	beego.Debug("map_chip_req:%v", map_chip_req)

	var chip_list []*cspb.ChipInfo
	for req_chip_id, req_chip_num := range map_chip_req {
		beego.Debug("req_chip_id:%d, req_chip_num:%d", req_chip_id, req_chip_num)
		chip_db_info, find_chip := map_chip_db[req_chip_id]
		beego.Debug("chip_db_info.chip_id:%d, p_chip_db_info:%p",
			chip_db_info.ChipId, &chip_db_info)
		if find_chip {
			if req_chip_num > chip_db_info.ChipNum {
				beego.Error("chip is not enough, rep_num:%d, has_num:%d",
					req_chip_num, chip_db_info.ChipNum)
				return makePetLevelUpResPkg(req, res_list,
					int32(cspb.ErrorCode_ChipNumIsNotEnough))
			}
			chip_list = append(chip_list, makeChip(chip_db_info.ChipId,
				chip_db_info.ChipType,
				req_chip_num,
				int32(cspb.ChangeType_Deduct)))
			db.ChangeChip(res_list.GetSAccount(), chip_db_info.ChipType, -req_chip_num,
				chip_db_info.ChipId)
		} else {
			beego.Error("no find chip in db chip_id:%d", req_chip_id)
			return makePetLevelUpResPkg(req, res_list, int32(cspb.ErrorCode_ChipIdIsInvalid))
		}
	}

	var pet_list []*cspb.PetInfo
	new_pet := petLevelUp(req, &pet_db)
	//if new_pet != pet_db {
	pet_list = append(pet_list, makePet(new_pet.PetId,
		new_pet.PetLevel,
		new_pet.PetCurExp,
		new_pet.PetTotalExp,
		new_pet.PetStarLevel,
		new_pet.Petmedallevel,
		new_pet.PetmedalNum,
		new_pet.DressId,
	))
	db.SetPetInfo(res_list.GetSAccount(),
		new_pet.PetId, new_pet.PetLevel,
		new_pet.PetCurExp, new_pet.PetTotalExp,
		pet_db.PetStarLevel)
	//}

	//添加pet_ntf到res_list中
	makeChipNtf(chip_list, res_list)

	//添加pet_ntf到res_list中
	makePetNtf(pet_list, res_list)

	return makePetLevelUpResPkg(req, res_list, int32(0))
}

func petLevelUp(req *cspb.CSPkg, pet_db_info *db.Pet) db.Pet {
	beego.Debug("pet level up start , old_pet:%v ", pet_db_info)
	var provide_exp int32 = 0
	for _, chip_info := range req.GetBody().GetPetLevelUpReq().GetChipList() {
		res_chip := resmgr.ChipData.GetItems()[chip_info.GetChipId()-1]
		provide_exp += res_chip.GetProvideExp() * chip_info.GetChipNum()
	}
	beego.Debug("chip provide_exp:%d", provide_exp)
	if provide_exp <= 0 {
		beego.Error("provide_exp:%d", provide_exp)
		return *pet_db_info
	}
	var new_pet_db_info db.Pet = *pet_db_info

	new_pet_db_info.PetTotalExp += provide_exp

	var need_exp int32 = 0
	var temp_exp int32 = 0
	exp_pool := provide_exp + pet_db_info.PetCurExp
	pet_max_level := getPetMaxLvByStarLv(pet_db_info.PetId, pet_db_info.PetStarLevel)
	for exp_pool > 0 {
		if new_pet_db_info.PetLevel >= pet_max_level {
			beego.Debug("is max level pet_id:%d, level:%d, max_level:%d",
				pet_db_info.PetId, pet_db_info.PetLevel, pet_max_level)
			break
		}
		res_pet_level := resmgr.PetlevelData.GetItems()[new_pet_db_info.PetLevel-1]
		need_exp = res_pet_level.GetNextLevelNeedExp()
		temp_exp = exp_pool - need_exp
		if temp_exp >= 0 {
			new_pet_db_info.PetLevel++
			new_pet_db_info.PetCurExp = 0
			new_pet_db_info.PetTotalExp += need_exp
		} else {
			new_pet_db_info.PetCurExp = exp_pool
			new_pet_db_info.PetTotalExp += exp_pool
		}
		exp_pool -= need_exp
	}
	new_pet_db_info.Petcard[0].CardId = int32(1)
	new_pet_db_info.Petcard[0].CardNum = int32(5)
	beego.Debug("pet level up end , new_pet:%v ", new_pet_db_info)
	return new_pet_db_info
}

func makePetLevelUpResPkg(req *cspb.CSPkg,
	res_list *cspb.CSPkgList, ret int32) int32 {

	//填充petlistres回包
	res_data := new(cspb.CSPetLevelUpRes)
	*res_data = cspb.CSPetLevelUpRes{
		Ret:   proto.Int32(ret),
		PetId: proto.Int32(req.GetBody().GetPetLevelUpReq().GetPetId()),
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		PetLevelUpRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kPetLevelUpRes),
		res_pkg_body, res_list)
	return ret
}
