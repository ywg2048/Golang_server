package cs_handle

import cspb "protocol"

import proto "code.google.com/p/goprotobuf/proto"

//import player_db "module/db/collection"
import log "code.google.com/p/log4go"
import resmgr "tuojie.com/piggo/quickstart.git/res_mgr"
import rand "github.com/tuojie/utility"

func randomItemHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {

	use_random_item_id := req.GetBody().GetRandomItemReq().GetUseRandomItemId()
	if use_random_item_id == 0 {
		cost_diamond := resmgr.KeyvalueData.GetItems()[7].GetIntValue()
		log.Debug("need deduct diamond:%d", cost_diamond)
		makeAttrInt32Ntf(int32(cspb.AttrId_Diamond),
			cost_diamond,
			int32(cspb.ChangeType_Deduct),
			res_list)
	} else {
		makeAttrItemNtf(int32(use_random_item_id),
			int32(1),
			int32(cspb.ChangeType_Deduct),
			res_list)

		// player_db.ChangeFieldValue(res_list.GetSAccount(),
		// 	"random_item_free_times", -1)
	}

	item_index := randomItemFromRes()
	if item_index <= 0 {
		log.Error("random item from res error item_index:%d", item_index)
		return makeRandomItemResPkg(req, res_list,
			int32(cspb.ErrorCode_RandomItemFail), item_index)
	}

	log.Debug("random item_index :%d", item_index)
	return makeRandomItemResPkg(req, res_list, int32(0), item_index)
}

func makeRandomItemResPkg(req *cspb.CSPkg,
	res_list *cspb.CSPkgList, ret int32, item_index int32) int32 {

	//填充RandomItemRes回包
	res_data := new(cspb.CSRandomItemRes)
	*res_data = cspb.CSRandomItemRes{
		Ret:       proto.Int32(ret),
		ItemIndex: proto.Int32(item_index),
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		RandomItemRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kRandomItemRes),
		res_pkg_body, res_list)
	return ret
}
func randomItemFromRes() int32 {

	var max_num int32 = 0
	for _, data := range resmgr.RandomitemData.GetItems() {
		max_num += data.GetRandomNum()
	}

	var random_num int32 = rand.RandInt32(max_num)
	var cur_num int32
	var item_index int32 = -1

	for _, data := range resmgr.RandomitemData.GetItems() {
		cur_num += data.GetRandomNum()
		if cur_num >= random_num {
			item_index = data.GetIndex()
			break
		}
	}
	if item_index == -1 {
		log.Error("random item error max_num:%d, random_num:%d, cur_num:%d",
			max_num, random_num, cur_num)
	}
	log.Debug("random_item_index:%d", item_index)
	return item_index
}
