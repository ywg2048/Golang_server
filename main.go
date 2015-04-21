package main

import (
	"github.com/astaxie/beego"
	_ "tuojie.com/piggo/quickstart.git/routers"
)

//import "fmt"
//import "io/ioutil"
//import "net/http"
import "os"

//import "flag"
//import cspb "protocol"
//import proto "code.google.com/p/goprotobuf/proto"
import log "code.google.com/p/log4go"
import cs_handle "module/cs_handle"
import db_session "module/db/session"
import db_collection "module/db/collection"
import res_mgr "module/res_mgr"
import rand "github.com/tuojie/utility"
import config "config_data"

func main() {
	Init()
	beego.Run()
	finish()
}

func Init() {

	//log的初始化 log 初始化要放倒最前面，再log初始化之前无法打印log
	log.LoadConfiguration("./conf/zoo_log.xml")
	log.Debug("******Start main.Init******")
	//命令行初始化
	//flag.Parse()

	//msg handle 初始化
	log.Debug("cs_handle Init")
	cs_handle.Init()

	log.Debug("config init")
	config.InitConfig()

	//db  初始化
	log.Debug("db init")
	err := db_session.Init(config.Config().GetDbIp())
	//err := db_session.Init("127.0.0.1:12306")
	if err != nil {
		log.Error("Init db_session fail")
		os.Exit(1)
	}

	log.Debug("init user id")
	db_collection.InitUserId()

	//随机数初始化
	log.Debug("init rand succ")
	rand.Init()

	//resource 初始化
	log.Debug("init res_mgr succ")
	res_mgr.Init()

	log.Debug("******End main.Init******")
}

func finish() {
	log.Debug("******Start main.finish******")

	//DB 关闭
	log.Info("db session close")
	db_session.Finish()

	//log 关闭,注意顺序，log关闭后无法再打印log了。所以要保证放到最后
	log.Info("log close")
	log.Close()
}
