package cs_handle

import (
	// "fmt"
	"github.com/astaxie/beego"
	// "time"
	// models "tuojie.com/piggo/quickstart.git/models"
)
import cspb "protocol"

// import proto "code.google.com/p/goprotobuf/proto"

// import db "tuojie.com/piggo/quickstart.git/db/collection"
// import "labix.org/v2/mgo/bson"

// import db_session "tuojie.com/piggo/quickstart.git/db/session"
import resmgr "tuojie.com/piggo/quickstart.git/res_mgr"

func ZooHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********ZooHandle Start**********")
	req_data := req.GetBody().GetZooReq()
	beego.Info(req_data)
	ret := int32(1)

	var AnimalNtf []*cspb.CSAnimalNtf
	for i := range resmgr.ZootestData.GetItems() {
		AnimalNtf = append(AnimalNtf, makeAnimalNtf(resmgr.ZootestData.GetItems()[i].GetAnimalId(), resmgr.ZootestData.GetItems()[i].GetStatus()))
	}

	res_data := new(cspb.CSZooRes)

	*res_data = cspb.CSZooRes{
		AnimalNtf: AnimalNtf,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		ZooRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kZooRes),
		res_pkg_body, res_list)
	return ret

}
