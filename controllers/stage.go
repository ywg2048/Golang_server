package controllers

import (
	// "encoding/json"
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	// "os"

	"strconv"
	"time"
	cs_handle "tuojie.com/piggo/quickstart.git/cs_handle"
	db_collection "tuojie.com/piggo/quickstart.git/db/collection"
)

//import "io/ioutil"

//import cspb "protocol"
//import proto "code.google.com/p/goprotobuf/proto"

//import cs_handle "tuojie.com/piggo/quickstart.git/cs_handle"
type Stagerecord struct {
	Id           int
	Name         string `orm:"size(100)"`
	Useracount   string `orm:"size(32)"`
	Stagename    string `orm:"size(256)"`
	Stagenum     string `orm:"size(256)"`
	Stagenamesub string `orm:"size(256)"`
	Time         int64
}
type StageController struct {
	beego.Controller
}

func init() {
	orm.RegisterModel(new(Stagerecord))

}
func (c *StageController) Get() {
	var users []Stagerecord
	var cond *orm.Condition
	cond = orm.NewCondition()

	cond = cond.And("Id__gte", 1)
	var qs orm.QuerySeter
	qs = orm.NewOrm().QueryTable("stagerecord").SetCond(cond)
	cnt, err := qs.All(&users)
	c.Data["s"] = users
	beego.Debug(cnt, err, users)

	c.TplNames = "stage.tpl"
}
func (c *StageController) Post() {
	beego.Info("Post")
	c.Ctx.Request.ParseForm()
	useracount := c.Input().Get("useracount")
	stagenum := c.Input().Get("stagenum")
	stagename := c.Input().Get("stagename")
	stagenamesub := c.Input().Get("stagename_sub")

	o := orm.NewOrm()
	var stagerecord Stagerecord

	stagerecord.Useracount = useracount
	stagerecord.Stagename = stagename
	stagerecord.Stagenum = stagenum

	stagerecord.Stagenamesub = stagenamesub
	stagerecord.Name = "WF"
	stagerecord.Time = time.Now().Unix()

	beego.Debug(stagerecord)

	id, err := o.Insert(&stagerecord)
	if err == nil {

		beego.Debug("插入成功！！", id)

	} else {

		beego.Error("插入失败！！！")

	}

	//查询道具表
	var users []Stagerecord
	var cond *orm.Condition
	cond = orm.NewCondition()

	cond = cond.And("Id__gte", 1)
	cond = cond.And("Useracount__contains", useracount)
	var qs orm.QuerySeter
	qs = orm.NewOrm().QueryTable("stagerecord").SetCond(cond)
	cnt, err := qs.All(&users)

	c.Data["s"] = users
	beego.Debug(users, cnt, err)

	//mongo
	uid, err_uid := strconv.Atoi(useracount)
	if err_uid != nil {

		beego.Error(useracount, err_uid, uid)

	}

	accessory_type, err_accessory_type := strconv.Atoi(stagename)
	if err_accessory_type != nil {

		beego.Error("str_accessory_type to int_accessory_type fail str_accessory_type:%s, err_accessory_type:%v",
			err_accessory_type)

	}
	accessory_sub_type, err_accessory_sub_type := strconv.Atoi(stagenamesub)
	if err_accessory_sub_type != nil {

		beego.Error("str_accessory_sub_type to int_accessory_sub_type fail str_accessory_sub_type:%s err_accessory_sub_type:%v",
			err_accessory_sub_type)

	}
	accessory_num, err_accessory_num := strconv.Atoi(stagenum)
	if err_accessory_num != nil {

		beego.Error("str_accessory_num to int_accessory_num fail str_accessory_num:%s err_accessory_num:%v",
			err_accessory_num)

	}

	var mail db_collection.Mail
	mail.Uid = int64(uid)
	mail.AccessoryType = int32(accessory_type)
	mail.AccessorySubType = int32(accessory_sub_type)
	mail.AccessoryNum = int32(accessory_num)

	beego.Debug("mailinfo:%v", mail)
	if cs_handle.CheckMail(mail) == false {

		beego.Error("mail param error mailinfo:%v", mail)

	}
	ret := db_collection.AddMail(mail)
	if ret == 0 {

		beego.Info("succ mailinfo:%v", mail)

	} else {

		beego.Error("add mail error ret:%d, mailinfo:%v", ret, mail)

	}
	c.TplNames = "stage.tpl"
}
