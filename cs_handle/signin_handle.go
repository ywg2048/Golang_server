package cs_handle

import cspb "protocol"
import proto "code.google.com/p/goprotobuf/proto"
import db "tuojie.com/piggo/quickstart.git/db/collection"
import log "code.google.com/p/log4go"
import resmgr "tuojie.com/piggo/quickstart.git/res_mgr"

//import "resource"
import "time"

func signinHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {

	log.Debug("******signinHandle, req is %v, res is %v", req, res_list)
	req_flag := req.GetBody().GetSignInReq().GetFlag()

	//拉取player基本信息
	//	db_info, ret := db.InsertPlayer(res_list.GetSAccount())
	ret, db_info := db.LoadPlayer(res_list.GetCAccount(), res_list.GetSAccount(), res_list.GetUid())
	log.Debug("-----------db_info---------", db_info)
	if ret != 0 {
		log.Error("load player fail ret:%d", ret)
		return makeSigninResPkg(req, res_list,
			int32(cspb.ErrorCode_PlayerNotExist), req_flag, int32(-1))
	}

	signin_day := int32(0)
	cur_day := int32(-1)
	consume := int32(-1)
	free := int32(0)

	start_time_str := resmgr.KeyvalueData.GetItems()[1].GetStrValue()

	Time, _ := time.Parse("2006-01-02 15:04:05", start_time_str)
	start_time := Time.Unix() - int64(60*60*8)
	now_time := time.Now().Unix()
	log.Debug("^^^^^^db_info.LastSignIntime^^^^^^", db_info.WonderfulFriends.LastSignInTime)
	var last_signin_time int64
	var free_oper_time int64
	if db_info.LastSignInTime > db_info.WonderfulFriends.LastSignInTime {
		last_signin_time = db_info.LastSignInTime
		free_oper_time = db_info.FreeSignInOperTime
	} else {
		last_signin_time = db_info.WonderfulFriends.LastSignInTime
		free_oper_time = db_info.WonderfulFriends.FreeSignInOperTime
	}
	if last_signin_time > now_time || last_signin_time <= start_time {
		last_signin_time = start_time
	}

	last_signin_cycle := (last_signin_time - start_time) / 86400 / 30
	last_signin_day := (last_signin_time - start_time) / 86400 % 30
	now_cycle := (now_time - start_time) / 86400 / 30
	now_day := (now_time - start_time) / 86400 % 30

	log.Debug("last_signin_time:%d", last_signin_time)
	log.Debug("free_oper_time:%d", free_oper_time)
	log.Debug("now_time:%d", now_time)
	log.Debug("start_time:%d", start_time)
	log.Debug("last_signin_cycle:%d", last_signin_cycle)
	log.Debug("last_signin_day:%d", last_signin_day)
	log.Debug("now_cycle:%d", now_cycle)
	log.Debug("now_day:%d", now_day)

	can_signin := false

	for true {
		//不能签到
		if last_signin_cycle == now_cycle {
			if last_signin_day >= now_day {
				if last_signin_day == now_day && now_day == 0 {
					signin_day = int32(-1)
					cur_day = int32(0)
					can_signin = true
					break
				} else {
					//不能签到了，已经全部都签完了
					signin_day = int32(now_day)
					cur_day = int32(now_day)
					can_signin = false
					break
				}
			} else {
				//同一个周期,此时now_day大于last_signin_day
				signin_day = int32(last_signin_day)
				cur_day = int32(now_day)
				can_signin = true
				break
			}
		}

		if last_signin_cycle != now_cycle {
			signin_day = int32(-1)
			cur_day = int32(now_day)
			can_signin = true
			break
		}

		log.Error("oh fuck code has bug!!!!")
		break
	}

	if can_signin {
		//免费签到
		if now_time-free_oper_time >= int64(60*60*24) || signin_day == -1 {
			consume = int32(0)
			free = int32(1)
		} else { //还可以签到
			//花钻石签到
			log.Debug("signin_day:%d", signin_day)
			consume = resmgr.DailyloginData.GetItems()[signin_day].GetReSignInDiamond()
		}
	}

	if req_flag == 0 {
		//查询
		makeSigninNtf(req, res_list, signin_day+1, cur_day+1,
			consume, free_oper_time+int64(60*60*24)+int64(60*60*8))
		makeSigninResPkg(req, res_list, int32(0), req_flag, free)
		return int32(0)
	}
	//签到
	if can_signin {
		//保存DB
		if consume == 0 {
			free_oper_time = int64(
				start_time +
					int64(86400*30*now_cycle) +
					int64(86400*(now_day)))
		}
		//可以签到
		signin_day++
		reward(signin_day, res_list)
		makeSigninNtf(req, res_list, signin_day+1, cur_day+1,
			consume, free_oper_time+int64(60*60*24)+int64(60*60*8))
		makeAttrInt32Ntf(int32(cspb.AttrId_Diamond), consume,
			int32(cspb.ChangeType_Deduct), res_list)
		makeSigninResPkg(req, res_list, int32(0), req_flag, free)

		last_signin_time = int64(
			start_time +
				int64(86400*30*now_cycle) +
				int64(86400*(signin_day)))
		//存时间到数据表player
		ret = db.SetLastSignInTime(res_list.GetCAccount(), last_signin_time, free_oper_time)
		if ret != 0 {
			log.Error("SetLastSignInTime error ret:%d", ret)
		}
	} else {
		//不可以签到
		makeSigninNtf(req, res_list, signin_day+1, cur_day+1,
			consume, free_oper_time+int64(60*60*24)+int64(60*60*8))
		makeSigninResPkg(req, res_list, int32(-1), req_flag, free)
	}

	return int32(0)
}

func reward(signin_day int32, res_list *cspb.CSPkgList) {
	log.Debug("*****Signin_handle.reward:signin_day:%d", signin_day)

	if signin_day < 0 || signin_day >= 30 {
		log.Error("signin:%d is invalid", signin_day)
		return
	}

	var pet_list []*cspb.PetInfo
	var chip_list []*cspb.ChipInfo

	login_data := resmgr.DailyloginData.GetItems()[signin_day]
	for _, goods_info := range login_data.GetGoodsList() {
		if goods_info.GetType() < int32(cspb.AttrId_Diamond) ||
			goods_info.GetType() >= int32(cspb.AttrId_Max) {

			log.Error("goods_type:%d is invalid", goods_info.GetType())
			break
		}
		if goods_info.GetType() == int32(cspb.AttrId_Item) {
			makeAttrItemNtf(goods_info.GetSubType(), goods_info.GetNum(),
				int32(cspb.ChangeType_Add), res_list)
		} else if goods_info.GetType() == int32(8) {
			//碎片
			pet_list, chip_list = processChip(res_list.GetSAccount(),
				goods_info.GetSubType(), goods_info.GetNum(),
				pet_list, chip_list)

		} else if goods_info.GetType() == int32(9) {
			//宠物
			pet_list, chip_list = processPet(res_list.GetSAccount(),
				goods_info.GetSubType(), goods_info.GetNum(),
				pet_list, chip_list)
		} else {
			makeAttrInt32Ntf(goods_info.GetType(), goods_info.GetNum(),
				int32(cspb.ChangeType_Add), res_list)
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

func makeSigninNtf(req *cspb.CSPkg, res_list *cspb.CSPkgList,
	signin_day int32,
	cur_day int32,
	consume int32,
	free_time int64) int32 {

	//填充SigninNtf回包
	res_data := new(cspb.CSSignInNtf)
	*res_data = cspb.CSSignInNtf{
		SigninDay: proto.Int32(signin_day),
		CurDay:    proto.Int32(cur_day),
		Consume:   proto.Int32(consume),
		FreeTime:  proto.Int64(free_time),
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		SignInNtf: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kSignInNtf),
		res_pkg_body, res_list)
	return int32(0)
}

func makeSigninResPkg(req *cspb.CSPkg,
	res_list *cspb.CSPkgList, ret int32, flag int32, free int32) int32 {

	//填充SigninRes回包
	res_data := new(cspb.CSSignInRes)
	*res_data = cspb.CSSignInRes{
		Ret:  proto.Int32(ret),
		Free: proto.Int32(free),
		Flag: proto.Int32(flag),
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		SignInRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kSignInRes),
		res_pkg_body, res_list)
	return ret
}
