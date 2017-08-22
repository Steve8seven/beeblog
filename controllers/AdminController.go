package controllers

import (
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Get() {
	check := checkAccount(this.Ctx)
	if check{
		this.TplName = "admin.html"
	}else {
		this.Redirect("/login", 301)
		//this.TplName = "login.html"
	}
	return
}

func (this *AdminController) Post() {

}
