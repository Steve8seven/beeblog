package controllers

import (
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Get(){
	this.TplName = "admin.html"
}

func (this *AdminController) Post(){

}