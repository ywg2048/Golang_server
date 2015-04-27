package cs_handle

import (
	"github.com/astaxie/beego"
)
import cspb "protocol"
import proto "code.google.com/p/goprotobuf/proto"
import db "tuojie.com/piggo/quickstart.git/db/collection"

import resmgr "tuojie.com/piggo/quickstart.git/res_mgr"

func petReportHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {

	pet_id := req.GetBody().GetReportPetReq().GetPetId()
	pet_db, ret_pet := db.GetPetById(res_list.GetSAccount(), pet_id)
	if ret_pet < 0 {
		beego.Error("get pet info error ret:%d", ret_pet)
		return makePetReportResPkg(req, res_list,
			int32(cspb.ErrorCode_PetIdIsInvalid))
	}

	if ret_pet == 0 {
		var pet_list []*cspb.PetInfo
		pet_list = append(pet_list, makePet(pet_db.PetId,
			pet_db.PetLevel,
			pet_db.PetCurExp,
			pet_db.PetTotalExp,
			pet_db.PetStarLevel))

		//添加pet_ntf到res_list中
		makePetNtf(pet_list, res_list)

		beego.Error(" pet is exist pet_id:%d, ret:%d", pet_id, ret_pet)
		return makePetReportResPkg(req, res_list,
			int32(cspb.ErrorCode_PetIdIsExist))
	}

	start_pet_star_level := resmgr.PetData.GetItems()[pet_id-1].GetStartStarLevel()
	db.SetPetInfo(res_list.GetSAccount(),
		pet_id, int32(1),
		int32(0), int32(0), start_pet_star_level)

	var pet_list []*cspb.PetInfo
	pet_list = append(pet_list, makePet(pet_id,
		int32(1),
		int32(0),
		int32(0),
		start_pet_star_level))

	//添加pet_ntf到res_list中
	makePetNtf(pet_list, res_list)

	return makePetReportResPkg(req, res_list, int32(0))
}

func makePetReportResPkg(req *cspb.CSPkg,
	res_list *cspb.CSPkgList, ret int32) int32 {

	//填充petreportres回包
	res_data := new(cspb.CSReportPetRes)
	*res_data = cspb.CSReportPetRes{
		Ret: proto.Int32(ret),
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		ReportPetRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kReportPetRes),
		res_pkg_body, res_list)
	return ret
}
