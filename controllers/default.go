package controllers

import (
	"github.com/astaxie/beego"
)

import "io/ioutil"
import cspb "protocol"
import proto "code.google.com/p/goprotobuf/proto"
import cs_handle "module/cs_handle"

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.tuojie.com"
	c.Data["Email"] = "info@tuojie.com"
	c.TplNames = "index.tpl"
}
func (c *MainController) Post() {
	beego.Debug("***********Post Start***********")
	data_req, _ := ioutil.ReadAll(c.Ctx.Input.Request.Body)
	c.Ctx.Input.Request.Body.Close()

	var pkg_list_req cspb.CSPkgList
	if err := proto.Unmarshal(data_req, &pkg_list_req); err != nil {
		beego.Error("main.procHttpMsg: unmarshal data_req to CSPkgList fail err:%v", err)
		return
	}

	beego.Info("*******Start and req is %v******", pkg_list_req)
	if len(pkg_list_req.GetPkgs()) == 0 {
		beego.Error("req no pkg account c:%s,s:%s",
			pkg_list_req.GetCAccount(),
			pkg_list_req.GetSAccount())
		return
	}

	var pkg_list_res = cspb.CSPkgList{
		CAccount: proto.String(pkg_list_req.GetCAccount()),
		SAccount: proto.String(pkg_list_req.GetSAccount()),
		Pkgs:     nil,
		Seq:      proto.Int64(pkg_list_req.GetSeq()),
		Uid:      proto.Int64(pkg_list_req.GetUid()),
	}

	cs_handle.PkgListHandle(pkg_list_req, &pkg_list_res)
	beego.Debug("Res is:%v", pkg_list_res)

	data_res, err := proto.Marshal(&pkg_list_res)
	if err != nil {
		beego.Error("marshal fail err:%v", err)
		return
	}

	if _, err := c.Ctx.Output.Context.ResponseWriter.Write(data_res); err != nil {
		beego.Error("write data fail err:%v", err)
		return
	}
}
