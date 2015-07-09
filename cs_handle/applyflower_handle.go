package cs_handle

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
	"time"
	models "tuojie.com/piggo/quickstart.git/models"
)
import cspb "protocol"

// import proto "code.google.com/p/goprotobuf/proto"

// import db "tuojie.com/piggo/quickstart.git/db/collection"
import "labix.org/v2/mgo/bson"

import db_session "tuojie.com/piggo/quickstart.git/db/session"

// import resmgr "tuojie.com/piggo/quickstart.git/res_mgr"

func ApplyFlowerHandle(
	req *cspb.CSPkg,
	res_list *cspb.CSPkgList) int32 {
	beego.Info("*********ApplyFlowerHandle Start**********")
	req_data := req.GetBody().GetApplyflowerReq()
	beego.Info(req_data)
	ret := int32(1)

	now := time.Now().Unix()
	now_time := time.Unix(now, 0).Format("2006-01-02 15:04:05")

	var now_array []string
	now_array = strings.Split(now_time, "-")
	var now_array_3 []string
	now_array_3 = strings.Split(now_array[2], " ")
	year := now_array[0]
	month := now_array[1]
	day := now_array_3[0]

	beego.Info(year, month, day, now)

	c := db_session.DB("zoo").C("player")
	var player models.Player
	for i := range req_data.GetApplyFlowerList() {

		//查找用户是否存在
		err := c.Find(bson.M{"uid": req_data.GetApplyFlowerList()[i].GetFriendId()}).One(&player)
		if err != nil {
			beego.Error("找不到用户！")
		} else {
			beego.Info("找到用户！")
			var messagesList []models.Messages
			var cond *orm.Condition
			cond = orm.NewCondition()

			cond = cond.And("Id__gte", 1)
			cond = cond.And("Touid__contains", req_data.GetApplyFlowerList()[i].GetFriendId())
			cond = cond.And("Fromuid__contains", int32(res_list.GetUid()))
			cond = cond.And("Messagetype__contains", int32(1))
			cond = cond.And("ElementType__contains", int32(1))
			// cond = cond.And("IsFinish__contains", int32(0))
			beego.Info(cond)
			var qs orm.QuerySeter
			qs = orm.NewOrm().QueryTable("messages").SetCond(cond)

			cnt, err := qs.All(&messagesList)
			beego.Info(cnt, err, messagesList)

			maxTime := int64(0)
			if len(messagesList) > 0 {

				maxTime = messagesList[len(messagesList)-1].Time
			}
			str_time := time.Unix(maxTime, 0).Format("2006-01-02 15:04:05")

			var data_array []string
			data_array = strings.Split(str_time, "-")
			var data_array_3 []string
			data_array_3 = strings.Split(data_array[2], " ")
			maxYear := data_array[0]
			maxmonth := data_array[1]
			maxday := data_array_3[0]

			beego.Info("maxtime", maxTime, str_time)
			//消息提示，存在Mysql中
			// && stringToint(month) >= stringToint(maxmonth) && stringToint(day) > stringToint(maxday)

			if stringToint(year) > stringToint(maxYear) {
				var players models.Player
				c.Find(bson.M{"c_account": res_list.GetCAccount()}).One(&players)

				o := orm.NewOrm()
				var messages models.Messages
				messages.Fromuid = int32(res_list.GetUid())
				messages.Fromname = players.Name
				messages.FromStarId = players.StarId
				messages.Time = time.Now().Unix()
				messages.IsFinish = int32(0)
				messages.Messagetype = int32(1)
				messages.ElementType = int32(1)
				messages.Tag = int32(1)
				messages.Touid = req_data.GetApplyFlowerList()[i].GetFriendId()

				id, err := o.Insert(&messages)
				if err == nil {
					fmt.Println(id)
				}
			} else if stringToint(year) == stringToint(maxYear) {
				if stringToint(month) > stringToint(maxmonth) {
					var players models.Player
					c.Find(bson.M{"c_account": res_list.GetCAccount()}).One(&players)

					o := orm.NewOrm()
					var messages models.Messages
					messages.Fromuid = int32(res_list.GetUid())
					messages.Fromname = players.Name
					messages.FromStarId = players.StarId
					messages.Time = time.Now().Unix()
					messages.IsFinish = int32(0)
					messages.Messagetype = int32(1)
					messages.ElementType = int32(1)
					messages.Tag = int32(1)
					messages.Touid = req_data.GetApplyFlowerList()[i].GetFriendId()

					id, err := o.Insert(&messages)
					if err == nil {
						fmt.Println(id)
					}
				} else if stringToint(month) == stringToint(maxmonth) {
					if stringToint(day) > stringToint(maxday) {
						var players models.Player
						c.Find(bson.M{"c_account": res_list.GetCAccount()}).One(&players)

						o := orm.NewOrm()
						var messages models.Messages
						messages.Fromuid = int32(res_list.GetUid())
						messages.Fromname = players.Name
						messages.FromStarId = players.StarId
						messages.Time = time.Now().Unix()
						messages.IsFinish = int32(0)
						messages.Messagetype = int32(1)
						messages.ElementType = int32(1)
						messages.Tag = int32(1)
						messages.Touid = req_data.GetApplyFlowerList()[i].GetFriendId()

						id, err := o.Insert(&messages)
						if err == nil {
							fmt.Println(id)
						}

					} else {
						beego.Info("同一天不能再次赠送")
					}
				} else {
					beego.Info("条件不足")
				}
			} else {
				beego.Info("条件不足")
			}

		}

	}
	res_data := new(cspb.CSApplyFlowerRes)
	*res_data = cspb.CSApplyFlowerRes{}

	res_pkg_body := new(cspb.CSBody)
	*res_pkg_body = cspb.CSBody{
		ApplyflowerRes: res_data,
	}
	res_list = makeCSPkgList(int32(cspb.Command_kApplyFlowerRes),
		res_pkg_body, res_list)
	return ret

}
func stringToint(value string) int {
	b, error := strconv.Atoi(value)
	if error != nil {
		fmt.Println("字符串转换成整数失败")
	}
	return b
}
