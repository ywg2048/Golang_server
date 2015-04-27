package cs_handle

import (
	"github.com/astaxie/beego"
)
import cspb "protocol"
import proto "code.google.com/p/goprotobuf/proto"
import db "tuojie.com/piggo/quickstart.git/db/collection"

import "time"

func rechargeHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {

	beego.Debug("******loginHandle, req is %v, res is %v", req, res_list)
	ret, player := db.LoadPlayer(res_list.GetCAccount(), res_list.GetSAccount(), res_list.GetUid())
	beego.Debug("Db LoadPlayer Player, ret is %d, player is %v", ret, player)

	req_data := req.GetBody().GetRechargeReportReq()
	beego.Debug("req_data:%s", req_data.String())
	var recharge_flow db.RechargeFlowData

	// recharge_flow.Account = res_list.GetSAccount()
	// recharge_flow.Uid = res_list.GetUid()
	recharge_flow.Rmb = req_data.GetRmb()
	recharge_flow.GoodsType = req_data.GetGoodsType()
	recharge_flow.GoodsSubType = req_data.GetGoodsSubType()
	recharge_flow.GoodsNum = req_data.GetGoodsNum()
	recharge_flow.RechargeTime = time.Now().Unix() - int64(60*60*8)
	recharge_flow.Version = req_data.GetVersion()
	recharge_flow.Code = req_data.GetCode()
	recharge_flow.Channel = req_data.GetChannel()

	player.RechargeFlow = recharge_flow

	//ret := AddRechargeFlow(&recharge_flow)

	if ret != 0 {
		beego.Error("add recharge fail ret:%d", ret)
		ret = int32(-1)
		return makeRechargeResPkg(req, res_list, ret)
	}

	return makeRechargeResPkg(req, res_list, int32(0))
}

func makeRechargeResPkg(req *cspb.CSPkg, res_list *cspb.CSPkgList, ret int32) int32 {

	//填充RechargeRes回包
	res_data := new(cspb.CSRechargeReportRes)
	*res_data = cspb.CSRechargeReportRes{
		Ret:     proto.Int32(ret),
		OrderId: proto.Int64(req.GetBody().GetRechargeReportReq().GetOrderId()),
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		RechargeReportRes: res_data,
	}

	res_list = makeCSPkgList(int32(cspb.Command_kRechargeReportRes),
		res_pkg_body, res_list)
	return ret
}
