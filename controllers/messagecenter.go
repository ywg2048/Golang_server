package controllers

import (
	// "encoding/json"
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	// "os"
	"strconv"
	"time"
	models "tuojie.com/piggo/quickstart.git/models"
	// cs_handle "tuojie.com/piggo/quickstart.git/cs_handle"
	// db_collection "tuojie.com/piggo/quickstart.git/db/collection"
)

type MessagecenterController struct {
	beego.Controller
}

func init() {

	orm.RegisterDataBase("default", "mysql", "root:@/Monsters?charset=utf8")

}
func (c *MessagecenterController) Get() {

	id := c.Input().Get("id")
	opration := c.Input().Get("opration")
	beego.Info(id)
	beego.Info(opration)

	ids, err := strconv.Atoi(id)
	id1 := int32(ids)
	beego.Info(id1)
	// 激活操作
	// if opration == "active" {
	// 	beego.Info("激活")

	// 	o := orm.NewOrm()
	// 	messagecenter := []models.Messagecenter{Id: id1}
	// 	if o.Read(&messagecenter) == nil {
	// 		messagecenter.IsActive = 1
	// 		if num, err := o.Update(&messagecenter); err == nil {
	// 			fmt.Println(num)
	// 		}
	// 	}
	// } else {
	// 	beego.Info("关闭")
	// 	o := orm.NewOrm()
	// 	messagecenter := []models.Messagecenter{Id: id1}
	// 	if o.Read(&messagecenter) == nil {
	// 		messagecenter.IsActive = 0
	// 		if num, err := o.Update(&messagecenter); err == nil {
	// 			fmt.Println(num)
	// 		}
	// 	}
	// }
	//读取列表
	var message []models.Messagecenter
	var cond *orm.Condition
	cond = orm.NewCondition()

	cond = cond.And("Id__gte", 1)

	var qs orm.QuerySeter
	qs = orm.NewOrm().QueryTable("messagecenter").Limit(20).SetCond(cond)
	cnt, err := qs.All(&message)

	c.Data["message"] = message
	beego.Debug(message, cnt, err)
	c.TplNames = "announcement.tpl"
}
func (c *MessagecenterController) Post() {
	beego.Info("Post")
	c.Ctx.Request.ParseForm()
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	title2 := c.Input().Get("title2")
	content2 := c.Input().Get("content2")

	o := orm.NewOrm()
	var messagecenter models.Messagecenter

	messagecenter[0].Title = title
	messagecenter[0].Content = content
	messagecenter[0].Title2 = title2
	messagecenter[0].Content2 = content2
	messagecenter[0].Time = time.Now().Unix()
	messagecenter[0].IsActive = 0
	beego.Debug(messagecenter)

	id, err := o.Insert(&messagecenter)
	if err == nil {

		beego.Debug("插入成功！！", id)

	} else {

		beego.Error("插入失败！！！")

	}

	var message []models.Messagecenter
	var cond *orm.Condition
	cond = orm.NewCondition()

	cond = cond.And("Id__gte", 1)

	var qs orm.QuerySeter
	qs = orm.NewOrm().QueryTable("messagecenter").Limit(20).SetCond(cond)
	cnt, err := qs.All(&message)

	c.Data["message"] = message
	beego.Debug(message, cnt, err)

	c.TplNames = "announcement.tpl"
}
