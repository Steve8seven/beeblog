package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Email"] = "inhangzhouzucc@163.com"
	c.Data["isHome"] = true
	c.Data["isAdmin"] = false

	isAdmin := checkAccount(c.Ctx);

	c.Data["IsLogin"] = isAdmin
	if isAdmin {
		c.Data["isAdmin"] = true
	}
	c.TplName = "home.html"
}
