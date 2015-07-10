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
	pet_star_level int32, Petmedallevel int32, PetmedalNum int32, DressId int32, PetCardNtf []*cspb.CSPetCardNtf) *cspb.PetInfo {

	pet_info := new(cspb.PetInfo)
	*pet_info = cspb.PetInfo{
		PetId:        proto.Int32(pet_id),
		PetLevel:     proto.Int32(pet_level),
		PetCurExp:    proto.Int32(pet_cur_exp),
		PetTotalExp:  proto.Int32(pet_total_exp),
		PetStarLevel: proto.Int32(pet_star_level),
		PetMedalLevl: proto.Int32(Petmedallevel),
		PetMedalNum:  proto.Int32(PetmedalNum),
		DressId:      proto.Int32(DressId),
		PetCardNtf:   PetCardNtf,
	}

	beego.Debug("pet_info:%v", pet_info)
	return pet_info
}
func makecardNtf(CardId int32, CardNum int32) *cspb.CSPetCardNtf {
	card_ntf := new(cspb.CSPetCardNtf)
	*card_ntf = cspb.CSPetCardNtf{
		CardId: proto.Int32(CardId),
		CarNum: proto.Int32(CardNum),
	}
	return card_ntf
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
func makeRank(playuid int32, playname string,
	level int32, medalnum int32, rankid int32, medalLevelID int32, starid int32) *cspb.CSRankNtf {

	rank_ntf := new(cspb.CSRankNtf)
	*rank_ntf = cspb.CSRankNtf{
		Playuid:  proto.Int32(playuid),
		Playname: proto.String(playname),

		Level:        proto.Int32(level),
		Medalnum:     proto.Int32(medalnum),
		Rankid:       proto.Int32(rankid),
		MedalLevelID: proto.Int32(medalLevelID),
		Starid:       proto.Int32(starid),
	}

	return rank_ntf
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
func makeFriendntf(Playid int32, Starid int32, Name string, CardNtf []*cspb.CSPetCardNtf) *cspb.CSFriendNtf {
	friend_ntf := new(cspb.CSFriendNtf)
	*friend_ntf = cspb.CSFriendNtf{
		Playid:  proto.Int32(Playid),
		Starid:  proto.Int32(Starid),
		Name:    proto.String(Name),
		CardNtf: CardNtf,
	}
	return friend_ntf
}
func makeFriendmessagelistNtf(MessageType int32, ElementType int32, Friendntf []*cspb.CSFriendNtf, MessageId int32) *cspb.CSFriendmessagelistNtf {
	friendmessagelist_ntf := new(cspb.CSFriendmessagelistNtf)
	*friendmessagelist_ntf = cspb.CSFriendmessagelistNtf{
		MessageType: proto.Int32(MessageType),
		ElementType: proto.Int32(ElementType),
		Friendntf:   Friendntf,

		MessageId: proto.Int32(MessageId),
	}
	return friendmessagelist_ntf

}
func makefriendlist(FriendListId int32, Playerid int32, Name string, Starid int32, Starname string, Fighting int32, DressId int32, Dress string, Level int32, Medal int32, MedalLevelID int32, Stagelevel int32) *cspb.CSFriendListNtf {
	friendlist_ntf := new(cspb.CSFriendListNtf)
	*friendlist_ntf = cspb.CSFriendListNtf{
		FriendListId: proto.Int32(FriendListId),
		Playerid:     proto.Int32(Playerid),
		Name:         proto.String(Name),
		Starid:       proto.Int32(Starid),
		Starname:     proto.String(Starname),
		Fighting:     proto.Int32(Fighting),
		DressId:      proto.Int32(DressId),
		Dress:        proto.String(Dress),
		Level:        proto.Int32(Level),
		Medal:        proto.Int32(Medal),
		MedalLevelID: proto.Int32(MedalLevelID),
		Stagelevel:   proto.Int32(Stagelevel),
	}
	return friendlist_ntf
}

func makeAnimalNtf(AnimalId int32, Status int32) *cspb.CSAnimalNtf {
	animal_ntf := new(cspb.CSAnimalNtf)
	*animal_ntf = cspb.CSAnimalNtf{

		AnimalId: proto.Int32(AnimalId),
		Status:   proto.Int32(Status),
	}
	return animal_ntf
}

func makeStarInfo(StarId int32, DressId int32) *cspb.CSStarInfo {
	star_info := new(cspb.CSStarInfo)
	*star_info = cspb.CSStarInfo{
		StarId:  proto.Int32(StarId),
		DressId: proto.Int32(DressId),
	}
	return star_info
}
func makeFriendStagentf(FriendId int32, Name string, Stage int32) *cspb.CSFriendStageNtf {
	friendstage_ntf := new(cspb.CSFriendStageNtf)
	*friendstage_ntf = cspb.CSFriendStageNtf{
		FriendId: proto.Int32(FriendId),
		Name:     proto.String(Name),
		Stage:    proto.Int32(Stage),
	}
	return friendstage_ntf
}
func makeAchievementNtf(Achievementid int32, Status int32) *cspb.CSAchievementNtf {
	achievement_ntf := new(cspb.CSAchievementNtf)
	*achievement_ntf = cspb.CSAchievementNtf{
		Achievementid: proto.Int32(Achievementid),
		Status:        proto.Int32(Status),
	}
	return achievement_ntf
}
func makemessageTipsntf(Messagetype int32, Number int32) *cspb.CSmessageTipsntf {
	messageTips_ntf := new(cspb.CSmessageTipsntf)
	*messageTips_ntf = cspb.CSmessageTipsntf{
		Messagetype: proto.Int32(Messagetype),
		Number:      proto.Int32(Number),
	}
	return messageTips_ntf
}
