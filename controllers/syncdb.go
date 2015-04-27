package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//import "io/ioutil"
//import cspb "protocol"
//import proto "code.google.com/p/goprotobuf/proto"
//import cs_handle "tuojie.com/piggo/quickstart.git/cs_handle"

// Model Struct
type MonsterUser struct {
	Id      int
	Name    string   `orm:"size(100)"`
	Profile *Profile `orm:"rel(one)"` // OneToOne relation
}
type Profile struct {
	Id          int
	Age         int16
	MonsterUser *MonsterUser `orm:"reverse(one)"` // 设置反向关系(可选)
}

type SyncdbController struct {
	beego.Controller
}

func (c *SyncdbController) Get() {
	orm.RegisterDataBase("default", "mysql", "root:@/Monsters?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(MonsterUser), new(Profile))

	// create table
	orm.RunSyncdb("default", false, true)
}
