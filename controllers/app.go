package controllers

import (
	"github.com/astaxie/beego"
)

type AppController struct {
	beego.Controller
}

func (c *AppController) Get() {

	c.TplNames = "app.tpl"
}
func (c *AppController) Post() {

}
