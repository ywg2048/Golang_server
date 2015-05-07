package routers

import (
	"github.com/astaxie/beego"
	"github.com/beego/admin"
	"tuojie.com/piggo/quickstart.git/controllers"
)

func init() {

	beego.Router("/", &controllers.MainController{})
	beego.Router("/stage", &controllers.StageController{})
	beego.Router("/mobileweb", &controllers.MobilewebController{})
	beego.Router("/foshan/rest/game/app", &controllers.AppController{})
	beego.Router("/message", &controllers.MessagecenterController{})
	admin.Run()
}
