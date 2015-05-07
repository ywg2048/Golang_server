package cs_handle

import (
	"github.com/astaxie/beego"
)
import cspb "protocol"

// cmd proc func array
var procFuncArray [cspb.Command_kMaxCount]cmdProcess

func Init() {
	beego.Debug("******cs_handle framehandle Init")
	//初始化process array
	for i := 0; i < int(cspb.Command_kMaxCount); i++ {
		procFuncArray[i] = nil
	}

	//注册handle
	cmdProcessRegister(cspb.Command_kRegistPlayerReq, registPlayerHandle)
	cmdProcessRegister(cspb.Command_kPetListReq, getPetListHandle)
	cmdProcessRegister(cspb.Command_kChipListReq, getChipListHandle)
	cmdProcessRegister(cspb.Command_kKeyRandomReq, keyRandomHandle)
	cmdProcessRegister(cspb.Command_kPetLevelUpReq, petLevelUpHandle)
	cmdProcessRegister(cspb.Command_kRandomItemReq, randomItemHandle)
	cmdProcessRegister(cspb.Command_kSignInReq, signinHandle)
	cmdProcessRegister(cspb.Command_kPetStarUpReq, petStarUpHandle)
	cmdProcessRegister(cspb.Command_kReportPetReq, petReportHandle)
	cmdProcessRegister(cspb.Command_kServerTimeReq, serverTimeHandle)
	cmdProcessRegister(cspb.Command_kUpdateInfoReq, updateInfoHandle)
	cmdProcessRegister(cspb.Command_kRechargeReportReq, rechargeHandle)
	cmdProcessRegister(cspb.Command_kLoginReq, loginHandle)
	cmdProcessRegister(cspb.Command_kStageReportReq, stageReportHandle)
	cmdProcessRegister(cspb.Command_kMessageCenterReq, MessageCenterHandle)

}

func PkgListHandle(
	pkg_list_req cspb.CSPkgList,
	pkg_list_res *cspb.CSPkgList) {
	for _, req := range pkg_list_req.GetPkgs() {
		cmd := req.GetHead().GetCmd()
		proc := getCmdProc(cmd)
		if proc == nil {
			beego.Error("no find cmd proc cmd:%d", cmd)
			continue
		}
		beego.Debug("Reqinfo cmd is%d & command name is %s req is %s",
			cmd,
			cspb.Command_name[cmd],
			req.String())
		ret := proc(req, pkg_list_res)
		if ret != 0 {
			beego.Error("cmd proc fail ret:%d", ret)
			continue
		}
	}
}

type cmdProcess func(
	*cspb.CSPkg,
	*cspb.CSPkgList) int32

func getCmdProc(cmd int32) cmdProcess {
	if cmd >= int32(cspb.Command_kMaxCount) || cmd <= 0 {
		return nil
	}

	return procFuncArray[cmd]
}

// Register your client cmd proc function
func cmdProcessRegister(cmd cspb.Command, proc cmdProcess) int {
	if cmd >= cspb.Command_kMaxCount || cmd <= 0 {
		return -1
	}

	if procFuncArray[cmd] != nil {
		return 1
	}

	procFuncArray[cmd] = proc
	return 0
}
