package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//这里一个get请求
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	//这个data表示返回数据字段和内容
	c.TplName = "index.tpl"
//	表示指向的模板文件，在view里面
}
