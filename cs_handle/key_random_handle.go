package cs_handle

import (
	"github.com/astaxie/beego"
)
import cspb "protocol"
import proto "code.google.com/p/goprotobuf/proto"
import db "tuojie.com/piggo/quickstart.git/db/collection"

import resmgr "tuojie.com/piggo/quickstart.git/res_mgr"
import rand "github.com/tuojie/utility"
import "resource"
import "time"
import "strings"

var pool_goods_num int = 16

type assignInfo struct {
	Index    int32
	PoolType int32
	GoodsId  int32
}

func keyRandomHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {

	//检查随机类型是否合法
	random_type := req.GetBody().GetKeyRandomReq().GetRandomType()
	if !checkRandomType(random_type) {
		makeKeyRandomResPkg([]int32{}, []int32{}, req, res_list,
			int32(cspb.ErrorCode_KeyRandomTypeInvalid),
			random_type)
		return -1
	}
	str_version_s := req.GetBody().GetKeyRandomReq().GetVersion()
	str_version := strings.Replace(str_version_s, "Version ", "", -1)
	int_version := IpStringToInt(str_version)
	if int_version == 0 {
		int_version = IpStringToInt("1.0.0")
	}
	beego.Debug("req_str_version:%s, req_int_version:%d", str_version, int_version)
	//拉取pool_random信息
	pool_db_list, ret := db.GetRandomPool(res_list.GetSAccount(),
		getKeyRandomPoolType(random_type))
	if ret != 0 {
		beego.Error("get random pools fail ret:%d", ret)
		makeKeyRandomResPkg([]int32{}, []int32{}, req, res_list,
			int32(cspb.ErrorCode_PlayerNotExist), random_type)
		return -1
	}

	var goods_list []int32
	var next_goods_list []int32
	var goods_id int32 = -1

	assign_list, last_signin_time := makeAssignList(pool_db_list)
	if isPoolRandom(random_type) {
		now := time.Now().Unix()
		free, hour, today := isFreeChangePool(random_type, last_signin_time)
		if !free {
			//通知钻石变化
			cost_diamond := resmgr.KeyvalueData.GetItems()[0].GetIntValue()
			makeAttrInt32Ntf(int32(cspb.AttrId_Diamond),
				cost_diamond,
				int32(cspb.ChangeType_Deduct),
				res_list)
		} else {
			beego.Debug("free random key")
		}
		if hour >= 0 && today >= 0 {
			makePoolNextFreeNtf(hour, today, res_list)
		}
		last_signin_time = now

	} else {
		if len(assign_list) < pool_goods_num {
			beego.Error("assign_list is error len:%d", int(len(assign_list)))
			makeKeyRandomResPkg([]int32{}, []int32{}, req, res_list,
				int32(cspb.ErrorCode_SysError),
				random_type)
			return -1
		}
		var random_num int = 0
		if isBatchRandom(random_type) {
			random_num = int(resmgr.KeyvalueData.GetItems()[3].GetIntValue())
		} else {
			random_num = int(resmgr.KeyvalueData.GetItems()[2].GetIntValue())
		}

		for i := 0; i < random_num; i++ {
			var index int32
			goods_id, index = randomFromAssign(random_type, assign_list)
			if goods_id != -1 {
				goods_list = append(goods_list, goods_id)
				assign_list = deleteAssignList(res_list.GetSAccount(), index,
					goods_id, last_signin_time, assign_list)
			} else {
				beego.Error("random goods error")
				makeKeyRandomResPkg([]int32{}, []int32{}, req, res_list,
					int32(cspb.ErrorCode_SysError),
					random_type)
				return -1
			}
		}

		//通知扣掉钥匙
		makeAttrInt32Ntf(getAttrIdFromRandomType(random_type), int32(random_num),
			int32(cspb.ChangeType_Deduct), res_list)
		//发放 chip 或者 pet ntf
		goodsParseMakeNtf(goods_list, res_list)
	}

	for i := 0; i < pool_goods_num; i++ {
		goods_id = randomFromAll(random_type, int_version, defaultSelecter, randomKeyNum)
		if goods_id == -1 {
			beego.Error("random error")
			makeKeyRandomResPkg([]int32{}, []int32{}, req, res_list,
				int32(cspb.ErrorCode_SysError),
				random_type)
			return -1
		}
		next_goods_list = append(next_goods_list, goods_id)
		if isPoolRandom(random_type) {
			goods_list = append(goods_list, goods_id)
		}
	}
	coverKeyRandomInfo(res_list.GetSAccount(), getKeyRandomPoolType(random_type),
		next_goods_list, last_signin_time)

	return makeKeyRandomResPkg(next_goods_list, goods_list, req,
		res_list, 0, random_type)
}

func checkRandomType(random_type int32) bool {
	if random_type == int32(cspb.KeyRandomType_rCopperKey) ||
		random_type == int32(cspb.KeyRandomType_rSilverKey) ||
		random_type == int32(cspb.KeyRandomType_rGoldKey) ||
		random_type == int32(cspb.KeyRandomType_rCopperKeyPool) ||
		random_type == int32(cspb.KeyRandomType_rSilverKeyPool) ||
		random_type == int32(cspb.KeyRandomType_rGoldKeyPool) ||
		random_type == int32(cspb.KeyRandomType_rSilverKeyBatch) ||
		random_type == int32(cspb.KeyRandomType_rGoldKeyBatch) {
		return true
	}
	beego.Error("random type is invalid random_type:%d", random_type)
	return false
}

func getKeyRandomPoolType(random_type int32) int32 {
	switch random_type {
	case int32(cspb.KeyRandomType_rCopperKey):
		return int32(cspb.KeyRandomType_rCopperKeyPool)

	case int32(cspb.KeyRandomType_rSilverKey):
		return int32(cspb.KeyRandomType_rSilverKeyPool)

	case int32(cspb.KeyRandomType_rGoldKey):
		return int32(cspb.KeyRandomType_rGoldKeyPool)

	case int32(cspb.KeyRandomType_rCopperKeyPool):
		return int32(cspb.KeyRandomType_rCopperKeyPool)

	case int32(cspb.KeyRandomType_rSilverKeyPool):
		return int32(cspb.KeyRandomType_rSilverKeyPool)

	case int32(cspb.KeyRandomType_rGoldKeyPool):
		return int32(cspb.KeyRandomType_rGoldKeyPool)

	case int32(cspb.KeyRandomType_rSilverKeyBatch):
		return int32(cspb.KeyRandomType_rSilverKeyPool)

	case int32(cspb.KeyRandomType_rGoldKeyBatch):
		return int32(cspb.KeyRandomType_rGoldKeyPool)

	default:
		beego.Error(" random_type is invalid random_type:%d",
			random_type)
		return -1
	}
}

func isPoolRandom(random_type int32) bool {
	if random_type == int32(cspb.KeyRandomType_rCopperKeyPool) ||
		random_type == int32(cspb.KeyRandomType_rSilverKeyPool) ||
		random_type == int32(cspb.KeyRandomType_rGoldKeyPool) {
		return true
	}
	return false
}

func isFreeChangePool(random_type int32, last_signin_time int64) (bool, int32, int32) {

	var hour_list []int32
	if random_type == int32(cspb.KeyRandomType_rCopperKeyPool) {
		hour_list = resmgr.KeyvalueData.GetItems()[4].GetIntListValue()
	} else if random_type == int32(cspb.KeyRandomType_rSilverKeyPool) {
		hour_list = resmgr.KeyvalueData.GetItems()[5].GetIntListValue()
	} else if random_type == int32(cspb.KeyRandomType_rGoldKeyPool) {
		hour_list = resmgr.KeyvalueData.GetItems()[6].GetIntListValue()
	}

	size := len(hour_list)
	if size <= 0 {
		beego.Debug("hour_list`len:%d is invalid", size)
		beego.Debug("return %t", false)
		return false, -1, -1
	}

	last_hour := int32(time.Unix(last_signin_time, 0).Hour())
	last_year_day := int32(time.Unix(last_signin_time, 0).YearDay())
	now_hour := int32(time.Now().Hour())
	now_year_day := int32(time.Now().YearDay())

	beego.Debug("hour_list:%d", hour_list)
	beego.Debug("last_hour:%d", last_hour)
	beego.Debug("last_year_day:%d", last_year_day)
	beego.Debug("now_hour:%d", now_hour)
	beego.Debug("now_year_day:%d", now_year_day)

	if last_signin_time == int64(0) {
		beego.Debug("last_signin_time:%d return %t", last_signin_time, true)
		return true, hour_list[0], int32(1)
	}
	var ret bool

	if last_year_day != now_year_day {
		//不在同一天
		if now_hour >= hour_list[0] {
			//当前的小时数大于等于本天第一次就可以免费
			ret = true
		} else {
			beego.Debug("return %t", false)
			ret = false
		}

		if len(hour_list) >= 1 {
			beego.Debug("ret:%t, hour:%d, today:%d", ret, hour_list[1], 1)
			return ret, hour_list[1], int32(1)
		} else {
			beego.Debug("ret:%t, hour:%d, today:%d", ret, hour_list[0], 0)
			return ret, hour_list[0], int32(0)
		}

	} else {
		//在同一天
		for _, hour := range hour_list {
			if hour <= 0 || hour >= 24 {
				beego.Error("static data is invalid hour:%d", hour)
				beego.Debug("return %t", false)
				return false, -1, int32(-1)
			}

			if last_hour < hour && now_hour >= hour {
				beego.Debug("return %t", true)
				return true, hour, int32(1)
			}
		}
	}
	return false, -1, int32(-1)
}

func isBatchRandom(random_type int32) bool {
	if random_type == int32(cspb.KeyRandomType_rGoldKeyBatch) ||
		random_type == int32(cspb.KeyRandomType_rSilverKeyBatch) {
		return true
	}
	return false
}

func randomFromAssign(random_type int32, assign_list []*assignInfo) (int32, int32) {

	beego.Debug("assign_list:%v,", assign_list)

	var max_num int32 = 0
	for _, assign_goods := range assign_list {
		data := resmgr.RandomgoodsData.GetItems()[assign_goods.GoodsId-1]
		max_num += randomKeyNum(random_type, data)
	}
	//log.Debug(">>> max_num:%d", max_num)
	var random_num int32 = rand.RandInt32(max_num)
	var cur_num int32
	var goods_id int32 = -1
	var index int32

	for _, assign_goods := range assign_list {
		data := resmgr.RandomgoodsData.GetItems()[assign_goods.GoodsId-1]
		cur_num += randomKeyNum(random_type, data)
		if cur_num >= random_num {
			goods_id = data.GetElementId()
			index = assign_goods.Index
			break
		}
	}
	if goods_id == -1 {
		beego.Error("random error max_num:%d, random_num:%d, cur_num:%d",
			max_num, random_num, cur_num)
	}
	beego.Debug("random_type:%d, random_goods_id:%d, index:%d",
		random_type, goods_id, index)
	return goods_id, index
}

func randomKeyNum(random_type int32, data *resource.Randomgoods) int32 {
	switch random_type {
	case int32(cspb.KeyRandomType_rCopperKey):
		return data.GetCopperKeyRandom()

	case int32(cspb.KeyRandomType_rSilverKey):
		return data.GetSilverKeyRandom()

	case int32(cspb.KeyRandomType_rGoldKey):
		return data.GetGoldKeyRandom()

	case int32(cspb.KeyRandomType_rCopperKeyPool):
		return data.GetCopperKeyRandom()

	case int32(cspb.KeyRandomType_rSilverKeyPool):
		return data.GetSilverKeyRandom()

	case int32(cspb.KeyRandomType_rGoldKeyPool):
		return data.GetGoldKeyRandom()

	case int32(cspb.KeyRandomType_rSilverKeyBatch):
		return data.GetSilverKeyRandom()

	case int32(cspb.KeyRandomType_rGoldKeyBatch):
		return data.GetGoldKeyRandom()

	default:
		beego.Error(" random_type is invalid random_type:%d",
			random_type)
		return 0
	}
}

func makeAssignList(key_db_list []db.RandomPool) ([]*assignInfo, int64) {
	var assign_list []*assignInfo
	var last_signin_time int64
	for _, info := range key_db_list {
		//if info.HasRandom == 1 {
		//	continue
		//}
		assign_info := new(assignInfo)
		assign_info.GoodsId = info.GoodsId
		assign_info.Index = info.Index
		assign_info.PoolType = info.PoolType

		assign_list = append(assign_list, assign_info)

		last_signin_time = info.UpdateTime
	}
	beego.Debug("assign_list:%v, last_signin_time:%d",
		assign_list, last_signin_time)
	return assign_list, last_signin_time
}

func deleteAssignList(account string, index int32,
	goods_id int32, last_signin_time int64,
	assign_list []*assignInfo) []*assignInfo {

	var return_list []*assignInfo
	for _, assign_info := range assign_list {
		if assign_info.GoodsId == goods_id &&
			assign_info.Index == index {
			db.SetRandomPool(account, index, assign_info.PoolType,
				goods_id, last_signin_time, int32(1))
		} else {
			return_list = append(return_list, assign_info)
		}
	}
	return return_list
}
func coverKeyRandomInfo(account string, random_type int32,
	goods_list []int32, last_signin_time int64) int32 {

	beego.Debug("account:%s, random_type:%d, goods_list:%v",
		account, random_type, goods_list)

	if random_type != int32(cspb.KeyRandomType_rCopperKeyPool) &&
		random_type != int32(cspb.KeyRandomType_rSilverKeyPool) &&
		random_type != int32(cspb.KeyRandomType_rGoldKeyPool) {

		beego.Error("random_type:%d is invalid", random_type)
		return -1
	}
	count := len(goods_list)
	if count != pool_goods_num {
		beego.Error("goods count:%d", count)
		return -1
	}

	for i := 0; i < pool_goods_num; i++ {
		db.SetRandomPool(account, int32(i), random_type,
			goods_list[i], last_signin_time, int32(0))
	}
	return 0
}

func getAttrIdFromRandomType(random_type int32) int32 {
	switch random_type {
	case int32(cspb.KeyRandomType_rCopperKey):
		return int32(cspb.AttrId_CopperKey)

	case int32(cspb.KeyRandomType_rSilverKey):
		return int32(cspb.AttrId_SilverKey)

	case int32(cspb.KeyRandomType_rGoldKey):
		return int32(cspb.AttrId_GoldKey)

	case int32(cspb.KeyRandomType_rSilverKeyBatch):
		return int32(cspb.AttrId_SilverKey)

	case int32(cspb.KeyRandomType_rGoldKeyBatch):
		return int32(cspb.AttrId_GoldKey)

	default:
		beego.Error(" random_type is invalid random_type:%d",
			random_type)
		return 0
	}
}

func makeKeyRandomResPkg(next_goods_list []int32, goods_list []int32, req *cspb.CSPkg,
	res_list *cspb.CSPkgList, ret int32, random_type int32) int32 {

	//填充KeyRandomRes回包
	res_data := new(cspb.CSKeyRandomRes)
	*res_data = cspb.CSKeyRandomRes{
		Ret:           proto.Int32(ret),
		GoodsList:     goods_list,
		RandomType:    proto.Int32(random_type),
		NextGoodsList: next_goods_list,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		KeyRandomRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kKeyRandomRes),
		res_pkg_body, res_list)
	return ret
}
