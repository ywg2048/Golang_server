package cs_handle

import cspb "protocol"
import proto "code.google.com/p/goprotobuf/proto"
import chip_db "tuojie.com/piggo/quickstart.git/db/collection"
import log "code.google.com/p/log4go"

func getChipListHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {

	chip_db_list, ret := chip_db.GetChipList(res_list.GetSAccount())
	if ret == 1 {
		return makeChipListResPkg(req, res_list, ret)
	} else if ret == 0 {
		var chip_list []*cspb.ChipInfo
		for _, db_info := range chip_db_list {
			chip_list = append(chip_list, makeChip(db_info.ChipId,
				db_info.ChipType,
				db_info.ChipNum,
				int32(cspb.ChangeType_Total)))
		}
		//添加pet_ntf到res_list中

		makeChipNtf(chip_list, res_list)
	} else {
		log.Error("get chip list fail ret:%d", ret)
	}

	return makeChipListResPkg(req, res_list, ret)
}

func makeChipListResPkg(req *cspb.CSPkg,
	res_list *cspb.CSPkgList, ret int32) int32 {

	//填充chiplistres回包
	res_data := new(cspb.CSChipListRes)
	*res_data = cspb.CSChipListRes{
		Ret: proto.Int32(ret),
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		ChipListRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kChipListRes),
		res_pkg_body, res_list)
	return ret
}
