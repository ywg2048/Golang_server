package cs_handle

import (
	"github.com/astaxie/beego"
	// models "tuojie.com/piggo/quickstart.git/models"
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
	pet_cur_exp int32, Satisfaction int32,
	FightExp int32, Fighting int32, DressId int32) *cspb.PetInfo {

	pet_info := new(cspb.PetInfo)
	*pet_info = cspb.PetInfo{
		PetId:        proto.Int32(pet_id),
		PetLevel:     proto.Int32(pet_level),
		PetCurExp:    proto.Int32(pet_cur_exp),
		Satisfaction: proto.Int32(Satisfaction),
		FightExp:     proto.Int32(FightExp),
		Fighting:     proto.Int32(Fighting),
		DressId:      proto.Int32(DressId),
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
func makecardNtf1(CardId int32, CardNum int32) *cspb.CSCardNtf {
	card_ntf := new(cspb.CSCardNtf)
	*card_ntf = cspb.CSCardNtf{
		CardId:  proto.Int32(CardId),
		CardNum: proto.Int32(CardNum),
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
func makeFriendntf(Playid int32, Starid int32, Name string, MysqlId int32) *cspb.CSFriendNtf {
	friend_ntf := new(cspb.CSFriendNtf)
	*friend_ntf = cspb.CSFriendNtf{
		Playid:  proto.Int32(Playid),
		Starid:  proto.Int32(Starid),
		Name:    proto.String(Name),
		MysqlId: proto.Int32(MysqlId),
	}
	return friend_ntf
}
func makeFriendmessagelistNtf(MessageType int32, ElementType int32, CardId int32, ElementNum int32, Friendntf []*cspb.CSFriendNtf, MessageId int32) *cspb.CSFriendmessagelistNtf {
	friendmessagelist_ntf := new(cspb.CSFriendmessagelistNtf)
	*friendmessagelist_ntf = cspb.CSFriendmessagelistNtf{
		MessageType: proto.Int32(MessageType),
		ElementType: proto.Int32(ElementType),
		Friendntf:   Friendntf,
		CardId:      proto.Int32(CardId),
		ElementNum:  proto.Int32(ElementNum),
		MessageId:   proto.Int32(MessageId),
	}
	return friendmessagelist_ntf

}
func makefriendlist(FriendListId int32, Playerid int32, Name string, Starid int32, Starname string, Fighting int32, DressId int32, Dress string, Level int32, Medal int32, MedalLevelID int32, Stagelevel int32, iZooId int32, iAchievementID int32) *cspb.CSFriendListNtf {
	friendlist_ntf := new(cspb.CSFriendListNtf)
	*friendlist_ntf = cspb.CSFriendListNtf{
		FriendListId:   proto.Int32(FriendListId),
		Playerid:       proto.Int32(Playerid),
		Name:           proto.String(Name),
		Starid:         proto.Int32(Starid),
		Starname:       proto.String(Starname),
		Fighting:       proto.Int32(Fighting),
		DressId:        proto.Int32(DressId),
		Dress:          proto.String(Dress),
		Level:          proto.Int32(Level),
		Medal:          proto.Int32(Medal),
		MedalLevelID:   proto.Int32(MedalLevelID),
		Stagelevel:     proto.Int32(Stagelevel),
		IZooID:         proto.Int32(iZooId),
		IAchievementID: proto.Int32(iAchievementID),
	}
	return friendlist_ntf
}

func makeAnimalNtf(AnimalId int32, AnimalLevel int32, Islocked int32) *cspb.CSAnimalNtf {
	animal_ntf := new(cspb.CSAnimalNtf)
	*animal_ntf = cspb.CSAnimalNtf{

		AnimalId:    proto.Int32(AnimalId),
		AnimalLevel: proto.Int32(AnimalLevel),
		Islocked:    proto.Int32(Islocked),
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
func makeAchievementNtf(Achievementid int32, StarLevel int32, Process int32, IsReceive int32) *cspb.CSAchievementNtf {
	achievement_ntf := new(cspb.CSAchievementNtf)
	*achievement_ntf = cspb.CSAchievementNtf{
		Achievementid: proto.Int32(Achievementid),
		Process:       proto.Int32(Process),
		StarLevel:     proto.Int32(StarLevel),
		IsReceive:     proto.Int32(IsReceive),
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
func makeAttrInfo(Id int32, Value *cspb.AttrValue, ChangeType int32) *cspb.AttrInfo {
	attrinfo := new(cspb.AttrInfo)
	*attrinfo = cspb.AttrInfo{
		Id:         proto.Int32(Id),
		Value:      Value,
		ChangeType: proto.Int32(ChangeType),
	}
	return attrinfo
}
func makeStageNtf(StageId int32, StageLevel int32, StageScore int32, GetMedal int32, MedalIsAdd int32) *cspb.CSStageNtf {
	stage_ntf := new(cspb.CSStageNtf)
	*stage_ntf = cspb.CSStageNtf{
		StageId:    proto.Int32(StageId),
		StageLevel: proto.Int32(StageLevel),
		StageScore: proto.Int32(StageScore),
		GetMedal:   proto.Int32(GetMedal),
		MedalIsAdd: proto.Int32(MedalIsAdd),
	}
	return stage_ntf
}

// func makeAttrValue(Diamond int32, Gold int32, Flower int32, Solution int32, Medal int32) *cspb.AttrValue {
// 	attrvalue := new(cspb.AttrValue)
// 	*attrvalue = cspb.AttrValue{
// 		Diamond:  proto.Int32(Diamond),
// 		Gold:     proto.Int32(Gold),
// 		Flower:   proto.Int32(Flower),
// 		Solution: proto.Int32(Solution),
// 		Medal:    proto.Int32(Medal),
// 	}
// 	return attrvalue
// }
