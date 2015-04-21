package routers

import (
	"github.com/astaxie/beego"
	"tuojie.com/piggo/quickstart.git/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
