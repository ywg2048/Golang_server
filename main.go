package main

import (
	"github.com/astaxie/beego"
	"os"
	_ "tuojie.com/piggo/quickstart.git/routers"
)

import cs_handle "module/cs_handle"
import db_session "module/db/session"
import db_collection "module/db/collection"
import res_mgr "module/res_mgr"
import rand "github.com/tuojie/utility"
import config "config_data"

func main() {
	Init()
	beego.Run()
	Finish()
}

func Init() {
	beego.Debug("******Start main.Init******")

	//msg handle 初始化
	beego.Debug("cs_handle Init")
	cs_handle.Init()

	beego.Debug("config init")
	config.InitConfig()

	//db  初始化
	beego.Debug("db init")
	err := db_session.Init(config.Config().GetDbIp())
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
