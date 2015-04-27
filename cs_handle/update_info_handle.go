package cs_handle

import (
	"github.com/astaxie/beego"
)
import cspb "protocol"
import proto "code.google.com/p/goprotobuf/proto"

//import config "config_data"
import resmgr "tuojie.com/piggo/quickstart.git/res_mgr"

//import "github.com/astaxie/beego"

func updateInfoHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {

	beego.Debug("******updateInfoHandle******")
	req_version := req.GetBody().GetUpdateInfoReq().GetVersion()
	req_osType := req.GetBody().GetUpdateInfoReq().GetOSType()
	//	req_code := req.GetBody().GetUpdateInfoReq().GetCode()
	//req_channel := req.GetBody().GetUpdateInfoReq().GetChannel()
	strategy := int32(cspb.StrategyType_Newst)
	url := ""
	strategyName := ""
	hasUpdate := int32(0)
	latestVersion := ""
	canHotUpdate := int32(0)
	hotUpdateVersion := ""

	beego.Debug("req_osType is %s", req_osType)
	beego.Debug("UpdateconfigsData count is %d", int32(len(resmgr.UpdateconfigsData.GetItems())))
	for _, update_config := range resmgr.UpdateconfigsData.GetItems() {
		if req_version == update_config.GetVersion() && req_osType == update_config.GetOSType() {
			beego.Debug("update config is %v", update_config)
			strategy = update_config.GetStrategyType()
			url = update_config.GetUrl()
			strategyName = update_config.GetStrategyName()
			hasUpdate = update_config.GetHasUpdate()
			latestVersion = update_config.GetLatestVersion()
			canHotUpdate = update_config.GetCanHotUpdate()
			hotUpdateVersion = update_config.GetHotUpdateVersion()
			break
		}
	}

	ret := int32(0)

	//填充CSUpdateInfoRes回包
	beego.Debug("填充CSUpdateInfoRes回包")
	res_data := new(cspb.CSUpdateInfoRes)
	*res_data = cspb.CSUpdateInfoRes{
		Ret:              proto.Int32(ret),
		Strategy:         proto.Int32(strategy),
		Url:              proto.String(url),
		StrategyName:     proto.String(strategyName),
		HasUpdate:        proto.Int32(hasUpdate),
		LatestVersion:    proto.String(latestVersion),
		CanHotUpdate:     proto.Int32(canHotUpdate),
		HotUpdateVersion: proto.String(hotUpdateVersion),
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		UpdateInfoRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kUpdateInfoRes),
		res_pkg_body, res_list)
	return ret
}
