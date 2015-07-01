package cs_handle

import (
	// "fmt"
	"github.com/astaxie/beego"
	// "time"
	"github.com/astaxie/beego/orm"

	models "tuojie.com/piggo/quickstart.git/models"
)
import cspb "protocol"

import proto "code.google.com/p/goprotobuf/proto"

// import db "tuojie.com/piggo/quickstart.git/db/collection"
// import "labix.org/v2/mgo/bson"

// import db_session "tuojie.com/piggo/quickstart.git/db/session"
import resmgr "tuojie.com/piggo/quickstart.git/res_mgr"

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

	var qs orm.QuerySeter

	var res_rank []*cspb.CSRankNtf
	//测试代码

	for i := range resmgr.RankingtestData.GetItems() {
		res_rank = append(res_rank, makeRank(resmgr.RankingtestData.GetItems()[i].GetUID(), resmgr.RankingtestData.GetItems()[i].GetName(), resmgr.RankingtestData.GetItems()[i].GetLevel(), resmgr.RankingtestData.GetItems()[i].GetMedalNum(), resmgr.RankingtestData.GetItems()[i].GetRankId(), resmgr.RankingtestData.GetItems()[i].GetMedalLevelID(), resmgr.RankingtestData.GetItems()[i].GetStarId()))
	}

	//正式代码
	StarId := req_data.GetStarid()
	switch req_data.GetType() {
	case "1":
		//小伙伴排名
		qs = orm.NewOrm().QueryTable("ranking").SetCond(cond).OrderBy("-Medal")
		cnt, err := qs.All(&ranking)
		if err != nil {
			beego.Debug("查询数据库失败")
		}
		beego.Debug(ranking, cnt, err)
		for i := range ranking {
			res_rank = append(res_rank, makeRank(ranking[i].Uid, ranking[i].Name, ranking[i].Level, ranking[i].Medal, int32(i+1), ranking[i].MedalLevelId, StarId))
		}
	case "2":
		//明星排名
		cond = cond.And("StarId_contains", StarId)
		qs = orm.NewOrm().QueryTable("ranking").SetCond(cond).OrderBy("-Medal")
		cnt, err := qs.All(&ranking)
		if err != nil {
			beego.Debug("查询数据库失败")
		}
		beego.Debug(ranking, cnt, err)
		for i := range ranking {
			res_rank = append(res_rank, makeRank(ranking[i].Uid, ranking[i].Name, ranking[i].Level, ranking[i].Medal, int32(i+1), ranking[i].MedalLevelId, StarId))
		}

	case "3":
		//好友排名(查找朋友列表的数据)

	}

	SearchType := req_data.GetType()
	beego.Info(res_rank)
	ret = int32(1)
	res_data := new(cspb.CSRankRes)
	*res_data = cspb.CSRankRes{
		Ret:        proto.Int32(ret),
		SearchType: &SearchType,
		RankNtf:    res_rank,
	}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		RankRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kRankRes),
		res_pkg_body, res_list)
	return ret

}
