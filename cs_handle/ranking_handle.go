package cs_handle

import (
	// "fmt"
	"github.com/astaxie/beego"
	// "time"
	"github.com/astaxie/beego/orm"

	models "tuojie.com/piggo/quickstart.git/models"
)
import cspb "protocol"

// import proto "code.google.com/p/goprotobuf/proto"

// import db "tuojie.com/piggo/quickstart.git/db/collection"
// import "labix.org/v2/mgo/bson"

// import db_session "tuojie.com/piggo/quickstart.git/db/session"

func RankHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********RankHandle Start**********")
	req_data := req.GetBody().GetRankReq()
	beego.Info(req_data)

	ret := int32(0)

	var ranking []models.Ranking
	var cond *orm.Condition
	cond = orm.NewCondition()

	cond = cond.And("Id__gte", 1)
	cond = cond.And("IsActive__contains", 1)
	var qs orm.QuerySeter
	qs = orm.NewOrm().QueryTable("ranking").SetCond(cond)
	cnt, err := qs.All(&ranking)
	if err != nil {
		beego.Debug("查询数据库失败")
	}
	beego.Debug(ranking, cnt, err)

	var res_rank []*cspb.CSRankNtf
	// switch req_data.GetType() {
	// case int32(1):
	// 	//小伙伴排名
	// 	res_rank = append(res_rank, makeRank())
	// case int32(2):
	// 	//明星排名
	// 	res_rank = append(res_rank, makeRank())
	// case int32(3):
	// 	//好友排名
	// 	res_rank = append(res_rank, makeRank())
	// }
	res_data := new(cspb.CSRankRes)
	*res_data = cspb.CSRankRes{
		RankNtf: res_rank,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		RankRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kRankRes),
		res_pkg_body, res_list)
	return ret

}
