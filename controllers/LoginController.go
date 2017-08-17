package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"fmt"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get(){
	isExit := this.Input().Get("isExit")=="true"
	fmt.Println(isExit)
	if isExit {
		this.Ctx.SetCookie("uname", "", -1, "/")
		this.Ctx.SetCookie("password", "", -1, "/")
		this.Redirect("/admin", 301)
		return
	}
	this.TplName = "login.html"
}

func (this *LoginController) Post(){
	//this.TplName = "Login.html"
	uname := this.Input().Get("username");
	password := this.Input().Get("password");
	autoLogin := this.Input().Get("autoLogin") == "on"

	if beego.AppConfig.String("uname") == uname &&
	beego.AppConfig.String("password") == password{
		maxAge := 0
		if autoLogin{
			maxAge = 1<<31-1
		}
		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		this.Ctx.SetCookie("password", password, maxAge, "/")

	}
	this.Redirect("/admin", 301)
	return
}

func checkAccount(ctx *context.Context) bool{
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value;
	ck, err = ctx.Request.Cookie("password")
	if err != nil{
		return false
	}
	password := ck.Value

	return beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("password") == password
}


