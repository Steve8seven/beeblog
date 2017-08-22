package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) Get() {
	check := checkAccount(this.Ctx)
	fmt.Println(check)
	this.TplName = "article_edit.html"
	if check{
		this.TplName = "article_edit.html"
	}else {
		this.Redirect("/login", 301)
		//this.TplName = "login.html"
	}
	return
}

func (this *ArticleController) Post() {

}

func (this *ArticleController) SaveArticle()  {
	mdcode := this.Input().Get("mdcode")
	htmlcode := this.Input().Get("htmlcode")
	fmt.Println(mdcode)
	fmt.Println(htmlcode)
	this.Data["json"] = "okok"
	this.ServeJSON()
}
