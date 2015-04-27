package cs_handle

import (
	"github.com/astaxie/beego"
)
import cspb "protocol"
import proto "code.google.com/p/goprotobuf/proto"

import db "tuojie.com/piggo/quickstart.git/db/collection"

func loginHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {

	beego.Debug("******loginHandle, req is %v, res is %v", req, res_list)
	ret, player := db.LoadPlayer(res_list.GetCAccount(), res_list.GetSAccount(), res_list.GetUid())
	beego.Debug("Db LoadPlayer Player, ret is %d, player is %v", ret, player)

	processMail(res_list)
	return makeLoginResPkg(req, res_list, ret)
}

func makeLoginResPkg(req *cspb.CSPkg,
	res_list *cspb.CSPkgList, ret int32) int32 {

	//填充petlistres回包
	res_data := new(cspb.CSLoginRes)
	*res_data = cspb.CSLoginRes{
		Ret: proto.Int32(ret),
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		LoginRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kLoginRes),
		res_pkg_body, res_list)
	return ret
}
