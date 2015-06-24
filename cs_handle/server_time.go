package cs_handle

import (
	"github.com/astaxie/beego"
)
import cspb "protocol"
import proto "code.google.com/p/goprotobuf/proto"

import "time"

func serverTimeHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {

	beego.Debug("******serverTimeHandle")
	ret := int32(0)
	now_time := time.Now().Unix()
	beego.Debug("server time :%d", now_time)
	//填充ServerTimeRes回包
	messageList := int32(5)
	solution := int32(2)
	achievement := int32(1)
	zoo := int32(2)
	res_data := new(cspb.CSServerTimeRes)
	*res_data = cspb.CSServerTimeRes{
		Ret:         proto.Int32(ret),
		ServerTime:  proto.Int64(now_time),
		MessageList: &messageList,
		Solution:    &solution,
		Achievement: &achievement,
		Zoo:         &zoo,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		ServerTimeRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kServerTimeRes),
		res_pkg_body, res_list)
	beego.Debug("serverTimeHandle******")
	return ret
}
