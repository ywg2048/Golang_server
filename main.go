package main

import (
	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/auth"
	_ "github.com/go-sql-driver/mysql"
	"os"
	_ "tuojie.com/piggo/quickstart.git/docs"
	_ "tuojie.com/piggo/quickstart.git/routers"
)

import cs_handle "tuojie.com/piggo/quickstart.git/cs_handle"
import db_session "tuojie.com/piggo/quickstart.git/db/session"
import db_collection "tuojie.com/piggo/quickstart.git/db/collection"
import res_mgr "tuojie.com/piggo/quickstart.git/res_mgr"
import rand "github.com/tuojie/utility"

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

func main() {
	beego.EnableAdmin = true
	beego.AdminHttpAddr = "localhost"
	if beego.AppConfig.String("runmode") == "pro" {
		beego.SetLevel(beego.LevelInformational)
	}

	beego.SetLogger("file", `{"filename":"log/zoo_server.log"}`)
	Init()
	beego.InsertFilter("/stage", beego.BeforeRouter, auth.Basic(beego.AppConfig.String("authuser"), beego.AppConfig.String("authpwd")))
	beego.Run()

	Finish()
}

func Init() {

	beego.Debug("******Start main.Init******")
	orm.RegisterDataBase("default", "mysql", "root:@/Monsters?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(MonsterUser), new(Profile))

	// create table
	orm.RunSyncdb("default", false, true)
	//msg handle 初始化
	beego.Debug("cs_handle Init")
	cs_handle.Init()

	beego.Debug("config init")
	//config.InitConfig()

	//db  初始化
	beego.Debug("db init")
	err := db_session.Init(beego.AppConfig.String("mongodb_ip"))
	//err := db_session.Init("127.0.0.1:12306")
	if err != nil {
		beego.Error("Init db_session fail")
		os.Exit(1)
	}

	beego.Debug("init user id")
	db_collection.InitUserId()

	//随机数初始化
	beego.Debug("init rand succ")
	rand.Init()

	//resource 初始化
	beego.Debug("init res_mgr succ")
	res_mgr.Init()

	beego.Debug("******End main.Init******")
}

func Finish() {
	beego.Debug("******Start main.finish******")

	//DB 关闭
	beego.Info("db session close")
	db_session.Finish()

	//log 关闭,注意顺序，log关闭后无法再打印log了。所以要保证放到最后
	beego.Info("log close")
}
