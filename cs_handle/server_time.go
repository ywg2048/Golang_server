package cs_handle

import cspb "protocol"
import proto "code.google.com/p/goprotobuf/proto"
import log "code.google.com/p/log4go"
import "time"

func serverTimeHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {

	log.Debug("******serverTimeHandle")
	ret := int32(0)
	now_time := time.Now().Unix()
	log.Debug("server time :%d", now_time)
	//填充ServerTimeRes回包
	res_data := new(cspb.CSServerTimeRes)
	*res_data = cspb.CSServerTimeRes{
		Ret:        proto.Int32(ret),
		ServerTime: proto.Int64(now_time),
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		ServerTimeRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kServerTimeRes),
		res_pkg_body, res_list)
	log.Debug("serverTimeHandle******")
	return ret
}
