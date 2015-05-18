package cs_handle

import (
	"github.com/astaxie/beego"
)
import cspb "protocol"

import proto "code.google.com/p/goprotobuf/proto"

func makeAttrInt32Ntf(attr_id int32, value int32,
	change_type int32, res_list *cspb.CSPkgList) {

	attr_ntf := new(cspb.CSAttrNtf)
	*attr_ntf = cspb.CSAttrNtf{
		AttrList: nil,
	}
	attr_ntf.AttrList = addAttrInt32(attr_id, value, change_type)

	if len(attr_ntf.GetAttrList()) <= 0 {
		beego.Error("no attr to add res_list")
		return
	}
	//添加attr_ntf到res_list中
	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		AttrNtf: attr_ntf,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kAttrNtf),
		res_pkg_body, res_list)
}

func makeAttrItemNtf(item_id int32, item_num int32,
	change_type int32, res_list *cspb.CSPkgList) {

	attr_ntf := new(cspb.CSAttrNtf)
	*attr_ntf = cspb.CSAttrNtf{
		AttrList: nil,
	}
	attr_ntf.AttrList = addAttrItemOne(item_id, item_num, change_type)

	if len(attr_ntf.GetAttrList()) <= 0 {
		beego.Error("no attr to add res_list")
		return
	}
	//添加attr_ntf到res_list中
	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		AttrNtf: attr_ntf,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kAttrNtf),
		res_pkg_body, res_list)
}

func makePet(pet_id int32, pet_level int32,
	pet_cur_exp int32, pet_total_exp int32,
	pet_star_level int32) *cspb.PetInfo {

	pet_info := new(cspb.PetInfo)
	*pet_info = cspb.PetInfo{
		PetId:        proto.Int32(pet_id),
		PetLevel:     proto.Int32(pet_level),
		PetCurExp:    proto.Int32(pet_cur_exp),
		PetTotalExp:  proto.Int32(pet_total_exp),
		PetStarLevel: proto.Int32(pet_star_level),
	}

	beego.Debug("pet_info:%v", pet_info)
	return pet_info
}

func makePetNtf(pet_list []*cspb.PetInfo,
	res_list *cspb.CSPkgList) {

	pet_ntf := new(cspb.CSPetNtf)
	*pet_ntf = cspb.CSPetNtf{
		PetList: pet_list,
	}
	//添加pet_ntf到res_list中
	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		PetNtf: pet_ntf,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kPetNtf),
		res_pkg_body, res_list)
}

func makeChip(chip_id int32, chip_type int32,
	chip_num int32, change_type int32) *cspb.ChipInfo {

	chip_info := new(cspb.ChipInfo)
	*chip_info = cspb.ChipInfo{
		ChipId:     proto.Int32(chip_id),
		ChipType:   proto.Int32(chip_type),
		ChipNum:    proto.Int32(chip_num),
		ChangeType: proto.Int32(change_type),
	}

	beego.Debug("chip_info:%v", chip_info)
	return chip_info
}
func makeMessage(message_id int32, message_title string,
	message_content string, message_isActive int32, message_time int64) *cspb.CSMessageNtf {

	message_ntf := new(cspb.CSMessageNtf)
	*message_ntf = cspb.CSMessageNtf{
		Id:       proto.Int32(message_id),
		Title:    proto.String(message_title),
		Content:  proto.String(message_content),
		IsActive: proto.Int32(message_isActive),
		Time:     proto.Int64(message_time),
	}

	beego.Debug("message_ntf:%v", message_ntf)
	return message_ntf
}
func makeApplylist(playerid int32, isaccept int32, isrefuse int32, applytime int64, accepttime int64) *cspb.CSApplyListNtf {
	apply_ntf := new(cspb.CSApplyListNtf)
	*apply_ntf = cspb.CSApplyListNtf{
		Playerid:   proto.Int32(playerid),
		IsAccept:   proto.Int32(isaccept),
		IsRefuse:   proto.Int32(isrefuse),
		Applytime:  proto.Int64(applytime),
		Accepttime: proto.Int64(accepttime),
	}
	return apply_ntf
}
func makeChipNtf(chip_list []*cspb.ChipInfo,
	res_list *cspb.CSPkgList) {

	chip_ntf := new(cspb.CSChipNtf)
	*chip_ntf = cspb.CSChipNtf{
		ChipList: chip_list,
	}
	//添加chip_ntf到res_list中
	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		ChipNtf: chip_ntf,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kChipNtf),
		res_pkg_body, res_list)
}

func makePoolNextFreeNtf(hour int32, today int32,
	res_list *cspb.CSPkgList) {

	ntf := new(cspb.CSPoolNextFreeNtf)
	*ntf = cspb.CSPoolNextFreeNtf{
		Hour:  proto.Int32(hour),
		Today: proto.Int32(today),
	}
	//添加chip_ntf到res_list中
	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		PoolNextFreeNtf: ntf,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kPoolNextFreeNtf),
		res_pkg_body, res_list)
}
