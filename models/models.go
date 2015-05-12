package models

import (
	"github.com/astaxie/beego/orm"
)

type Messagecenter struct {
	Id      int32
	Title   string `orm:"size(100)"`
	Content string `orm:"size(256)"`

	IsActive int32
	Time     int64
}
type Userinfo struct {
	Id        int32
	Uid       int64
	userscore []Userscore
}
type Userscore struct {
	Id       int32
	Uid      int64
	Level    int32
	Score    int32
	Startnum int32
	Time     int64
}

func init() {

	orm.RegisterModel(new(Messagecenter), new(Userscore), new(Userinfo))

}
