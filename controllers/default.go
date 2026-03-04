package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.vip"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	// 方法1：直接写字符串到响应
	c.Ctx.WriteString("Hello, World!")
}
