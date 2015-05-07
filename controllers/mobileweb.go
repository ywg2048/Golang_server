package controllers

import (
	"github.com/astaxie/beego"
)

type MobilewebController struct {
	beego.Controller
}

func (c *MobilewebController) Get() {

	c.TplNames = "mobileweb.tpl"
}
func (c *MobilewebController) Post() {

}
