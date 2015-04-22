package cs_handle

import cspb "protocol"
import proto "code.google.com/p/goprotobuf/proto"
import log "code.google.com/p/log4go"
import db "tuojie.com/piggo/quickstart.git/db/collection"
import "resource"
import resmgr "tuojie.com/piggo/quickstart.git/res_mgr"
import rand "github.com/tuojie/utility"
import "strconv"
import "strings"

func makeCSPkgList(
	cmd int32,
	pkg_body *cspb.CSBody,
	pkg_list *cspb.CSPkgList) *cspb.CSPkgList {

	res_pkg_head := new(cspb.CSHead)
	*res_pkg_head = cspb.CSHead{
		Cmd: proto.Int32(cmd),
	}

	pkg := new(cspb.CSPkg)
	*pkg = cspb.CSPkg{
		Head: res_pkg_head,
		Body: pkg_body,
	}
	//append pkg 到的pkg_list中
	pkg_list.Pkgs = append(pkg_list.Pkgs, pkg)
	log.Debug("resinfo:%d->%s\n %s",
		cmd,
		cspb.Command_name[cmd],
		pkg.String())
	return pkg_list
}

func addAttrInt32(attr_id int32, value int32,
	change_type int32) []*cspb.AttrInfo {

	var attr_list []*cspb.AttrInfo
	attr_value := new(cspb.AttrValue)
	switch attr_id {
	case int32(cspb.AttrId_Diamond):
		*attr_value = cspb.AttrValue{
			Diamond: proto.Int32(value),
		}
	case int32(cspb.AttrId_Flower):
		*attr_value = cspb.AttrValue{
			Flower: proto.Int32(value),
		}
	case int32(cspb.AttrId_Gold):
		*attr_value = cspb.AttrValue{
			Gold: proto.Int32(value),
		}
	case int32(cspb.AttrId_CopperKey):
		*attr_value = cspb.AttrValue{
			CopperKey: proto.Int32(value),
		}
	case int32(cspb.AttrId_SilverKey):
		*attr_value = cspb.AttrValue{
			SilverKey: proto.Int32(value),
		}
	case int32(cspb.AttrId_GoldKey):
		*attr_value = cspb.AttrValue{
			GoldKey: proto.Int32(value),
		}

	default:
		log.Error("attr id is invalid attr_id:%d", attr_id)
		return nil
	}

	attr_info := new(cspb.AttrInfo)
	*attr_info = cspb.AttrInfo{
		Id:         proto.Int32(attr_id),
		Value:      attr_value,
		ChangeType: proto.Int32(change_type),
	}

	attr_list = append(attr_list, attr_info)
	log.Debug("attr_list:%v", attr_list)
	return attr_list
}

func addAttrItemOne(item_id int32, item_num int32,
	change_type int32) []*cspb.AttrInfo {

	var attr_list []*cspb.AttrInfo
	attr_value := new(cspb.AttrValue)
	*attr_value = cspb.AttrValue{
		ItemList: nil,
	}

	attr_item := new(cspb.AttrItem)
	*attr_item = cspb.AttrItem{
		ItemId:  proto.Int32(item_id),
		ItemNum: proto.Int32(item_num),
	}
	attr_value.ItemList = append(attr_value.GetItemList(), attr_item)

	attr_info := new(cspb.AttrInfo)
	*attr_info = cspb.AttrInfo{
		Id:         proto.Int32(int32(cspb.AttrId_Item)),
		Value:      attr_value,
		ChangeType: proto.Int32(change_type),
	}

	attr_list = append(attr_list, attr_info)
	log.Debug("attr_list:%v", attr_list)
	return attr_list
}

type SelecteFunc func(int32, int32) bool

func defaultSelecter(random_type int32, goods_id int32) bool {
	return true
}

type randomNumFunc func(int32, *resource.Randomgoods) int32

func randomFromAll(random_type int32, req_version int64,
	selecter SelecteFunc,
	randomNum randomNumFunc) int32 {

	var max_num int32 = 0
	for _, data := range resmgr.RandomgoodsData.GetItems() {
		if !selecter(random_type, data.GetElementId()) {
			continue
		}
		resource_version := IpStringToInt(data.GetVersion())
		if req_version < resource_version {
			continue
		}
		max_num += randomNum(random_type, data)
	}

	var random_num int32 = rand.RandInt32(max_num)
	var cur_num int32
	var goods_id int32 = -1

	for _, data := range resmgr.RandomgoodsData.GetItems() {
		if !selecter(random_type, data.GetElementId()) {
			continue
		}
		resource_version := IpStringToInt(data.GetVersion())
		if req_version < resource_version {
			continue
		}
		cur_num += randomNum(random_type, data)
		if cur_num >= random_num {
			goods_id = data.GetElementId()
			break
		}
	}
	if goods_id == -1 {
		log.Error("random error max_num:%d, random_num:%d, cur_num:%d, req_version:%d",
			max_num, random_num, cur_num, req_version)
	}
	log.Debug("random_type:%d, random_goods_id:%d, req_version:%d",
		random_type, goods_id, req_version)
	return goods_id
}

func processMail(res_list *cspb.CSPkgList) {
	log.Debug("******processMail")

	account := res_list.GetSAccount()
	uid := res_list.GetUid()
	log.Debug("processMail account:%s, uid:%d", account, uid)
	mail_list, ret := db.GetMailAll(uid)
	if ret != 0 || len(mail_list) <= 0 {
		log.Debug("no thing to process ret:%d, len:%d", ret, len(mail_list))
		return
	}
	var pet_list []*cspb.PetInfo
	var chip_list []*cspb.ChipInfo
	for _, mail := range mail_list {

		if mail.AccessoryType == 1 {
			makeAttrInt32Ntf(int32(cspb.AttrId_Diamond),
				mail.AccessoryNum,
				int32(cspb.ChangeType_Add),
				res_list)
		} else if mail.AccessoryType == 2 {
			makeAttrInt32Ntf(int32(cspb.AttrId_Gold),
				mail.AccessoryNum,
				int32(cspb.ChangeType_Add),
				res_list)
		} else if mail.AccessoryType == 3 {
			makeAttrInt32Ntf(int32(cspb.AttrId_Flower),
				mail.AccessoryNum,
				int32(cspb.ChangeType_Add),
				res_list)
		} else if mail.AccessoryType == 4 {
			makeAttrItemNtf(mail.AccessorySubType,
				mail.AccessoryNum,
				int32(cspb.ChangeType_Add),
				res_list)
		} else if mail.AccessoryType == 5 {
			pet_list, chip_list = processPet(res_list.GetSAccount(),
				mail.AccessorySubType, mail.AccessoryNum,
				pet_list, chip_list)
		} else if mail.AccessoryType == 6 {
			pet_list, chip_list = processChip(res_list.GetSAccount(),
				mail.AccessorySubType, mail.AccessoryNum,
				pet_list, chip_list)
		}
		db.RemoveMail(mail.MailId.Hex())
	}

	//发送 chip 或者 pet ntf
	if len(pet_list) > 0 {
		log.Debug("petlist:%v", pet_list)
		for _, pet := range pet_list {
			db.SetPetInfo(res_list.GetSAccount(), pet.GetPetId(),
				1, 0, 0, pet.GetPetStarLevel())
		}
		makePetNtf(pet_list, res_list)
	}

	if len(chip_list) > 0 {
		log.Debug("makeChipList chip_list:%v", chip_list)
		for _, chip := range chip_list {
			change_num := int32(0)
			if chip.GetChangeType() == int32(cspb.ChangeType_Add) {
				change_num = chip.GetChipNum()
			} else if chip.GetChangeType() == int32(cspb.ChangeType_Deduct) {
				change_num = -chip.GetChipNum()
			}
			db.ChangeChip(res_list.GetSAccount(),
				chip.GetChipType(), change_num, chip.GetChipId())
		}
		makeChipNtf(chip_list, res_list)
	}

}

func goodsParseMakeNtf(goods_list []int32, res_list *cspb.CSPkgList) {
	//goods_id 转换到 chip 或者 pet 或者 属性
	var pet_list []*cspb.PetInfo
	var chip_list []*cspb.ChipInfo
	for _, id := range goods_list {
		random_goods := resmgr.RandomgoodsData.GetItems()[id-1]
		if random_goods.GetElementType() == 1 {
			makeAttrInt32Ntf(int32(cspb.AttrId_Diamond),
				random_goods.GetNum(),
				int32(cspb.ChangeType_Add),
				res_list)
		} else if random_goods.GetElementType() == 2 {
			makeAttrInt32Ntf(int32(cspb.AttrId_Gold),
				random_goods.GetNum(),
				int32(cspb.ChangeType_Add),
				res_list)
		} else if random_goods.GetElementType() == 3 {
			makeAttrInt32Ntf(int32(cspb.AttrId_Flower),
				random_goods.GetNum(),
				int32(cspb.ChangeType_Add),
				res_list)
		} else if random_goods.GetElementType() == 4 {
			makeAttrItemNtf(random_goods.GetSubId(),
				random_goods.GetNum(),
				int32(cspb.ChangeType_Add),
				res_list)
		} else if random_goods.GetElementType() == 5 {
			pet_list, chip_list = processPet(res_list.GetSAccount(),
				random_goods.GetSubId(), random_goods.GetNum(),
				pet_list, chip_list)
		} else if random_goods.GetElementType() == 6 {
			pet_list, chip_list = processChip(res_list.GetSAccount(),
				random_goods.GetSubId(), random_goods.GetNum(),
				pet_list, chip_list)
		}

	}
	//发送 chip 或者 pet ntf
	if len(pet_list) > 0 {
		log.Debug("petlist:%v", pet_list)
		for _, pet := range pet_list {
			db.SetPetInfo(res_list.GetSAccount(), pet.GetPetId(),
				1, 0, 0, pet.GetPetStarLevel())
		}
		makePetNtf(pet_list, res_list)
	}

	if len(chip_list) > 0 {
		log.Debug("makeChipList chip_list:%v", chip_list)
		for _, chip := range chip_list {
			change_num := int32(0)
			if chip.GetChangeType() == int32(cspb.ChangeType_Add) {
				change_num = chip.GetChipNum()
			} else if chip.GetChangeType() == int32(cspb.ChangeType_Deduct) {
				change_num = -chip.GetChipNum()
			}
			db.ChangeChip(res_list.GetSAccount(),
				chip.GetChipType(), change_num, chip.GetChipId())
		}
		makeChipNtf(chip_list, res_list)
	}
}

func processPet(
	account string, pet_id int32, num int32,
	pet_list []*cspb.PetInfo,
	chip_list []*cspb.ChipInfo) ([]*cspb.PetInfo, []*cspb.ChipInfo) {

	_, ret := db.GetPetById(account, pet_id)
	if ret < 0 {
		log.Error("get pet db error saccount:%s, pet_id:%d", account, pet_id)
		return pet_list, chip_list
	}

	//自己没有该宠物，直接添加到pet_list
	start_star_level := resmgr.PetData.GetItems()[pet_id-1].GetStartStarLevel()
	if ret == 1 {
		pet_list = append(pet_list, makePet(pet_id, 1, 0, 0, start_star_level))
		if num == 1 {
			//增加的数量只有一个，直接返回啦
			return pet_list, chip_list
		}
	}

	//有该宠物的话 或者 增加的个数大于1个，转换成碎片
	for _, chip_res := range resmgr.ChipData.GetItems() {
		//这里 碎片的类型就是宠物的id
		if chip_res.GetChipType() == pet_id {
			add_chip_num := num * getNeedChipNum(pet_id, start_star_level)
			chip_list = append(chip_list, makeChip(chip_res.GetChipId(),
				chip_res.GetChipType(), add_chip_num,
				int32(cspb.ChangeType_Add)))
			log.Debug("find pet to chip petid:%d", pet_id)
			return pet_list, chip_list
		}
	}
	log.Error("pet_id:%d is not in res", pet_id)
	return pet_list, chip_list
}

func processChip(
	account string, chip_id int32, num int32,
	pet_list []*cspb.PetInfo,
	chip_list []*cspb.ChipInfo) ([]*cspb.PetInfo, []*cspb.ChipInfo) {

	res_chip := resmgr.ChipData.GetItems()[chip_id-1]

	//如果不是宠物碎片那就直接放到chip_list里面
	if res_chip.GetChipType() == 0 {
		chip_list = append(chip_list, makeChip(chip_id,
			res_chip.GetChipType(), num,
			int32(cspb.ChangeType_Add)))
		return pet_list, chip_list
	}

	chip_db_list, ret_chip := db.GetChipList(account)
	if ret_chip < 0 {
		log.Error("get chip db list error saccount:%s", account)
		return pet_list, chip_list
	}

	//随机到的碎片 自己是否拥有多少个
	var chip_db_num int32 = 0
	//var chip_type int32 = -1
	for _, chip_db := range chip_db_list {
		if chip_db.ChipId == chip_id {
			chip_db_num = chip_db.ChipNum
			break
		}
	}

	//该碎片需要多少个可以解锁该碎片对应的宠物
	start_star_level := resmgr.PetData.GetItems()[res_chip.GetChipType()-1].GetStartStarLevel()
	unlock_chip_num := getNeedChipNum(res_chip.GetChipType(), start_star_level)

	left_num := chip_db_num + num - unlock_chip_num
	change_chip_num := num
	//够解锁宠物的啦
	if left_num >= 0 {
		_, ret_pet := db.GetPetById(account, res_chip.GetChipType())
		if ret_pet < 0 {
			log.Error("get pet db error saccount:%s, pet_id:%d",
				account, res_chip.GetChipType())
			return pet_list, chip_list
		}
		if ret_pet == 1 {
			//db里面没有该宠物

			//通知客户端碎片的信息
			change_chip_num = num - unlock_chip_num

			//加上对应的宠物到pet_list
			pet_star_level := resmgr.PetData.GetItems()[res_chip.GetChipType()-1].GetStartStarLevel()
			pet_list = append(pet_list, makePet(res_chip.GetChipType(),
				1, 0, 0, pet_star_level))
		}
	}

	if change_chip_num > 0 {
		chip_list = append(chip_list, makeChip(chip_id,
			res_chip.GetChipType(), change_chip_num,
			int32(cspb.ChangeType_Add)))
	} else if change_chip_num < 0 {
		chip_list = append(chip_list, makeChip(chip_id,
			res_chip.GetChipType(), -change_chip_num,
			int32(cspb.ChangeType_Deduct)))
	}

	return pet_list, chip_list
}

func processChipEx(
	account string, chip_id int32, num int32,
	chip_list []*cspb.ChipInfo) []*cspb.ChipInfo {

	log.Debug("account:%s, chip_id:%d, num:%d",
		account, chip_id, num)
	if chip_id < 1 {
		log.Error("processChip fail chip_id(%d) is invalid ", chip_id)
		return chip_list
	}
	res_chip := resmgr.ChipData.GetItems()[chip_id-1]

	chip_list = append(chip_list, makeChip(chip_id,
		res_chip.GetChipType(), num,
		int32(cspb.ChangeType_Add)))
	return chip_list
}

func getNeedChipNum(pet_id int32, pet_star_lv int32) int32 {

	log.Debug("pet_id:%d, pet_star_lv:%d", pet_id, pet_star_lv)
	for _, star_data := range resmgr.StarData.GetItems() {
		if star_data.GetStarLevel() == pet_star_lv &&
			star_data.GetPetId() == pet_id {

			return star_data.GetNeedChipNum()
		}
	}
	return 0
}

func getPetMaxLvByStarLv(pet_id int32, pet_star_lv int32) int32 {

	for _, star_data := range resmgr.StarData.GetItems() {
		if star_data.GetStarLevel() == pet_star_lv &&
			star_data.GetPetId() == pet_id {

			return star_data.GetLevelLimitUp()
		}
	}
	return 0
}

func CheckMail(mail db.Mail) bool {
	valid := false
	if mail.AccessoryType == 1 ||
		mail.AccessoryType == 2 ||
		mail.AccessoryType == 3 {
		valid = true
	}
	if mail.AccessoryType == 4 {
		if mail.AccessorySubType > 0 &&
			mail.AccessorySubType < 12 &&
			mail.AccessoryNum > 0 {
			valid = true
		}
	} else if mail.AccessoryType == 5 {
		if mail.AccessorySubType > 0 &&
			mail.AccessoryNum > 0 &&
			mail.AccessorySubType <= int32(len(resmgr.PetData.GetItems())) {
			valid = true
		}

	} else if mail.AccessoryType == 6 {
		if mail.AccessorySubType > 0 &&
			mail.AccessoryNum > 0 &&
			mail.AccessorySubType <= int32(len(resmgr.ChipData.GetItems())) {
			valid = true
		}
	}
	return valid
}

func IpStringToInt(ip string) int64 {
	bits := strings.Split(ip, ".")
	if len(bits) < 3 {
		log.Debug("len:%d", len(bits))
		return 0
	}
	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	//b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 16
	sum += int64(b1) << 8
	sum += int64(b2)
	//sum += int64(b3)

	// sum += int64(b0) << 24
	// sum += int64(b1) << 16
	// sum += int64(b2) << 8
	// sum += int64(b3)

	return sum
}
