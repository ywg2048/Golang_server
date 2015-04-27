package routers

import (
	"github.com/astaxie/beego"
	"github.com/beego/admin"
	"tuojie.com/piggo/quickstart.git/controllers"
)

func init() {

	beego.Router("/", &controllers.MainController{})
	beego.Router("/stage", &controllers.StageController{})
	admin.Run()
}
