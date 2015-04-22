package cs_handle

import cspb "protocol"
import proto "code.google.com/p/goprotobuf/proto"
import db "tuojie.com/piggo/quickstart.git/db/collection"
import log "code.google.com/p/log4go"

func registPlayerHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {

	log.Debug("******registPlayerHandle, req is %v, res is %v", req, res_list)
	//req_data := req.GetBody().GetRegistPlayerReq()

	// var player db.Player
	// player.Caccount = res_list.GetCAccount()
	// player.RegistTime = time.Now().Unix()

	// uid := db.GetUid()
	// if uid < 0 {
	// 	log.Error("get uid fail uid:%d", uid)
	// 	ret_uid := int32(cspb.ErrorCode_PlayerInsertFail)
	// 	uid = int64(0)
	// 	return makeRegistPlayerResPkg(req, res_list, ret_uid, "", uid)
	// }
	// player.Uid = uid
	// 根据提供的client account创建一个player用户
	ret, player := db.LoadPlayer(res_list.GetCAccount(), res_list.GetSAccount(), res_list.GetUid())
	log.Debug("Db LoadPlayer Player, ret is %d, player is %v", ret, player)

	// switch ret {
	// case 0:
	// 	//		initOther(saccount)
	// 	ret_uid := int32(0)
	// case 1:
	// 	log.Error("insert player is exist")
	// 	ret_player = int32(cspb.ErrorCode_PlayerSAccountIsExist)
	// 	ret_uid := ret
	// case 2:
	// 	log.Error("get uid fail uid:%d", uid)
	// 	ret_uid := int32(cspb.ErrorCode_PlayerInsertFail)
	// 	player.Uid = int64(0)
	// }

	return makeRegistPlayerResPkg(req, res_list, ret, player.Saccount.String(), player.Uid)

	// if ret == 1 {
	// 	log.Error("insert player is exist")
	// 	ret_player = int32(cspb.ErrorCode_PlayerSAccountIsExist)
	// 	return makeRegistPlayerResPkg(req, res_list, ret_player, saccount, uid)
	// } else if ret != 0 {
	// 	log.Error("insert player fail ret:%d", ret_player)
	// 	ret_player = int32(cspb.ErrorCode_PlayerInsertFail)
	// 	return makeRegistPlayerResPkg(req, res_list, ret_player, saccount, uid)
	// }

	// if ret == 0 {
	// 	initOther(saccount)
	// }
	// return makeRegistPlayerResPkg(req, res_list, int32(0), saccount, uid)
}

func makeRegistPlayerResPkg(req *cspb.CSPkg,
	res_list *cspb.CSPkgList, ret int32, saccount string, uid int64) int32 {

	//填充RegistPlayerRes回包
	res_data := new(cspb.CSRegistPlayerRes)
	*res_data = cspb.CSRegistPlayerRes{
		Ret: proto.Int32(ret),
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		RegistPlayerRes: res_data,
	}
	*res_list.SAccount = saccount
	*res_list.Uid = uid

	res_list = makeCSPkgList(int32(cspb.Command_kRegistPlayerRes),
		res_pkg_body, res_list)
	return ret
}

func initOther(account string) {
	//db.SetPetInfo(account, 1, 1, 0, 0, 1)
}
