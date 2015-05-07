package models

import (
	"github.com/astaxie/beego/orm"
)

type Messagecenter []struct {
	Id       int32
	Title    string `orm:"size(100)"`
	Content  string `orm:"size(256)"`
	Title2   string `orm:"size(256)"`
	Content2 string `orm:"size(256)"`
	IsActive int32
	Time     int64
}

func init() {

	orm.RegisterModel(new(Messagecenter))

}
