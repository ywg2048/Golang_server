package controllers

import (
	"github.com/astaxie/beego"
)

import "fmt"
import "io/ioutil"
import "net/http"

//import "os"
//import "flag"
import cspb "protocol"
import proto "code.google.com/p/goprotobuf/proto"
import log "code.google.com/p/log4go"
import cs_handle "module/cs_handle"

//import "flag"

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.tpl"
}

func procHttpMsg(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数
	if r.Method == "GET" {
		fmt.Fprintf(w, "{\"error\":\"please use post\"}")
		log.Error("error: client is use get mode")

	} else if r.Method == "POST" {
		log.Info("RemoteAddr is :%s", r.RemoteAddr)

		data_req, _ := ioutil.ReadAll(r.Body)
		r.Body.Close()

		var pkg_list_req cspb.CSPkgList
		if err := proto.Unmarshal(data_req, &pkg_list_req); err != nil {
			log.Error("main.procHttpMsg: unmarshal data_req to CSPkgList fail err:%v", err)
			return
		}

		log.Info("*******Start and req is %v******", pkg_list_req)
		if len(pkg_list_req.GetPkgs()) == 0 {
			log.Error("req no pkg account c:%s,s:%s",
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
		log.Debug("Res is:%v", pkg_list_res)

		data_res, err := proto.Marshal(&pkg_list_res)
		if err != nil {
			log.Error("marshal fail err:%v", err)
			return
		}

		if _, err := w.Write(data_res); err != nil {
			log.Error("write data fail err:%v", err)
			return
		}
		log.Info("******End and res is %v******", pkg_list_res)
	}
}
